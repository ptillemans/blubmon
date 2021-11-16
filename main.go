package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os/exec"
	"strings"
	"time"
)

var logger *log.Logger

func IsOnline(service string) (bool, error) {
	out, err := exec.Command("/usr/bin/connmanctl", "services", service).Output()
	if err != nil {
		return false, err
	}
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if !strings.Contains(line, " = ") {
			continue
		}
		parts := strings.SplitN(line, " = ", 2)
		if parts[0] != "State" {
			continue
		}
		return parts[1] == "online", nil
	}

	return false, nil

}

func logPrintLines(text string) {
	for _, line := range strings.Split(text, "\n") {
		logger.Print(line)
	}
}

func DisconnectWifi(service string) error {
	out, err := exec.Command("/usr/bin/connmanctl", "disconnect", service).Output()
	if err != nil {
		return err
	}
	logPrintLines(string(out))
	return nil
}

func ConnectWifi(service string) error {
	out, err := exec.Command("/usr/bin/connmanctl", "connect", service).Output()
	if err != nil {
		return err
	}
	logPrintLines(string(out))
	return nil
}

func main() {
	var err error
	logger, err = syslog.NewLogger(syslog.LOG_WARNING|syslog.LOG_DAEMON,
		0)
	if err != nil {
		fmt.Printf("Error starting loggerer: ", err)
	}

	service := "wifi_001986811bff_536e616d656c6c697437_managed_psk"
	online, err := IsOnline(service)
	if err != nil {
		logger.Fatal(err)
	}
	if online {
		fmt.Println("Online")
		logger.Print("Online")
	} else {
		fmt.Println("Offline")
		logger.Print("Offline")

		err = DisconnectWifi(service)
		if err != nil {
			logger.Print("Disconnect Wifi:", err)
		}

		time.Sleep(15)
		err = ConnectWifi(service)
		if err != nil {
			logger.Print("Connect Wifi:", err)
		}

	}
}
