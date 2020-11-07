package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws/wsutil"

	"github.com/gobwas/ws"
)

var n net.Conn

func web(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Println("连接失败：", err)
	}
	for {
		defer conn.Close()
		var (
			state  = ws.StateServerSide
			reader = wsutil.NewReader(conn, state)            //创建一个读取器，从conn中读取状态以保持连接
			writer = wsutil.NewWriter(conn, state, ws.OpText) //创建一个新的缓存数据区,ws.OpText是操作码
		)
		header, _ := reader.NextFrame()                    //读取conn中的下一条数据
		writer.Reset(conn, state, header.OpCode)           //重置缓冲数据区，并给其制定新的操作码
		if _, err := io.Copy(writer, reader); err != nil { //将读取到的数据copy到writer
			fmt.Println("copy err :", err)
		}
		if err := writer.Flush(); err != nil { //Flush将所有缓冲数据写入底层的io.Writer(发送给客户端)
			fmt.Println("Flush err :", err)
		}

	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/web.html")
	})
	http.HandleFunc("/ws", web)
	http.ListenAndServe(":1234", nil)
}
