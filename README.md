# Room Temperature Control over MQTT

This is a simple demonstration on how to use MQTT to monitor and control room temperature.

The room temperature is reported periodically over the `/readings/temperature` topic and the heating valve openness is published to the `/actuators/room-1` topic.

## Dependencies

```
go get github.com/eclipse/paho.mqtt.golang
go get golang.org/x/net/websocket
go get golang.org/x/net/proxy
```

## How to simulate the system

- You will need an MQTT broker. One way to achieve that is by running `mosquitto` locally on your machine.
- Having the MQTT broker running you can then lunch the application:
```
go build
./room_temperature_mqtt tcp://localhost:1883
```
- You can simulate temperature readings by publishing messages to the 
`/readings/temperature` topic
```
mosquitto_pub -t "/readings/temperature" -m '{"sensorID": "sensor-1","type": 
"temperature","value": 15}' -q 1 -r
```

### Temperature control logic
- In this example is assumed that there is a linear relationship between the temperature and the valve openness. In the real world to achieve accurate results this relation needs to be tuned and calibrated for the specific environment.
- It is assumed that temperature ranges between 0°C and 35°C, for 0°C the valve is set to 100% openness and for 35°C the valve is set to 0% openness.