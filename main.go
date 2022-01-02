package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	//User string `json:"user"`
	Info string `json:"info"`
}

var People map[*websocket.Conn]int

var BroadCast chan []byte

//var upgrader websocket.Upgrader

func main() {
	//fmt.Println("Go语言标准包里面没有提供对WebSocket的支持，但是在由官方维护的go.net子包中有对这个的支持 go get golang.org/x/net/websocket")
	//打印这个信息就，os.Exit(1)  退出程序
	//log.Fatal("shiming")  todo  草拟吗 啊   看清楚啊   后面的域名的地址 有个老子的名字啊
	//http.Handle("/shiming", websocket.Handler(Echo))
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}
	People = make(map[*websocket.Conn]int, 10)
	BroadCast = make(chan []byte)
	router := gin.Default()
	router.GET("/chatws", Ping)
	router.Run(":8989")

	//listener, err := net.Listen("tcp", "localhost:8989")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//go broadcaster()
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Println(err)
	//		continue
	//	}
	//	go handleConn(conn)
	//}
	//http.HandleFunc("/shiming", serveWs)
	//http.ListenAndServe("localhost:8989", nil)
}

func Ping(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	go broadcaster()

	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// 客户端关闭连接时也会进入
			fmt.Println(err)
			break
		}
		if message == nil {
			continue
		}
		// msg := &data{}
		// json.Unmarshal(message, msg)
		// fmt.Println(msg)
		fmt.Println(mt)
		fmt.Println(message)
		fmt.Println(string(message))
		fmt.Println("远程主机:", ws.RemoteAddr())
		var msg Message
		msg.Info = string(message)

		People[ws] = mt
		BroadCast <- message

		//fmt.Println("jjj")

		// 如果客户端发送ping就返回pong,否则数据原封不动返还给客户端
		if string(message) == "ping" {
			message = []byte("pong")
		}
		// 写入ws数据 二进制返回
		//err = ws.WriteMessage(mt, message)
		// 返回JSON字符串，借助gin的gin.H实现
		// v := gin.H{"message": msg}
		// err = ws.WriteJSON(v)
		//if err != nil {
		//	break
		//}
	}
}

//type client chan<- string
//
//var (
//	entering = make(chan client)
//	leaving  = make(chan client)
//	message  = make(chan string)
//)

func broadcaster() {
	//clients := make(map[client]bool)
	for {
		select {
		case msg := <-BroadCast:
			for k, v := range People {
				//这里的cli就是handleConn里的ch channel，
				//writeToCLient goroutine一直在监听ch channel，读取channel中的内容，并写入客户端连接
				//cli <- msg

				k.WriteMessage(v, msg)
			}
			//case cli := <-entering:
			//	clients[cli] = true
			//case cli := <-leaving:
			//	delete(clients, cli)
			//	close(cli)
		}
	}
}

//func handleConn(conn net.Conn) {
//	ch := make(chan string)
//	//写入消息到客户端的连接
//	go writeToCLient(conn, ch)
//
//	who := conn.RemoteAddr().String()
//	//当客户端连接过来时，给客户端一条消息
//	//注意，这时的ch会立马被writeToCLient goroutine读取，并发送到当前客户端
//	//所以已连接的其他客户端不会接受到该条消息
//	ch <- "You are " + who
//	//这里的message channel会被broadcaster读取，广播给所有已连接的客户端
//	//注意，这时当前客户端还没给entering，所以当前客户端不会接受到该条消息
//	message <- who + " are arrived"
//	//将当前客户端发送给entering channel，broadcaster会将当前客户端添加到已连接的客户端集合中
//	entering <- ch
//
//	input := bufio.NewScanner(conn)
//	//阻塞监听客户端输入
//	for input.Scan() {
//		//获取客户端输入，并发送到message channel，然后broadcaster会将它广播给所有连接的客户端
//		//因为这时，当前客户端已经添加到clients集合中，所以当前客户端也会接受到消息
//		message <- who + ": " + input.Text()
//	}
//
//	//客户端断开连接
//	leaving <- ch
//	message <- who + " are left"
//	conn.Close()
//}
//
//func writeToCLient(conn net.Conn, ch <-chan string) {
//	for msg := range ch {
//		fmt.Fprintln(conn, msg)
//	}
//}

//func serveWs(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Println("upgrade:", err)
//		return
//	}
//
//	for {
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		if err := conn.WriteMessage(messageType, []byte("收到:"+string(p))); err != nil {
//			log.Println(err)
//			return
//		}
//		fmt.Println(string(p))
//	}
//}

//func Echo(w *websocket.Conn) {
//var error error
//for {
//	var reply string
//	if error = websocket.Message.Receive(w, &reply); error != nil {
//		fmt.Println("不能够接受消息 error==", error)
//		break
//	}
//	fmt.Println("能够接受到消息了--- ", reply)
//	msg := "我已经收到消息 Received:" + reply
//	//  连接的话 只能是   string；类型的啊
//	fmt.Println("发给客户端的消息： " + msg)
//
//	if error = websocket.Message.Send(w, msg); error != nil {
//		fmt.Println("不能够发送消息 悲催哦")
//		break
//	}
//}
