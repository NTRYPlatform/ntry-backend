package ws

import (
	"fmt"
	"net/http"

	"github.com/NTRYPlatform/ntry-backend/eth"
	"github.com/gorilla/mux"

	"github.com/gorilla/websocket"
)

//TODO: create abstraction
var regSubscribers = make(map[string]*websocket.Conn)
var contractSubscribers = make(map[string]*websocket.Conn)

func WriteToRegisterChannel(register <-chan string, err chan<- struct{}) {

	// Grab the next message from the register channel
	for {
		select {
		case m := <-register:
			// Send it out to the user it needs to go to
			if client, ok := regSubscribers[m]; ok {
				err := client.WriteJSON("{\"registered\":true, \"uid\":\"" + m + "\"}")
				if err != nil {
					fmt.Printf("Error writing to connection for user: %s\n", m)
				}
				client.Close()
				delete(regSubscribers, m)
			}
		default:

		}
	}
	err <- struct{}{}
}

func WriteToContractChannel(contract <-chan interface{}, err chan<- struct{}) {

	// Grab the next message from the contract channel
	for {
		select {
		case m := <-contract:
			c, ok := m.(eth.ContractNotification)
			// Send it out to the user it needs to go to
			if ok {
				if client, ok := contractSubscribers[c.NotifyParty]; ok {
					err := client.WriteJSON(m)
					if err != nil {
						fmt.Printf("Error writing to connection for user: %s\n", m)
					}
				}
			}
		default:
		}
	}
	err <- struct{}{}
}

func ServeRegWs(w http.ResponseWriter, r *http.Request) {
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

	regSubscribers[v["uid"]] = ws
}

func ServeContractWs(w http.ResponseWriter, r *http.Request) {
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

	contractSubscribers[v["uid"]] = ws
}
