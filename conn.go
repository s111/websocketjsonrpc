// Package websocketjsonrpc takes a *github.com/gorilla/websocket.Conn and wraps it with the io.ReaderWriterCloser interface.
package websocketjsonrpc

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type conn struct {
	ws *websocket.Conn
}

// NewConn returns a new connection implementing the io.ReadWriteCloser interface
func NewConn(ws *websocket.Conn) *conn {
	return &conn{ws}
}

func (c *conn) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	var raw json.RawMessage
	err := c.ws.ReadJSON(&raw)

	copy(p, raw)

	return len(raw), err
}

func (c *conn) Write(p []byte) (int, error) {
	return len(p), c.ws.WriteMessage(websocket.TextMessage, p)
}

func (c *conn) Close() error {
	return c.ws.Close()
}
