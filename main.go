package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("Initialize Room Temperature Monitor")

    // initialize mqtt client
    if mqttServer := getServerName(os.Args); mqttServer != "" {
        initializeClient(mqttServer)
    } else {
        return
    }

    // subscribe to temperature readings
    subscribeToTemperature()

    // block forever
    select{ }

    // shouldn't reach this line
    defer deinitializeClient()
}

func getServerName (input []string) string {
    const usageMsg = "Usage: " +
        "`./room_temperature_mqtt tcp://mqtt-broker-address:1883`"

    if len(input) != 2 {
        fmt.Println("Invalid number of inputs")
        fmt.Println(usageMsg)
        return ""
    }

    serverName := input[1]
    if !strings.Contains(serverName, "tcp://") {
        fmt.Println("Missing valid network protocol")
        fmt.Println(usageMsg)
        return ""
    }
    if !strings.Contains(serverName, ":1883") {
        fmt.Println("Missing MQTT valid port")
        fmt.Println(usageMsg)
        return ""
    }
    return serverName
}

// room temperature control settings
const maxTemperature   = 35
const minTemperature   = 0
const maxValveOpenness = 100
const minValveOpenness = 0

// handle room temperature readings
func handleRoomTemperature(temperature float32) {
    fmt.Printf("Handle Temperature: %g\n", temperature)

    newValveOpenness := simpleValveControl(temperature)
    if newValveOpenness >= minTemperature || newValveOpenness <= maxTemperature {
        dispachValveOpenness(newValveOpenness)
    }
}

// simple valve control
// assuming linear relation between temperature and valve openness
func simpleValveControl (roomTemp float32) int {
    if roomTemp < minTemperature {
        return maxValveOpenness
    }
    if roomTemp > maxTemperature {
        return minValveOpenness
    }

    // y = mx + b
    // m = -2.857142857142857
    // b = 100
    return int(-2.857142857142857 * roomTemp + 100)
}
