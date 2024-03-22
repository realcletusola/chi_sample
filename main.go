package main  

// import packages 
import (

	"net/http"
	"log"

	"github.com/cletushunsu/chi_sample/Router"
)


// main function 
func main() {

	// get the router from routes.go 
	r := routes.GoChiRouter()

	// start the server 
	log.Fatal(http.ListenAndServe(":8090", r))
}

