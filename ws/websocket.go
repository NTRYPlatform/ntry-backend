package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/websocket"
)

var subscribers = make(map[string]*websocket.Conn)

func WriteToRegisterChannel(register <-chan string, err chan<- struct{}) {

	// Grab the next message from the register channel
	for {
		select {
		case m := <-register:
			// Send it out to the user it needs to go to
			if client, ok := subscribers[m]; ok {
				err := client.WriteJSON("{\"registered\":true, \"uid\":\"" + m + "\"}")
				if err != nil {
					fmt.Printf("Error writing to connection for user: %s\n", m)
				}
				client.Close()
				delete(subscribers, m)
			}
		default:

		}
	}
	err <- struct{}{}
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// log.Println(err)
		return
	}
	//TODO: figure out a way to unregister/close

	subscribers[v["uid"]] = ws
}
