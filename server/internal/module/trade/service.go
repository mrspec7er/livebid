package trade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mrspec7er/livebid/server/internal/database"
)

type Client struct {
	conn *websocket.Conn
}

type Service struct {
	Store   database.DBConn
	Clients []*Client
}

func (s *Service) WebsocketConnection() websocket.Upgrader {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return upgrader
}

func (s *Service) BroadcastMessage(msgType int, message []byte) {
	for _, client := range s.Clients {
		err := client.conn.WriteMessage(msgType, message)
		if err != nil {
			fmt.Println("error while broadcasting: ", err)
			client.conn.Close()
		}
	}
}

func (s *Service) ProcessMessage(conn *websocket.Conn, itemId string) (int, error) {
	defer conn.Close()

	client := &Client{
		conn: conn,
	}
	s.Clients = append(s.Clients, client)

	defer func() {
		client.conn.Close()

		for i, c := range s.Clients {
			if c == client {
				s.Clients = append(s.Clients[:i], s.Clients[i+1:]...)
				break
			}
		}
	}()
	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			return 500, err
		}

		newMsg := &database.TradeMessage{
			ID:         time.Now().GoString(),
			CreatorID:  "user01",
			ItemNumber: itemId,
			Message:    string(data),
		}

		msgPayload, err := json.Marshal(newMsg)
		if err != nil {
			return 500, err
		}

		s.BroadcastMessage(msgType, msgPayload)
	}
}
