package handler

import (
	"fmt"
	"net/http"
)

func GetHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
