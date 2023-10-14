package hello

import (
	"fmt"
	"io"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello, user!\n")
}
