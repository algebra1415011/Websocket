package homepage


import (
    "fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home Page")
}