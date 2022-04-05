package main

import(
	"net/http"
)
func main()  {

	systemFolder := http.FileServer(http.Dir("./"))
	http.Handle("/", systemFolder)
	http.ListenAndServe(":8000", nil)
}