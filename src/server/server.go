package server
import(
	
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"strconv"
	"fmt"

)

type Server struct{
	Broadcast chan string
	Clients map[*websocket.Conn]bool
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
// pingConnectedClient put the message on broadcast channel then a function handleMessages will invoke everytime that will Send the message to all connected client
func (s *Server)PingConnectedClient(msg string){
	fmt.Println("Server pings after every 5 sec")
	s.Broadcast <- msg
}

func (s *Server)StartServer(port string ){
	fmt.Println("server starting ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server)HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-s.Broadcast
		// Send it out to every client that is currently connected
		for conn,val := range s.Clients {
			if s.Clients[conn] {
				log.Printf("msg: %v %v", val,msg)
				if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
					log.Printf("sending message to client: %v",msg)
				}else{
					log.Printf("error in sending message to client: %v",err)
					delete(s.Clients, conn)
				}
			}
		}
	}
}


func (s *Server)HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial request to a websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		messageType, d, err := ws.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}
		
		if err := ws.WriteMessage(messageType, d); err != nil {
			log.Println(err)
			return
		}
		log.Println(string(d))
		if !s.Clients[ws] {
			log.Println("New Client Connected")
			s.Clients[ws] = true
		}else{
			log.Println("Existing Client Connected")
		}
	}
}

// activeClient function is for getting the resutls for connected clients
func (s *Server)ActiveClient(w http.ResponseWriter, r *http.Request) {
	cnt:=0
	for conn,_ := range s.Clients {
		if s.Clients[conn] {
			cnt=cnt+1
		}
	}
	activeClients :="Active clients are "+strconv.Itoa(cnt)
	fmt.Println(activeClients)
    fmt.Fprintf(w,activeClients)
}