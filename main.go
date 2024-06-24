package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"openprio_proxy/openprio_pt_position_data"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	c := make(chan openprio_pt_position_data.LocationMessage)
	go startMQTTPublisher(c)
	startTCPServer(c)
}

// TCPConfig is the configuration of the TCP server.
type TCPConfig struct {
	DataFormat    string `envconfig:"DATA_FORMAT_OF_SOURCE" default:"csv"`
	ListenAddress string `envconfig:"TCP_LISTEN_ADDRESS" default:"127.0.0.1:9001"`
}

// start a new TCPServer.
func startTCPServer(c chan<- openprio_pt_position_data.LocationMessage) {
	var cfg TCPConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal("Loading TCPConfig went wrong:", err)
	}

	ln, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		log.Fatal("Starting tcp server went wrong: ", err)
	}
	log.Printf("Started tcp server on %s", cfg.ListenAddress)
	log.Print(cfg.ListenAddress)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("Something went wrong when accepting connection:", err)
		}
		go handleTCPConnection(conn, cfg.DataFormat, c)
	}
}

var lastSavedPositions = map[string]*openprio_pt_position_data.Position{}

func handleTCPConnection(conn net.Conn, dataFormat string, c chan<- openprio_pt_position_data.LocationMessage) {
	timeoutDuration := 300 * time.Second
	bufReader := bufio.NewReader(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			log.Print("Something went wrong while reading data from TCP", err)
			return
		}
		result, err := processMessage(bytes, dataFormat)

		// Calculate bearing based on previous position.
		key := fmt.Sprintf("%s:%d", result.VehicleDescriptor.GetDataOwnerCode(), result.VehicleDescriptor.GetVehicleNumber())
		if previousPosition, ok := lastSavedPositions[key]; ok && result.Position.GetBearing() < 0 {
			newPosition := *result.GetPosition()
			// update bearing when vehicle runs at least 1 m/s and when the distance since the previous point is at least 1.5m
			if result.GetPosition().GetSpeed() >= 1.0 && flatEarthDistance(previousPosition.Latitude, previousPosition.Longitude, newPosition.Latitude, newPosition.Longitude) >= 1.5 {
				result.Position.Bearing = BearingBetweenPositions(*previousPosition, newPosition, result.Position.Hdop)
				lastSavedPositions[key] = result.Position
			} else {
				result.Position.Bearing = previousPosition.Bearing
			}
		} else {
			lastSavedPositions[key] = result.Position
		}

		if err != nil {
			log.Print(err)
		} else {
			c <- result
		}
	}
}
