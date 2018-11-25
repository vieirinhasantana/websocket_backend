package main

import (
	"encoding/json"
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
		defer jsonFile.Close()
		time.Sleep(3 * time.Second)
	}
}

func processData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode("OKOK")
}
