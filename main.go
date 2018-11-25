package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	data := []Documents{
		Documents{
			Name:  "BTOW3",
			Value: 37,
		},
		Documents{
			Name:  "CCRO3",
			Value: 11,
		},
		Documents{
			Name:  "CIEL3",
			Value: 9,
		},
		Documents{
			Name:  "CMIG4",
			Value: 12,
		},
		Documents{
			Name:  "CPLE6",
			Value: 28,
		},
	}

	m := Message{
		Payload: data,
	}

	b, err := json.Marshal(m)

	err = socket.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println(err)
	}
}

func processData(w http.ResponseWriter, r *http.Request) {

}
