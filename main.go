package main
import "github.com/akohring/slaythewild_basecamp/uart"

// import (
//     "bufio"
//     "encoding/json"
//     "flag"
//     "fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"
// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/websocket"
// )

// // Configuration is used to load external overrides
// type Configuration struct {
//     address string
//     imageStore string
// }

// var (
// 	addr    = flag.String("addr", "127.0.0.1:8080", "http service address")
//     cmdPath string
//     configuration = Configuration{}
// )

// const (
// 	// Time allowed to write a message to the peer.
// 	writeWait = 10 * time.Second

// 	// Maximum message size allowed from peer.
// 	maxMessageSize = 8192

// 	// Time allowed to read the next pong message from the peer.
// 	pongWait = 60 * time.Second

// 	// Send pings to peer with this period. Must be less than pongWait.
// 	pingPeriod = (pongWait * 9) / 10

// 	// Time to wait before force close on connection.
//     closeGracePeriod = 10 * time.Second
// )

// func initConfig(fileName string) {
//     file, _ := os.Open(fileName)
//     decoder := json.NewDecoder(file)
//     config := Configuration{}
//     err := decoder.Decode(&config)
//     if err != nil {
//         log.Println("error:", err)
//     }
//     log.Println(config.address)
// }

// func pumpStdin(ws *websocket.Conn, w io.Writer) {
// 	defer ws.Close()
// 	ws.SetReadLimit(maxMessageSize)
// 	ws.SetReadDeadline(time.Now().Add(pongWait))
// 	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
// 	for {
// 		_, message, err := ws.ReadMessage()
// 		if err != nil {
// 			break
// 		}
// 		message = append(message, '\n')
// 		if _, err := w.Write(message); err != nil {
// 			break
// 		}
// 	}
// }

// func pumpStdout(ws *websocket.Conn, r io.Reader, done chan struct{}) {
// 	defer func() {
// 	}()
// 	s := bufio.NewScanner(r)
// 	for s.Scan() {
// 		ws.SetWriteDeadline(time.Now().Add(writeWait))
// 		if err := ws.WriteMessage(websocket.TextMessage, s.Bytes()); err != nil {
// 			ws.Close()
// 			break
// 		}
// 	}
// 	if s.Err() != nil {
// 		log.Println("scan:", s.Err())
// 	}
// 	close(done)

// 	ws.SetWriteDeadline(time.Now().Add(writeWait))
// 	ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
// 	time.Sleep(closeGracePeriod)
// 	ws.Close()
// }

// func ping(ws *websocket.Conn, done chan struct{}) {
// 	ticker := time.NewTicker(pingPeriod)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
// 				log.Println("ping:", err)
// 			}
// 		case <-done:
// 			return
// 		}
// 	}
// }

// func internalError(ws *websocket.Conn, msg string, err error) {
// 	log.Println(msg, err)
// 	ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
// }

// var upgrader = websocket.Upgrader{}

// func webSocketHandler(w http.ResponseWriter, r *http.Request) {
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("upgrade:", err)
// 		return
// 	}

// 	defer ws.Close()

// 	outr, outw, err := os.Pipe()
// 	if err != nil {
// 		internalError(ws, "stdout:", err)
// 		return
// 	}
// 	defer outr.Close()
// 	defer outw.Close()

// 	inr, inw, err := os.Pipe()
// 	if err != nil {
// 		internalError(ws, "stdin:", err)
// 		return
// 	}
// 	defer inr.Close()
// 	defer inw.Close()

// 	proc, err := os.StartProcess(cmdPath, flag.Args(), &os.ProcAttr{
// 		Files: []*os.File{inr, outw, outw},
// 	})
// 	if err != nil {
// 		internalError(ws, "start:", err)
// 		return
// 	}

// 	inr.Close()
// 	outw.Close()

// 	stdoutDone := make(chan struct{})
// 	go pumpStdout(ws, outr, stdoutDone)
// 	go ping(ws, stdoutDone)

// 	pumpStdin(ws, inw)

// 	// Some commands will exit when stdin is closed.
// 	inw.Close()

// 	// Other commands need a bonk on the head.
// 	if err := proc.Signal(os.Interrupt); err != nil {
// 		log.Println("inter:", err)
// 	}

// 	select {
// 	case <-stdoutDone:
// 	case <-time.After(time.Second):
// 		// A bigger bonk on the head.
// 		if err := proc.Signal(os.Kill); err != nil {
// 			log.Println("term:", err)
// 		}
// 		<-stdoutDone
// 	}

// 	if _, err := proc.Wait(); err != nil {
// 		log.Println("wait:", err)
// 	}
// }

// func staticHandler(w http.ResponseWriter, r *http.Request) {
//     http.ServeFile(w, r, fmt.Sprintf("static/%s", r.URL.Path[1:]))
// }

func main() {
	uart.Run()
//     if(len(os.Args) < 2) {
//         log.Println("error:", "Provide an external configuration file")
//     } else {
// 		initConfig(os.Args[1])
// 		router := mux.NewRouter()
// 		router.HandleFunc("/people", GetPeople).Methods("GET")

//         http.FileServer(http.Dir("./static"))
//         http.HandleFunc("/", staticHandler)
//         http.HandleFunc("/ws", webSocketHandler)
//         log.Println(configuration)
//         log.Fatal(http.ListenAndServe(configuration.address, router))
//     }
}