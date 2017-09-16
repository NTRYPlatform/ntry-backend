package notary

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var register = make(chan string)
var notifications = make(chan string)

// 	ws       *websocket.Conn
// 	sender   chan []byte
// 	receiver chan string
// }

// type hub struct {
// 	clients    map[*Client]bool
// 	unicast    chan string
// 	register   chan *Client
// 	unregister chan *Client

// 	content string
// }

func WriteToRegisterChannel(msg string) {

	go func(msg string) { register <- msg }(msg)
	// Grab the next message from the register channel
	select {
	case m := <-register:
		// Send it out to every client that is currently connected (TODO: shouldn't be broadcast)
		for client := range clients {
			err := client.WriteJSON(m)
			if err != nil {
				// log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method not allowed", 405)
	// 	return
	// }
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
	// defer ws.Close()
	//TODO: figure out a way to unregister/close

	clients[ws] = true
}
