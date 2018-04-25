package main

import (
    "encoding/json"
    "fmt"
)

type ValveOpenness struct {
    Level int `json:"level"`
}

const valveControlTopic = "/actuators/room-1"

func publishToValve(message string) {
    publishToTopic(valveControlTopic, message)
}

func encodeValveIntoJson(valveOpenness int) string {
    valveStruct := ValveOpenness{}
    valveStruct.Level = valveOpenness
    valveJson, _ := json.Marshal(valveStruct)
    return string(valveJson)
}

func dispachValveOpenness(valveOpenness int) {
    fmt.Printf("Set Valve Openness: %d\n", valveOpenness)
    publishToValve(encodeValveIntoJson(valveOpenness))
}
