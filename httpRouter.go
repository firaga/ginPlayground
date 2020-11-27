package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func Index2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Blog:%s \nWechat:%s", "www.flysnow.org", "flysnow_org")
}
func main() {
	router := httprouter.New()
	router.GET("/", Index2)

	log.Fatal(http.ListenAndServe(":8080", router))
}
