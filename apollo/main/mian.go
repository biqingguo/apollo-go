package main

import (
	"apollo"
	"encoding/json"
	"log"
)

func main() {
	readyConfig := &agollo.AppConfig{
		AppId:         "10000",
		Cluster:       "diamond",
		NamespaceName: "application",
		Ip:            "localhost:8080",
	}

	agollo.InitCustomConfig(func() (*agollo.AppConfig, error) {
		return readyConfig, nil
	})

	agollo.Start()
	value := agollo.GetStringValue("batch", "2")
	log.Printf("key02: %s", value)
	event := agollo.ListenChangeEvent()
	changeEvent := <-event
	bytes, _ := json.Marshal(changeEvent)
	log.Printf("key02: %s", string(bytes))
	//fmt.Println("event:", string(bytes))
}
