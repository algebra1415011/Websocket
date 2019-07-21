package routes
import(
	"net/http"
	"homepage"
	"server"
	
)

func SetupRoutes(servertemp *server.Server) {
    http.HandleFunc("/", homepage.HomePage)
	http.HandleFunc("/connect", servertemp.HandleConnections)
	http.HandleFunc("/activeclient", servertemp.ActiveClient)
	
}
