package sync_handler

/*type SyncHandlerService interface {
	New(db runway.RunwayRepository, conflictDB conflict_repository.ConflictObjRepository) (SyncHandler, error)
	CheckVersionMismatch(localRunway domain.Runway) (bool, error)
	GetConflictingFields(localRunway domain.Runway)
}

type SyncHandler struct {
	db         runway.RunwayRepository
	conflictDB conflict_repository.ConflictObjRepository
}

func New(db runway.RunwayRepository, conflictDB conflict_repository.ConflictObjRepository) (SyncHandler, error) {
	return SyncHandler{
		db:         db,
		conflictDB: conflictDB,
	}, nil
}

func (s SyncHandler) CheckVersionMismatch(localRunway domain.Runway) (bool, error) {
	remoteRunway, err := s.db.GetByDesignator(localRunway.Designator)

	if err != nil {
		return false, err
	}

	if remoteRunway.LatestVersion == localRunway.LatestVersion {
		return false, nil
	}

	return true, nil
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
			fmt.Printf("Type not recognized %v\n", f.Type().String())
		}
	}

	return domain.Conflict{
		//ID:               uuid.New(),
		Remote:           diff["REMOTE"],
		Local:            diff["LOCAL"],
		ResolutionMethod: "LOCAL",
	}
}

func getJSONTag(tag string, key string) string {
	fmt.Printf("Tag %v\n", tag)

	if tag == "" {
		return key
	}

	jsonFullTagNameRegex := regexp.MustCompile("json:\".*?\"")
	jsonTagNameRegex := regexp.MustCompile("\".*?\"")

	match := jsonFullTagNameRegex.FindString(tag)

	fmt.Printf("Full match %v\n", match)

	if match == "" {
		return key
	}

	tagNameMatch := jsonTagNameRegex.FindStringSubmatch(match)

	if tagNameMatch == nil {
		return key
	}

	fmt.Printf("Return tag %v\n", tagNameMatch[0])
	returnString := strings.ReplaceAll(tagNameMatch[0], "\"", "")

	return returnString
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
		applyObjChanges(remoteRunway, conflictObj.Local)
		fmt.Printf("Remote runway after %v\n", remoteRunway)
	case "REMOTE":
		applyObjChanges(remoteRunway, conflictObj.Remote)
		fmt.Printf("Remote runway after %v\n", remoteRunway)
	default:
		fmt.Println("Impossible strategy")
		os.Exit(1)
	}

	fmt.Printf("DB runway: %v\n", remoteRunway)
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

func applyObjChanges(rwy domain.Runway, conflictObj map[string]interface{}) {
	rwyMap, err := convertToMapInterface(rwy)

	if err != nil {
		panic("Something went wrong!!!")
	}

	fmt.Printf("This is rwy map: %v\n", rwyMap)

	for conflictObjKey, conflictObjVal := range conflictObj {
		if _, keyExists := rwyMap[conflictObjKey]; !keyExists {
			fmt.Printf("Key does not exist: %v\n", conflictObjKey)
			continue
		}

		if isLoopable(conflictObjVal) {
			switch reflect.ValueOf(conflictObjVal).Kind() {
			case reflect.Map:
				fmt.Println("Value is a map")
				for _, key := range reflect.ValueOf(conflictObjVal).MapKeys() {
					strct := reflect.ValueOf(conflictObjVal).MapIndex(key)
					rwyMap[conflictObjKey].(map[string]interface{})[key.String()] = strct.Interface()
				}
			default:
				fmt.Printf("Value is loopable but not ready to be handled\n")
			}
			continue
		}

		switch reflect.ValueOf(conflictObjVal).Kind() {
		case reflect.Int, reflect.Float64, reflect.Bool:
			rwyMap[conflictObjKey] = conflictObjVal
		default:
			fmt.Printf("Cannot set val of type %v\n", reflect.ValueOf(conflictObjVal).Kind())
		}

	}
	fmt.Printf("This i rwy map after %v\n", rwyMap)
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
}*/
