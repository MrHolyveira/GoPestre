package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func chat(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"chat.html")
}
func main()  {
/*
	systemFolder := http.FileServer(http.Dir("./"))
	http.Handle("/", systemFolder)
	http.HandleFunc("/chat", chat)
	http.ListenAndServe(":8000", nil)
*/

	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "chat.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":8080")
}
