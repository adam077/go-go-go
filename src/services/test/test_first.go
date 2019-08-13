package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket 请求ping 返回pong
func Haha2(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func(ws *websocket.Conn) {
		fmt.Println(1)
		ws.Close()
	}(ws)
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		//if err != nil {
		//	fmt.Println(err)
		//	break
		//}
		if string(message) == "ping" {
			message = []byte("pong")
		} else {
			message = []byte("pong2")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
