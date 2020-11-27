package main

import (
	"fmt"
	"log"
	"net/http"
)

type CustomMux struct {
}

func (cm *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Blog:www.flysnow.org\nwechat:flysnow_org")
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &CustomMux{}))
}
