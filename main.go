package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"os"
)
func main()  {

	r := gin.Default()
	m := melody.New()
	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "chat.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(os.Getenv("PORT"))
}
