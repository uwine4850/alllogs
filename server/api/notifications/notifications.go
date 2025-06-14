package notifications

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type NotificationInfoMessage struct {
	rest.ImplementDTOMessage
	TypNotificationInfoMessage rest.TypeId `dto:"-typeid"`
	PID                        string      `dto:"PID"`
	Text                       string      `dto:"Text"`
}

const (
	TYPE_ERROR = iota
	TYPE_INFO
	TYPE_GROUP_INVITE
	TYPE_PROJECT
)

type WSMessage struct {
	Type    int
	UID     int
	Payload interface{}
}

var connections = map[int][]*websocket.Conn{}
var connectionToUser = map[*websocket.Conn]int{}

func Notification(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	socket := router.NewWebsocket(router.Upgrader)
	socket.OnConnect(func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		authJWT := r.URL.Query().Get("authJWT")
		if authJWT == "" {
			if err := conn.WriteJSON(&WSMessage{Type: TYPE_ERROR, Payload: "No authJWT."}); err != nil {
				fmt.Println("Send message error:", err)
			}
			return
		}
		_claims := &auth.JWTClaims{}
		_, err := jwt.ParseWithClaims(authJWT, _claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(manager.Key().HashKey()), nil
		})
		if err != nil {
			if err := conn.WriteJSON(&WSMessage{Type: TYPE_ERROR, Payload: err.Error()}); err != nil {
				fmt.Println("Send message error:", err)
			}
			return
		}
		connections[_claims.Id] = append(connections[_claims.Id], conn)
		connectionToUser[conn] = _claims.Id
	})
	socket.OnClientClose(func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		UID, ok := connectionToUser[conn]
		if ok {
			index := slices.Index(connections[UID], conn)
			if index != -1 {
				delete(connectionToUser, conn)
				newConnections := slices.Delete(connections[UID], index, index+1)
				if len(newConnections) == 0 {
					delete(connections, UID)
				} else {
					connections[UID] = newConnections
				}
			}
		}
	})
	socket.OnMessage(func(messageType int, msgData []byte, conn *websocket.Conn) {
		var wsMessage WSMessage
		fmt.Println(string(msgData))
		if err := json.Unmarshal(msgData, &wsMessage); err != nil {
			if err := conn.WriteJSON(&WSMessage{Type: TYPE_ERROR, Payload: err.Error()}); err != nil {
				fmt.Println("Send message error:", err)
			}
			return
		}
		userConnections := connections[wsMessage.UID]
		for i := 0; i < len(userConnections); i++ {
			userConn := userConnections[i]
			if err := userConn.WriteMessage(websocket.TextMessage, msgData); err != nil &&
				!websocket.IsUnexpectedCloseError(err, websocket.CloseNoStatusReceived) {
				fmt.Println("Send message error:", err)
			}
		}
	})
	if err := socket.ReceiveMessages(w, r); err != nil {
		fmt.Println("Receive messages error:", err)
	}
	return nil
}
