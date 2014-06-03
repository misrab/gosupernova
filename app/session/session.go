package session


import (
	"fmt"
	"net/http"
)

func SessionHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	// GET
	case "GET":
		fmt.Fprintf(w, "No such route")
	// DELETE
	case "DELETE":
		fmt.Fprintf(w, "deleting session")
	default:
		fmt.Fprintf(w, "No such route")
	}
}