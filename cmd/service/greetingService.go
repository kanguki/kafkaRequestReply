package httpService

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hi there,  %s!", vars["name"])
}
