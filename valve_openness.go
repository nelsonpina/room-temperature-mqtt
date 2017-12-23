package main

import (
    "encoding/json"
    "fmt"
)

type ValveOpenness struct {
    Level int `json:"level"`
}

func publishToValve(message string) {
    publishToTopic("/actuators/room-1", message)
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
