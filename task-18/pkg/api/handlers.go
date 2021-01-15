package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}

// handleMessage аутентифицирует пользователя, принимает сообщение и рассылает его другим участникам чата
func (a *API) handleMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	mt, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		_ = conn.WriteMessage(mt, []byte(err.Error()))
		return
	}

	if string(message) != "password" {
		err = fmt.Errorf("invalid password")
		fmt.Println(err)
		_ = conn.WriteMessage(mt, []byte(err.Error()))
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("OK"))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, message, err = conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		_ = conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	a.chat.Broadcast(string(message))
}

// handleMessages позволяет участникам чата подписаться на новые сообщения
func (a *API) handleMessages(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ch := make(chan string)
	memberId := a.chat.Subscribe(ch)
	defer a.chat.Unsubscribe(memberId)

	for {
		msg := <-ch
		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
