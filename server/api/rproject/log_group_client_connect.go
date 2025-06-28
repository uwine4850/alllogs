package rproject

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"slices"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/rest"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

const (
	TYPE_ERROR   = 0
	TYPE_LOGITEM = 1
)

type MsgLogItem struct {
	rest.ImplementDTOMessage
	TypLogItemMessage rest.TypeId       `dto:"-typeid"`
	Type              int               `dto:"Type"`
	Token             string            `dto:"Token"`
	Error             string            `dto:"Error"`
	Payload           MsgLogItemPayload `dto:"Payload"`
}

type MsgLogItemPayload struct {
	rest.ImplementDTOMessage
	TypLogItemPayload rest.TypeId `dto:"-typeid" json:"-"`
	Id                int         `dto:"Id" db:"id"`
	LogGroupId        int         `dto:"LogGroupId" db:"log_group_id"`
	Text              string      `dto:"Text" db:"text"`
	Type              string      `dto:"Type" db:"type"`
	Tag               string      `dto:"Tag" db:"tag"`
	Datetime          string      `dto:"Datetime" db:"datetime"`
}

var connections = map[string][]*websocket.Conn{}
var connectionsToClient = map[*websocket.Conn]string{}

func LogClientSocket(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	socket := router.NewWebsocket(router.Upgrader)
	socket.OnConnect(func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		token := r.URL.Query().Get("token")
		if token == "" {
			if err := conn.WriteJSON(&MsgLogItem{Type: TYPE_ERROR, Error: "token not found"}); err != nil {
				fmt.Println("Client log send message error:", err)
			}
		}
		connections[token] = append(connections[token], conn)
		connectionsToClient[conn] = token
	})
	socket.OnClientClose(func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		removeConnection(conn)
	})
	socket.OnMessage(func(messageType int, msgData []byte, conn *websocket.Conn) {
		var logItemMessage MsgLogItem
		if err := json.Unmarshal(msgData, &logItemMessage); err != nil {
			if err := conn.WriteJSON(&MsgLogItem{Type: TYPE_ERROR, Error: err.Error()}); err != nil {
				fmt.Println("Client log send message error:", err)
			}
			return
		}

		insertedId, err := writeLogItem(&logItemMessage)
		if err != nil {
			if err := conn.WriteJSON(&MsgLogItem{Type: TYPE_ERROR, Error: err.Error()}); err != nil {
				fmt.Println("Client log send message error:", err)
			}
			return
		}
		logItemMessage.Payload.Id = int(insertedId)

		logItemMessageBytes, err := json.Marshal(logItemMessage)
		if err != nil {
			if err := conn.WriteJSON(&MsgLogItem{Type: TYPE_ERROR, Error: err.Error()}); err != nil {
				fmt.Println("Client log send message error:", err)
			}
			return
		}

		clientConnections := connections[logItemMessage.Token]
		closedConnections := []*websocket.Conn{}
		for i := 0; i < len(clientConnections); i++ {
			clientConnection := clientConnections[i]
			if err := clientConnection.WriteMessage(websocket.TextMessage, logItemMessageBytes); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseNoStatusReceived) || errors.Is(err, net.ErrClosed) {
					closedConnections = append(closedConnections, clientConnection)
				} else {
					fmt.Println("Client log send message error:", err)
				}
			}
		}
		for i := 0; i < len(closedConnections); i++ {
			removeConnection(closedConnections[i])
		}
	})
	if err := socket.ReceiveMessages(w, r); err != nil {
		fmt.Println("Client log messages error:", err)
	}
	return nil
}

func removeConnection(conn *websocket.Conn) {
	token, ok := connectionsToClient[conn]
	if ok {
		index := slices.Index(connections[token], conn)
		if index != -1 {
			delete(connectionsToClient, conn)
			newConnections := slices.Delete(connections[token], index, index+1)
			if len(newConnections) == 0 {
				delete(connections, token)
			} else {
				connections[token] = newConnections
			}
		}
	}
}

func writeLogItem(logItemMessage *MsgLogItem) (int64, error) {
	params, err := mapper.ParamsValueFromDbStruct(typeopr.Ptr{}.New(&logItemMessage.Payload), []string{})
	if err != nil {
		return -1, err
	}
	paramsKeys, paramsValues := dbutils.ParseParams(params)

	authzCheck := qb.Exists(
		qb.SQ(false,
			qb.NewNoDbQB().SelectFrom(1, cnf.DBT_PROFILE).
				InnerJoin(cnf.DBT_PROJECT_LOG_GROUP, qb.Compare(
					cnf.DBT_PROJECT_LOG_GROUP+".id", qb.EQUAL, logItemMessage.Payload.LogGroupId,
				)).
				InnerJoin(cnf.DBT_PROJECT, qb.NoArgsCompare(
					cnf.DBT_PROJECT+".id", qb.EQUAL, cnf.DBT_PROJECT_LOG_GROUP+".project_id",
				)).
				Where(qb.Compare(
					cnf.DBT_PROFILE+".token", qb.EQUAL, logItemMessage.Token,
				), qb.AND, qb.NoArgsCompare(
					cnf.DBT_PROJECT+".user_id", qb.EQUAL, cnf.DBT_PROFILE+".user_id",
				)),
		),
	)

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	insertString := fmt.Sprintf("INSERT INTO %s (%s) SELECT %s",
		cnf.DBT_LOG_ITEM, strings.Join(paramsKeys, ","), dbutils.RepeatValues(len(paramsKeys), ","))
	newQB.Custom(insertString, paramsValues...).Where(authzCheck).Merge()
	res, err := newQB.Exec()
	if err != nil {
		return -1, err
	}
	if res["rowsAffected"].(int64) != 1 {
		return -1, errors.New("error adding log: access denied")
	}
	return res["insertID"].(int64), nil
}
