package main

import (
	"fmt"
	"log"
	"openprio_proxy/openprio_pt_position_data"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/protobuf/proto"
)

type MQTTConfig struct {
	MQTTHost     string `envconfig:"MQTT_HOST" required:"true"`
	MQTTUsername string `envconfig:"MQTT_USERNAME" required:"true"`
	MQTTPassword string `envconfig:"MQTT_PASSWORD" required:"true"`
}

func createClientOptions(cfg MQTTConfig) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	log.Print(cfg.MQTTHost)
	opts.AddBroker(fmt.Sprintf("ssl://%s", cfg.MQTTHost))
	opts.SetUsername(cfg.MQTTUsername)
	// opts.SetClientID(cfg.MQTTUsername)
	opts.SetPassword(cfg.MQTTPassword)
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(onConnectionLostHandler)
	opts.SetReconnectingHandler(onReConnect)
	return opts
}

func startMQTTPublisher(c <-chan openprio_pt_position_data.LocationMessage) {
	var cfg MQTTConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Loading MQTTConfig went wrong: %s", err)
	}

	mqttOpts := createClientOptions(cfg)
	client := mqtt.NewClient(mqttOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	counter := 0
	lastTimeReset := time.Now()
	for {
		data := <-c
		topic := fmt.Sprintf("/prod/pt/position/%s/vehicle_number/%d",
			data.VehicleDescriptor.GetDataOwnerCode(), data.VehicleDescriptor.GetVehicleNumber())
		out, err := proto.Marshal(&data)
		if err != nil {
			log.Fatalln("Failed to encode location message:", err)
		}

		t := client.Publish(topic, 0, false, out)
		go func() {
			_ = t.Wait()
			if t.Error() != nil {
				log.Print(t.Error())
			}
		}()
		counter = counter + 1
		if time.Now().Sub(lastTimeReset) >= 1*time.Minute {
			log.Printf("%d messages sent over MQTT since previous reset (normally 1 minute ago)", counter)
			lastTimeReset = time.Now()
			counter = 0
		}

	}

}

func onReConnect(client mqtt.Client, c *mqtt.ClientOptions) {
	log.Println("MQTT Tries to reconnect.")
}

func onConnectionLostHandler(client mqtt.Client, err error) {
	log.Println("MQTT lost connection")
	log.Println(err)
}
