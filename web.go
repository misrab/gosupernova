package main


import (
	// Stanford Library
	"encoding/json"
	"fmt"
	"net/http"
	//"html/template"
	"io/ioutil"
	//"io"
	"os"
	"log"
	//"regexp"

	// Github external
	"github.com/gorilla/mux"

	// Custom internal
	"github.com/Misrab/gosupernova/app/session"
)


// error response contains everything we need to use http.Error
type handlerError struct {
	Error   error
	Message string
	Code    int
}

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

// attach the standard ServeHTTP method to our handler so the http library can call it
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here we could do some prep work before calling the handler if we wanted to

	// call the actual handler
	response, err := fn(w, r)

	// check for errors
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Message), err.Code)
		return
	}
	if response == nil {
		log.Printf("ERROR: response from method is nil\n")
		http.Error(w, "Internal server error. Check the logs.", http.StatusInternalServerError)
		return
	}

	// turn the response into JSON
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// send the response and log
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}


/*
 *	Wild card routing
 */
/*
type route struct {
    pattern *regexp.Regexp
    handler http.Handler
}

type RegexpHandler struct {
    routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
    h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
    h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    for _, route := range h.routes {
        if route.pattern.MatchString(r.URL.Path) {
            route.handler.ServeHTTP(w, r)
            return
        }
    }
    // no pattern matched; send 404 response
    http.NotFound(w, r)
}
*/


/*
 *	Main
 */

func main() {
	// Handle non-api requests by serving Angular
	dir := "angular/"
	fs := http.Dir(dir)
	fileHandler := http.FileServer(fs)

	// Routes
	router := mux.NewRouter()
	
	// Angular
	router.HandleFunc("/", serveAngular)
	router.HandleFunc("/create", serveAngular)
	router.HandleFunc("/workspace", serveAngular)
	
	// API
	router.HandleFunc("/api/v1/session", session.SessionHandler)

	// static files
	router.PathPrefix("/angular/").Handler(http.StripPrefix("/angular", fileHandler))
	http.Handle("/", router)


	// Listening
	port := os.Getenv("PORT")
	fmt.Println("Listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}



func serveAngular(w http.ResponseWriter, req *http.Request) {
	// read in index.html
	body, err := ioutil.ReadFile("./angular/index.html")
    if err != nil {
        panic(err)
    }
	fmt.Fprintf(w, string(body))
}