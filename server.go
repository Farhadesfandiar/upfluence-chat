package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// type Message struct {
// 	Greeting string `json:"greeting"`
// }

// var (
// 	wsUpgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}
// 	wsConn *websocket.Conn
// )

// func WsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
// 		// check http request
// 		return true
// 	}

// 	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Printf("could not upgrade: %s\n", err.Error())
// 		return
// 	}
// 	defer wsConn.Close()

// 	for {
// 		var msg Message
// 		err := wsConn.ReadJSON(&msg)
// 		if err != nil {
// 			fmt.Printf("error reading JSON %s\n", err.Error())
// 			break
// 		}

// 		fmt.Printf("Message Received: %s\n", msg.Greeting)
// 	}
// }

// func main() {
// 	rotuer := mux.NewRouter()
// 	log.Fatal(http.ListenAndServe("localhost:9100", rotuer))
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		err2 := conn.WriteMessage(messageType, []byte("Hello Client"))
		if err2 != nil {
			fmt.Printf("error reading message: ", err.Error())

		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		log.Printf("listening ...")

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)

}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat Assignment v0.01")
	setupRoutes()
	http.ListenAndServe(":9100", nil)
}
