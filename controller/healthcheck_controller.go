package controller

import (
	"fmt"
	"net/http"
)

func GetHealthcheckController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
