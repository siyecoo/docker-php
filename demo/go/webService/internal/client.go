package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	// send buffer size
	bufSize = 256
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	heartBeating chan string
}

type MsgInfo struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (c *Client) heartBeatingMethod() {

	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case <-c.heartBeating:
			fmt.Println(fmt.Sprintf("当前时间[%s],客户端[%s]:心跳检测成功.", time.Now().String(), c.conn.RemoteAddr().String()))
		case <-time.After(writeWait):
			fmt.Println(fmt.Sprintf("当前时间[%s],客户端[%s]:超过心跳时间,退出结束.", time.Now().String(), c.conn.RemoteAddr().String()))
			c.conn.Close()
			return
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		fmt.Println("readPump")
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(fmt.Sprintf("当前时间[%s],客户端[%s]:读取数据失败,ERROR[%s]", time.Now().String(), c.conn.RemoteAddr().String(), err))
			}
			break
		}

		var msgInfo MsgInfo

		if err := json.Unmarshal(message, &msgInfo); err != nil {
			fmt.Println(fmt.Sprintf("当前时间[%s],客户端[%s]:数据格式解析异常[%s]", time.Now().String(), c.conn.RemoteAddr().String(), err))
		}

		switch msgInfo.Type {
		case "ping":
			fmt.Println(fmt.Sprintf("当前时间[%s],客户端[%s]:客户端发送心跳包[%s]", time.Now().String(), c.conn.RemoteAddr().String(), string(message)))
			c.heartBeating <- msgInfo.Message
		case "broadcast":
			if msgInfo.Message == "S5jqlCFsHAErBTKO" {
				c.hub.broadcast <- msgInfo.Message
			}
		}

	}
	fmt.Println("-----------------------------------结束readPump")
}

func (c *Client) writePump() {

	defer func() {
		c.conn.Close()
	}()

	for {
		fmt.Println("writePump")
		select {
		case message, ok := <-c.send:
			fmt.Println(fmt.Sprintf("当前时间[%s],客户端[%s]:后台推送广播数据到该客户端,推送内容[%s]", time.Now().String(), c.conn.RemoteAddr().String(), string(message)))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
	fmt.Println("----------------------------------------结束writePump")

}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:          hub,
		conn:         conn,
		send:         make(chan []byte, bufSize),
		heartBeating: make(chan string, bufSize),
	}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.

	go client.readPump()
	go client.writePump()
	client.heartBeatingMethod()
	fmt.Println("取消了")

}
