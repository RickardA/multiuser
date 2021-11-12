package service

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/RickardA/multiuser/pkg/aggregate"
	"github.com/RickardA/multiuser/pkg/repository/runway"
)

type SyncHandlerService interface {
	New(db runway.RunwayRepository) (SyncHandler, error)
	CheckVersionMismatch(localRunway aggregate.Runway) (bool, error)
	GetConflictingFields(localRunway aggregate.Runway)
}

type SyncHandler struct {
	db runway.RunwayRepository
}

func New(db runway.RunwayRepository) (SyncHandler, error) {
	return SyncHandler{
		db: db,
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

func (s SyncHandler) GetConflictingFields(localRunway aggregate.Runway) {
	remoteRunway, err := s.db.GetByDesignator(localRunway.Designator)

	if err != nil {
		os.Exit(1)
	}

	/*if err != nil {
		fmt.Printf("Error: %v", err)
	}

	if err != nil {
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}*/

	localRunwayElems := reflect.ValueOf(&localRunway).Elem()
	remoteRunwayElems := reflect.ValueOf(&remoteRunway).Elem()

	typeOfT := localRunwayElems.Type()
	typeOfRemoteElem := remoteRunwayElems.Type()

	diff := make(map[string]map[string]interface{})
	diff["LOCAL"] = make(map[string]interface{})
	diff["REMOTE"] = make(map[string]interface{})

	for i := 0; i < localRunwayElems.NumField(); i++ {
		f := localRunwayElems.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfRemoteElem.Field(i).Name, remoteRunwayElems.Field(i).Type(), remoteRunwayElems.Field(i).Interface())
		//fmt.Printf("Is Value Equal: %v\n", f.Interface() == remoteRunwayElems.Field(i).Interface())
		switch f.Type().String() {
		case "bool", "string", "int":
			fmt.Println("Got easy type")
			fmt.Printf("Easy type differs: %v\n", compareSingleValue(f.Interface(), remoteRunwayElems.Field(i).Interface()))
			if compareSingleValue(f.Interface(), remoteRunwayElems.Field(i).Interface()) {
				diff["LOCAL"][typeOfT.Field(i).Name] = f.Interface()
				diff["REMOTE"][typeOfRemoteElem.Field(i).Name] = remoteRunwayElems.Field(i).Interface()
			}
		case "map[string]int":
			fmt.Println("Got map type")
			fmt.Printf("Map differs: %v\n", compareMapChanges(f.Interface().(map[string]int), remoteRunwayElems.Field(i).Interface().(map[string]int)))
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

	t, _ := json.Marshal(diff)
	fmt.Printf("This is diff")
	fmt.Printf("%v", string(t))
}

func compareMapChanges(local map[string]int, remote map[string]int) []string {
	returnVal := []string{}
	for key := range local {
		fmt.Printf("Local val: %v\n", local[key])
		fmt.Printf("Remote val: %v\n", remote[key])
		if local[key] != remote[key] {
			returnVal = append(returnVal, key)
			break
		}
	}
	return returnVal
}

func compareSingleValue(local interface{}, remote interface{}) bool {
	returnVal := false
	fmt.Printf("Local val: %v\n", unpackPointerVal(local))
	fmt.Printf("Remote val: %v\n", unpackPointerVal(remote))
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
