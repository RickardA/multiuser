package service

import (
	"fmt"
	"os"
	"reflect"

	conflictObj "github.com/RickardA/multiuser/pkg/repository/conflict_obj"

	"github.com/google/uuid"

	"github.com/RickardA/multiuser/pkg/aggregate"
	"github.com/RickardA/multiuser/pkg/repository/runway"
)

type SyncHandlerService interface {
	New(db runway.RunwayRepository, conflictDB conflictObj.ConflictObjRepository) (SyncHandler, error)
	CheckVersionMismatch(localRunway aggregate.Runway) (bool, error)
	GetConflictingFields(localRunway aggregate.Runway)
}

type SyncHandler struct {
	db         runway.RunwayRepository
	conflictDB conflictObj.ConflictObjRepository
}

func New(db runway.RunwayRepository, conflictDB conflictObj.ConflictObjRepository) (SyncHandler, error) {
	return SyncHandler{
		db:         db,
		conflictDB: conflictDB,
	}, nil
}

func (s SyncHandler) CheckVersionMismatch(localRunway aggregate.Runway) (bool, error) {
	remoteRunway, err := s.db.GetByDesignator(localRunway.Designator)

	if err != nil {
		return false, err
	}

	if remoteRunway.LatestVersion == localRunway.LatestVersion {
		return false, nil
	}

	return true, nil
}

func (s SyncHandler) GetConflictingFields(localRunway aggregate.Runway) aggregate.ConflictObj {
	remoteRunway, err := s.db.GetByDesignator(localRunway.Designator)

	if err != nil {
		os.Exit(1)
	}

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
				diff["LOCAL"][typeOfT.Field(i).Name] = f.Interface()
				diff["REMOTE"][typeOfRemoteElem.Field(i).Name] = remoteRunwayElems.Field(i).Interface()
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
				diff["LOCAL"][typeOfT.Field(i).Name] = localDiffs
				diff["REMOTE"][typeOfT.Field(i).Name] = remoteDiffs
			}
		default:
			fmt.Printf("Type not recognized %v\n", f.Type().String())
		}
	}

	return aggregate.ConflictObj{
		ID:               uuid.New(),
		Remote:           diff["REMOTE"],
		Local:            diff["LOCAL"],
		ResolutionMethod: "LOCAL",
	}
}

func (s SyncHandler) applyChanges(conflictID uuid.UUID, strategy string) {
	fmt.Println("Apply Changes")
	remoteRunway, err := s.db.GetByDesignator("10-23")

	fmt.Printf("Remote runway before %v\n", remoteRunway)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conflictObj, err := s.conflictDB.GetByID(conflictID)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch strategy {
	case "LOCAL":
		elems := reflect.ValueOf(&remoteRunway).Elem()
		//typesOfRemoteRunway := elems.Type()
		for key, val := range conflictObj.Local {
			fmt.Printf("This is key %v\n", key)
			fmt.Printf("This is val %v\n", val)
			f := elems.FieldByName(key)

			if f.IsValid() && f.CanSet() {
				if isLoopable(f.Interface()) {
					switch f.Kind() {
					case reflect.Map:
						for key, val := range f.Interface.(map[string]interface{}) {

						}
					default:
						fmt.Printf("Value is loopable but not ready to be handled\n")
					}
					fmt.Printf("Value is loopable, kind: %v\n", f.Kind())
					continue
				}

				switch f.Kind() {
				case reflect.Int:
					f.Set(reflect.ValueOf(val.(int)))
				case reflect.Bool:
					f.SetBool(val.(bool))
				default:
					fmt.Printf("Cannot set val of type %v\n", f.Kind())
				}
			}

			/*for i := 0; i < elems.NumField(); i++ {
				// f := elems.Field(i)
				if typesOfRemoteRunway.Field(i).Name == key {
					fmt.Printf("Match on field %v with field %v \n", key, typesOfRemoteRunway.Field(i).Name)
					if isLoopable(elems.Field(i).Interface()) {
						fmt.Println("Value is loopable")
					} else {
						fmt.Println("Value is NOT loopable")
						f.s
					}

				}
			}*/
		}

		//s.db.Update(remoteRunway)

		fmt.Printf("Remote runway after %v\n", remoteRunway)

	case "REMOTE":
	default:
		fmt.Println("Impossible strategy")
		os.Exit(1)
	}

	fmt.Printf("DB runway: %v\n", remoteRunway)
	fmt.Printf("DB Conflcit: %v\n", conflictObj)
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
