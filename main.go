package main

import "fmt"

func main() {
    fmt.Println("Initialize Room Temperature Monitor")

    // initialize mqtt client
    initializeClient()

    // subscribe to temperature readings
    subscribeToTemperature()

    // block forever
    select{ }

    // shouldn't reach this line
    defer deinitializeClient()
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

    // y = ax + b
    // a = -2.857142857142857
    // b = 100
    return int(-2.857142857142857 * roomTemp + 100)
}
