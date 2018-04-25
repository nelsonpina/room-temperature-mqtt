package main

import "encoding/json"

type InputTemperature struct {
    SensorID string  `json:"sensorID"`
    Type     string  `json:"type"`
    Value    float32 `json:"value"`
}

const temperatureReadingsTopic = "/readings/temperature"

func subscribeToTemperature() {
    subscribeToTopic(temperatureReadingsTopic)
}

func decodeTemperatureFromJson(tempJson []byte) float32 {
    tempStruct := InputTemperature{}
    json.Unmarshal(tempJson, &tempStruct)
    return tempStruct.Value
}
