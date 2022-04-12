package main

import(
	"net/http"
)

func chat(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"chat.html")
}
func main()  {

	systemFolder := http.FileServer(http.Dir("./"))
	http.Handle("/", systemFolder)
	http.HandleFunc("/chat", chat)
	http.ListenAndServe(":8000", nil)
}
