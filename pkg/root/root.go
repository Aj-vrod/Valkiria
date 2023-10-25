package root

import (
	"fmt"
	"io"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "This is my website!\n")
}
