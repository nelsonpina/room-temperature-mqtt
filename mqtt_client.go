package main

import (
    "github.com/eclipse/paho.mqtt.golang"
    "fmt"
    "os"
    "strings"
)

// mqtt client
const server   = "tcp://localhost:1883"
const clientID = "room-1-temp"

// mqtt variables
var mqttOptions *mqtt.ClientOptions
var mqttClient mqtt.Client

// publish messages handler
var publishHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    // handle temperature message
    if strings.Compare(msg.Topic(), temperatureReadingsTopic) == 0 {
        inputTemp := decodeTemperatureFromJson(msg.Payload())
        handleRoomTemperature(inputTemp)
    }
}

func initializeClient() {
    fmt.Println("Initialize MQTT Client")
    mqttOptions = mqtt.NewClientOptions().AddBroker(server).SetClientID(
        clientID)
    mqttOptions.SetDefaultPublishHandler(publishHandler)

    mqttClient = mqtt.NewClient(mqttOptions)
    if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
}

func deinitializeClient() {
    fmt.Println("Disconnect MQTT Client")
    mqttClient.Disconnect(250)
}

func subscribeToTopic(topic string) {
    fmt.Printf("Subscribe to Topic: %s\n", topic)
    if token := mqttClient.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
        os.Exit(1)
    }
}

func publishToTopic(topic string, pubMessage string) {
    fmt.Printf("Publish to Topic: %s\n", topic)
    token := mqttClient.Publish(topic, 0, false, pubMessage)
    token.Wait()
}
