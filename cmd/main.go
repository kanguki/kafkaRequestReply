package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"mo.io/kafkaReqRep/cmd/router"
	"mo.io/kafkaReqRep/pkg/kafka"
)

func main() {
	kafka.InitKafka()
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	port := os.Getenv("PORT")
	server := http.TimeoutHandler(router.GetHandler(), time.Second*3, "Time out!")
	fmt.Printf("Run server on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}
