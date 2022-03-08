package sync_handler

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
)

var _ SyncHandlerService = &SyncHandler{}

type SyncHandlerService interface {
	CheckVersionMismatch(localRunway domain.Runway) (bool, error)
	GetConflictingFields(localRunway domain.Runway, remoteRunway domain.Runway) domain.Conflict
	CreateConflict(localRunway domain.Runway, clientID string) (domain.Conflict, error)
}

type SyncHandler struct {
	db repository.Repository
}

func New(db repository.Repository) (SyncHandler, error) {
	return SyncHandler{
		db: db,
	}, nil
}

func (s SyncHandler) CheckVersionMismatch(localRunway domain.Runway) (bool, error) {
	remoteRunway, err := s.db.GetRunwayByID(localRunway.ID)

	if err != nil {
		log.WithError(err).WithField("id", localRunway.ID).Error("Could not get runway from db")
		return false, err
	}

	if remoteRunway.LatestVersion == localRunway.LatestVersion {
		log.WithFields(log.Fields{
			"localVersion":  localRunway.LatestVersion,
			"remoteVersion": remoteRunway.LatestVersion,
			"id":            localRunway.ID,
		}).Info("Versions matching")
		return false, nil
	}

	log.WithFields(log.Fields{
		"localVersion":  localRunway.LatestVersion,
		"remoteVersion": remoteRunway.LatestVersion,
		"id":            localRunway.ID,
	}).Info("Versions are not matching")
	return true, nil
}

func (s SyncHandler) CreateConflict(localRunway domain.Runway, clientID string) (domain.Conflict, error) {
	remoteRunway, err := s.db.GetRunwayByID(localRunway.ID)

	if err != nil {
		log.WithError(err).WithField("id", localRunway.ID).Error("Could not get runway from db")
		return domain.Conflict{}, err
	}

	log.WithField("id", localRunway.ID).Info("Getting conflicting fields")
	conflict := s.GetConflictingFields(localRunway, remoteRunway)

	conflict.ClientID = clientID
	conflictID, err := s.db.CreateConflict(conflict)

	conflict.ID = conflictID

	if err != nil {
		log.WithError(err).WithField("id", localRunway.ID).Error("Could not create conflict in db")
		return domain.Conflict{}, err
	}

	log.WithFields(log.Fields{"id": localRunway.ID, "conflictID": conflictID}).Info("Conflict created in db")
	return conflict, nil
}

func (s SyncHandler) GetConflictingFields(localRunway domain.Runway, remoteRunway domain.Runway) domain.Conflict {
	localRunwayElems := reflect.ValueOf(&localRunway).Elem()
	remoteRunwayElems := reflect.ValueOf(&remoteRunway).Elem()

	typeOfT := localRunwayElems.Type()
	typeOfRemoteElem := remoteRunwayElems.Type()

	diff := make(map[string]map[string]interface{})
	diff["LOCAL"] = make(map[string]interface{})
	diff["REMOTE"] = make(map[string]interface{})

	for i := 0; i < localRunwayElems.NumField(); i++ {
		f := localRunwayElems.Field(i)
		switch f.Type().String() {
		case "bool", "string", "int":
			if compareSingleValue(f.Interface(), remoteRunwayElems.Field(i).Interface()) {
				diff["LOCAL"][getJSONTag(string(typeOfRemoteElem.Field(i).Tag), typeOfRemoteElem.Field(i).Name)] = f.Interface()
				diff["REMOTE"][getJSONTag(string(typeOfRemoteElem.Field(i).Tag), typeOfRemoteElem.Field(i).Name)] = remoteRunwayElems.Field(i).Interface()
			}
		case "map[string]int":
			diffingFields := compareMapChanges(f.Interface().(map[string]int), remoteRunwayElems.Field(i).Interface().(map[string]int))

			if len(diffingFields) > 0 {
				localMap := f.Interface().(map[string]int)
				remoteMap := remoteRunwayElems.Field(i).Interface().(map[string]int)
				localDiffs := make(map[string]int)
				remoteDiffs := make(map[string]int)
				for _, field := range diffingFields {
					localDiffs[field] = localMap[field]
					remoteDiffs[field] = remoteMap[field]
				}
				diff["LOCAL"][getJSONTag(string(typeOfT.Field(i).Tag), typeOfT.Field(i).Name)] = localDiffs
				diff["REMOTE"][getJSONTag(string(typeOfT.Field(i).Tag), typeOfT.Field(i).Name)] = remoteDiffs
			}
		default:
			log.WithField("Type", f.Type().String()).Error("Type not recognized")
		}
	}

	return domain.Conflict{
		RunwayID:         localRunway.ID,
		Remote:           diff["REMOTE"],
		Local:            diff["LOCAL"],
		ResolutionMethod: "LOCAL",
	}
}

func getJSONTag(tag string, key string) string {
	if tag == "" {
		return key
	}

	jsonFullTagNameRegex := regexp.MustCompile("json:\".*?\"")
	jsonTagNameRegex := regexp.MustCompile("\".*?\"")

	match := jsonFullTagNameRegex.FindString(tag)

	if match == "" {
		return key
	}

	tagNameMatch := jsonTagNameRegex.FindStringSubmatch(match)

	if tagNameMatch == nil {
		return key
	}

	returnString := strings.ReplaceAll(tagNameMatch[0], "\"", "")

	return returnString
}

func (s SyncHandler) ApplyChanges(remoteRunway domain.Runway, conflict domain.Conflict, strategy domain.ResolutionStrategy) (domain.Runway, error) {
	switch strategy {
	case domain.APPLY_LOCAL:
		return applyObjChanges(remoteRunway, conflict.Local)
	case domain.APPLY_REMOTE:
		return applyObjChanges(remoteRunway, conflict.Remote)
	default:
		return domain.Runway{}, ErrorImpossibleStrategy
	}
}

func convertToMapInterface(rwy domain.Runway) (returnMap map[string]interface{}, err error) {
	asJSON, err := json.Marshal(rwy)

	if err != nil {
		return returnMap, err
	}

	err = json.Unmarshal(asJSON, &returnMap)

	if err != nil {
		return returnMap, err
	}

	return returnMap, nil
}

func convertToStruct(m map[string]interface{}) (domain.Runway, error) {
	data, _ := json.Marshal(m)
	var result domain.Runway
	err := json.Unmarshal(data, &result)
	return result, err
}

func applyObjChanges(rwy domain.Runway, conflictObj map[string]interface{}) (domain.Runway, error) {
	rwyMap, err := convertToMapInterface(rwy)

	if err != nil {
		return domain.Runway{}, err
	}

	for conflictObjKey, conflictObjVal := range conflictObj {
		// Check if key exists
		// Or if key is latest version, do not merge
		if _, keyExists := rwyMap[conflictObjKey]; !keyExists || conflictObjKey == "LatestVersion" {
			continue
		}

		if isLoopable(conflictObjVal) {
			switch reflect.ValueOf(conflictObjVal).Kind() {
			case reflect.Map:
				for _, key := range reflect.ValueOf(conflictObjVal).MapKeys() {
					strct := reflect.ValueOf(conflictObjVal).MapIndex(key)
					rwyMap[conflictObjKey].(map[string]interface{})[key.String()] = strct.Interface()
				}
			default:
				log.Error("Value is loopable but not ready to be handled")
			}
			continue
		}

		switch reflect.ValueOf(conflictObjVal).Kind() {
		case reflect.Int, reflect.Float64, reflect.Bool:
			rwyMap[conflictObjKey] = conflictObjVal
		default:
			log.WithField("Type", reflect.ValueOf(conflictObjVal).Kind()).Error("Cannot set value of type")
		}

	}

	return convertToStruct(rwyMap)
}

func compareMapChanges(local map[string]int, remote map[string]int) []string {
	returnVal := []string{}
	for key := range local {
		if local[key] != remote[key] {
			returnVal = append(returnVal, key)
			break
		}
	}
	return returnVal
}

func compareSingleValue(local interface{}, remote interface{}) bool {
	returnVal := false
	if unpackPointerVal(local) != unpackPointerVal(remote) {
		return true
	}
	return returnVal
}

func unpackPointerVal(v interface{}) interface{} {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		return &v
	}
	return v
}

func isLoopable(v interface{}) (res bool) {
	defer func() {
		if recover() != nil {
			res = false
		}
	}()
	reflect.ValueOf(v).Len()
	return true
}
