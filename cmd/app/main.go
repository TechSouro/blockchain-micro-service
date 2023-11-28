package main

import (
	"context"
	"log"
	service "service/config"
)

func main() {
	log.Println("Starting blockchain monitor...")

	service.InitEventSystem()

	<-context.Background().Done()
}
