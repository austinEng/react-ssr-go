package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
	"path/filepath"
	"github.com/gorilla/mux"
	"app/ssr"
)

func main() {
	ssr.Init()

	page_requests := make(chan ssr.PageRenderRequest)

	for i := 0; i < 3; i++ {
		go ssr.NewPageRenderer(page_requests, i)
	}

	router := mux.NewRouter()

	staticPath := os.Args[1]
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the absolute path to prevent directory traversal
		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			// if we failed to get the absolute path respond with a 400 bad request
			// and stop
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		path = filepath.Join(staticPath, path)
		info, err := os.Stat(path)
		if os.IsNotExist(err) || info.IsDir() {

			// render the page
			reply := make(chan string)
			page_requests <- ssr.PageRenderRequest{Request: r, Reply: reply}
			fmt.Fprintf(w, <- reply)
			return

		} else if err != nil {
			// if we got an error (that wasn't that the file doesn't exist) stating the
			// file, return a 500 internal server error and stop
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// otherwise, use http.FileServer to serve the static dir
		http.FileServer(http.Dir(staticPath)).ServeHTTP(w, r)
	}))

	server := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Now server is running on port 3000")
	log.Fatal(server.ListenAndServe())
}
