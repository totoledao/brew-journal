package utils

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func ping(ws *websocket.Conn) {
	defer ws.Close()

	for {
		// Receive a ping message
		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Sleep for 500 milliseconds
		time.Sleep(500 * time.Millisecond)

		// Send a ping message
		err = websocket.Message.Send(ws, "ping")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func PingHandler(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			ping(ws)
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	}
