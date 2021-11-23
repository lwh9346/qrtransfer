package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

//go:embed dist
var dist embed.FS

type fsWithPrefix struct {
	Fs embed.FS
}

func (f *fsWithPrefix) Open(p string) (fs.File, error) {
	return f.Fs.Open(path.Join("dist", p))
}

var room *Room
var roomLock sync.Mutex

func main() {
	room = &Room{Alive: true, QR: "N"}
	go func() {
		ticker := time.Tick(600 * time.Second)
		for {
			<-ticker
			roomLock.Lock()
			roomOld := room
			room = &Room{QR: "N", Alive: true}
			roomOld.CloseRoom()
			roomLock.Unlock()
		}
	}()
	r := gin.Default()
	r.POST("/update", UpdateQRCodeHanlder)
	r.GET("/get", GetQRCodeHandler)
	r.StaticFS("/ui", http.FS(&fsWithPrefix{dist}))
	r.GET("/", RedirectMainPage)
	r.Run("0.0.0.0:8888")
}

func RedirectMainPage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "ui")
}

func UpdateQRCodeHanlder(c *gin.Context) {
	type req struct {
		QR string `json:"qr"`
	}
	var r req
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(400, gin.H{})
	}
	roomLock.Lock()
	room.UpdateQRCode(r.QR)
	roomLock.Unlock()
	c.JSON(200, gin.H{})
}

func GetQRCodeHandler(c *gin.Context) {
	if !c.IsWebsocket() {
		c.JSON(400, gin.H{})
		return
	}
	websocket.Handler(func(conn *websocket.Conn) {
		defer conn.Close()
		signal := make(chan (int))
		report := make(chan (error))

		roomLock.Lock()
		_, err := conn.Write([]byte(room.QR))
		if err != nil {
			roomLock.Unlock()
			return
		}
		room.AddClient(&Client{Signal: signal, Report: report, Alive: true})
		log.Println("adding client")
		roomLock.Unlock()
		for {
			s := <-signal
			if s == 1 {
				log.Println("closing client")
				return
			}
			_, err := conn.Write([]byte(room.QR))
			report <- err
			if err != nil {
				return
			}
		}
	}).ServeHTTP(c.Writer, c.Request)
}
