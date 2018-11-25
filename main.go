package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Payload []Documents `json:"payload"`
}

type Documents struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/send_data", processData)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	for true {
		jsonFile, err := os.Open("dataset.json")

		fmt.Println("Successfully Opened dataset.json")

		byteValue, _ := ioutil.ReadAll(jsonFile)

		if err != nil {
			fmt.Println(err)
		}

		err = socket.WriteMessage(websocket.TextMessage, byteValue)
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second)
		defer jsonFile.Close()
	}
}

func processData(w http.ResponseWriter, r *http.Request) {

}
