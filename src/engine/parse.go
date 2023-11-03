/*
 This file parsers JSON file fed into the program and creates
 an instance of a "struct" to store the l-system configuraiton
*/

package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
)

type LConfig struct {
	iteration int32
	rules     map[string]string
	cmds      map[string]string
	axiom     string
	magintude int32
	angle     int32
	ix        int32
	iy        int32
	width     int32
	height    int32
}

// Reads a field safely from a map, returns an empty string if key is not found
func GetSafe(data map[string]interface{}, key string) interface{} {
	val, ok := data[key]
	if ok {
		return val
	}
	return ""
}

// Reads int32 safely from data, and stores into a memory location pointed by 'intValRef'
func ReadInt(intValRef *int32, data *map[string]interface{}, intField string, errMessage string) {
	intVal, err := strconv.ParseInt(GetSafe(*data, intField).(string), 10, 32)
	if err != nil {
		log.Fatalf("[ERROR]: %s", errMessage)
	}
	*intValRef = int32(intVal)
}

func ReadMap(mapRef *map[string]string, data *map[string]interface{}, field string) {
	datamap := GetSafe(*data, field)
	switch datamap.(type) {
	case string: //string means there is no map data type found in file
		log.Fatalf("[ERROR]: key '%s' not found! It is required!", field)
	}
	dmap, ok := datamap.(map[string]string)
	if !ok {
		log.Fatalf("[ERROR]: couldn't process field %s", field)
	}
	*mapRef = dmap
}

// Constructs LConfig from 'data' map
func CreateLConfig(data map[string]interface{}) LConfig {
	config := LConfig{}

	//Read String
	config.axiom = GetSafe(data, "axiom").(string)

	//Read ints
	ReadInt(&config.magintude, &data, "magnitude", "Cannot read 'magnitude' from json")
	ReadInt(&config.ix, &data, "ix", "Cannot read 'ix' from json")
	ReadInt(&config.iy, &data, "iy", "Cannot read 'iy' from json")
	ReadInt(&config.width, &data, "width", "Cannot read 'width' from json")
	ReadInt(&config.height, &data, "height", "Cannot read 'height' from json")
	ReadInt(&config.angle, &data, "angle", "Cannot read 'angle' from json")
	ReadInt(&config.iteration, &data, "iteration", "Cannot read 'iteration' from json")

	//Read maps
	ReadMap(&config.cmds, &data, "cmds")
	ReadMap(&config.rules, &data, "rules")

	return config
}

func Parse(filename string) LConfig {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("[ERROR] opening "+filename, ": ", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("[ERROR] reading "+filename, ": ", err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return CreateLConfig(result)
}
