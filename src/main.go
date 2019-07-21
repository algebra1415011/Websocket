package main

import (
    "fmt"
	"github.com/gorilla/websocket"
	"routes"
	"cronjob"
	"server"
)



func main() {
	fmt.Println("Hello World")
	
	
	var servertemp= server.Server{
		Broadcast:   make(chan string),
		Clients: 	 make(map[*websocket.Conn]bool),
	}
	routes.SetupRoutes(&servertemp)
	go servertemp.HandleMessages()
	cronjob.StartcronJob(&servertemp,"PING!")
	servertemp.StartServer(":8080")
	//cron job that executes after every 30sec
	
	routes.SetupRoutes(&servertemp)
	// handleMessages is a lightweight thread that will invoke everytime we put contect in the broadcast channel
	
	
	// cronJob := cron.New()
	// cronJob.AddFunc("*/30 * * * * *", func() {  server.pingConnectedClient() })	
	// cronJob.Start()
	
}
