package httpService

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"mo.io/reqRep/pkg/kafka"
)

func Produce(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msg := vars["message"]
	kafka.Produce(msg)
	fmt.Fprintf(w, "Kafka produced msg %s!", msg)
}
