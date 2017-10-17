package ws

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/NTRYPlatform/ntry-backend/eth"
	"github.com/gorilla/mux"

	"github.com/gorilla/websocket"
)

//TODO: create abstraction
<<<<<<< HEAD
=======
var regSubscribers = make(map[string]*websocket.Conn)
var contractSubscribers = make(map[string]*websocket.Conn)
var approvedContractSubscribers = make(map[string]*websocket.Conn)
>>>>>>> yg-dev

type ChannelSub struct {
	sync.Mutex
	subscribers map[string]*websocket.Conn
}

type Channels struct {
	regChannel ChannelSub
	conChannel ChannelSub
}

func NewChannels() *Channels {
	ch := Channels{regChannel: ChannelSub{subscribers: make(map[string]*websocket.Conn)}, conChannel: ChannelSub{subscribers: make(map[string]*websocket.Conn)}}
	return &ch
}

// var regSubscribers = make(map[string]*websocket.Conn)
// var contractSubscribers = make(map[string]*websocket.Conn)

func (ch *Channels) WriteToRegisterChannel(register <-chan string, err chan<- struct{}) {

	// Grab the next message from the register channel
	for {
		select {
		case m := <-register:
			// Send it out to the user it needs to go to
			if client, ok := ch.regChannel.subscribers[m]; ok {
				err := client.WriteJSON("{\"registered\":true, \"uid\":\"" + m + "\"}")
				if err != nil {
					fmt.Printf("Error writing to connection for user: %s\n", m)
				}
				client.Close()
				delete(ch.regChannel.subscribers, m)
			}
		default:

		}
	}
	err <- struct{}{}
}

func (ch *Channels) WriteToContractChannel(contract <-chan interface{}, err chan<- struct{}) {

	// Grab the next message from the contract channel
	for {
		select {
		case m := <-contract:
			fmt.Printf("%v\n", m)
			c, ok := m.(eth.ContractNotification)
			// Send it out to the user it needs to go to
			if ok {
<<<<<<< HEAD
				if client, ok := ch.conChannel.subscribers[c.NotifyParty]; ok {
					err := client.WriteJSON(c.Contract)
					if err != nil {
						fmt.Printf("Error writing to connection for user: %s\n", m)
=======
				if c.Type == "new" {
					if client, ok := contractSubscribers[c.NotifyParty]; ok {
						err := client.WriteJSON(c)
						if err != nil {
							fmt.Printf("Error writing to connection for user: %s\n", m)
						}
					}
				} else {
					if client, ok := approvedContractSubscribers[c.NotifyParty]; ok {
						err := client.WriteJSON(c)
						if err != nil {
							fmt.Printf("Error writing to connection for user: %s\n", m)
						}
>>>>>>> yg-dev
					}
				}
			}
		default:
		}
	}
	err <- struct{}{}
}

func (ch *Channels) ServeRegWs(w http.ResponseWriter, r *http.Request) {
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
<<<<<<< HEAD
	ch.regChannel.Lock()
	ch.regChannel.subscribers[v["uid"]] = ws
	ch.regChannel.Unlock()
=======

	regSubscribers[v["uid"]] = ws
>>>>>>> yg-dev
}

func (ch *Channels) ServeContractWs(w http.ResponseWriter, r *http.Request) {
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
	ch.conChannel.Lock()
	ch.conChannel.subscribers[v["uid"]] = ws
	ch.conChannel.Unlock()
}

func ServeApprovedContractWs(w http.ResponseWriter, r *http.Request) {
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

	approvedContractSubscribers[v["uid"]] = ws
}
