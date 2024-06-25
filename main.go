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

var lastSavedPositions = map[string]*openprio_pt_position_data.LocationMessage{}

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
		if previousLocationMessage, ok := lastSavedPositions[key]; ok && result.Position.GetBearing() < 0 {
			// previousPosition
			previousPosition := previousLocationMessage.Position
			previousDrivingDirection := previousLocationMessage.VehicleDescriptor.DrivingDirection

			// newPosition
			newPosition := *result.GetPosition()
			newDrivingDirection := result.VehicleDescriptor.DrivingDirection
			distance := flatEarthDistance(previousPosition.Latitude, previousPosition.Longitude, newPosition.Latitude, newPosition.Longitude)

			// Turng bearing with 180deg if tram turns around and a bearing was set in previous locationMessage.
			if previousPosition.Bearing >= 0.0 && newDrivingDirection != openprio_pt_position_data.DrivingDirection_UNDEFINED && previousDrivingDirection != openprio_pt_position_data.DrivingDirection_UNDEFINED && newDrivingDirection != previousDrivingDirection {
				newBearing := (previousPosition.Bearing + 180.0)
				if newBearing >= 360.0 {
					newBearing = newBearing - 360.0
				}
				result.Position.Bearing = newBearing
				lastSavedPositions[key] = &result
			} else if result.GetPosition().GetSpeed() >= 1.0 && distance > 150 {
				// Set bearing to unknown when position changes with >150m
				result.Position.Bearing = -1
				lastSavedPositions[key] = &result
			} else if result.GetPosition().GetSpeed() >= 1.0 && ((newPosition.Hdop <= 1.0 && distance >= 1.5) || distance >= 5.0) {
				// Update bearing when tram has some speed and moved over a certain distance depending on the quality of the GPS signal.
				result.Position.Bearing = BearingBetweenPositions(*previousPosition, newPosition, result.Position.Hdop)
				lastSavedPositions[key] = &result
			} else {
				// Reuse earlier bearing
				result.Position.Bearing = previousPosition.Bearing
			}
		} else {
			lastSavedPositions[key] = &result
		}

		if err != nil {
			log.Print(err)
		} else {
			c <- result
		}
	}
}
