package router

import (
	"github.com/gorilla/mux"
	"mo.io/reqRep/cmd/service"
)

func GetHandler() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/{name}", httpService.Greeting)
	mux.HandleFunc("/kafka/{message}", httpService.Produce)
	return mux
}
