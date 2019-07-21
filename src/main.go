package main

import (
    "fmt"
	"github.com/gorilla/websocket"
	"routes"
	"cronjob"
	"server"
)


// var broadcast = make(chan string)           // broadcast channel
// clients map[*websocket.Conn]bool // connected clients
// We'll need to define an Upgrader
// this will require a Read and Write buffer size


// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  1024,
//     WriteBufferSize: 1024,
// }


// func homePage(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Home Page")
// }

// // handleConnections function handles request coming from client
// // if its a new client the connection is stored in map clients

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// 		// Upgrade initial request to a websocket
// 		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
// 		ws, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
		
// 		for {
// 			messageType, d, err := ws.ReadMessage()
// 			        if err != nil {
// 			            log.Println(err)
// 			            return
// 					}
			
// 			if err := ws.WriteMessage(messageType, d); err != nil {
// 				log.Println(err)
// 				return
// 			}
// 			log.Println(string(d))
// 			if !clients[ws] {
// 				log.Println("New Client Connected")
// 				clients[ws] = true
// 			}else{
// 				log.Println("Existing Client Connected")
// 			}
// 		}
// 	}

// // handleMessages is for sending message "ping" to the connected clients if it's connected
// // if there is any problem in connecting with clients if will remove the map clients 	

// func handleMessages() {
// 	for {
// 		// Grab the next message from the broadcast channel
// 		msg := <-broadcast
// 		// Send it out to every client that is currently connected
// 		for conn,val := range clients {
// 			if clients[conn] {
// 				log.Printf("msg: %v %v", val,msg)
// 				if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
// 					log.Printf("sending message to client: %v",msg)
// 				}else{
// 					log.Printf("error in sending message to client: %v",err)
// 								delete(clients, conn)
// 				}
// 			}
// 		}
// 	}
// }



// // func setupRoutes() {
// //     http.HandleFunc("/", homePage)
// // 	http.HandleFunc("/connect", handleConnections)
// // 	http.HandleFunc("/activeclient", activeClient)
	
// // }


// // activeClient function is for getting the resutls for connected clients
// func activeClient(w http.ResponseWriter, r *http.Request) {
// 	cnt:=0
// 	for conn,_ := range clients {
// 		if clients[conn] {
// 			cnt=cnt+1
// 		}
// 	}
// 	activeClients :="Active clients are "+strconv.Itoa(cnt)
// 	fmt.Println(activeClients)
//     fmt.Fprintf(w,activeClients)
// }



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
