package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ChatWith(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "welcome to chatroom"
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("chat.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func Register() {
	log.Println("start...")
	http.Handle("/", websocket.Handler(ChatWith))
	http.HandleFunc("/chat", ClientHandler)
}

func main() {
	Register()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
