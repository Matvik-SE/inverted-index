package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var threadsNum = 1
var dataDir = "./data/"
var dataExt = ".txt"
var filesArray []string
var upgrader = websocket.Upgrader{}

func main() {
	filesArray = getAllFiles(dataDir, dataExt)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/threads", threadsHandler)
	http.HandleFunc("/socket", socketHandler)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inverted Index Sockets App")
}

func threadsHandler(w http.ResponseWriter, r *http.Request) {
	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)

	if err == nil && number > 0 {
		threadsNum = number
		fmt.Fprintf(w, "Number of threads has been changed: %d", number)
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Error during connection:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error during socket message reading:", err)
			break
		}

		log.Printf("Received string: '%s'", message)

		result := runIndexThreads(filesArray, threadsNum, string(message))

		if len(result) > 0 {
			for key, value := range result {
				err = conn.WriteMessage(messageType, []byte(key+" - "+value.(string)+" times"))

				if err != nil {
					break
				}
			}
		}
	}
}
