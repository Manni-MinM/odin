package main

import "net/http"

// create a client to request a get concurrently very fast we want to add load to server.
// the path is localhost:8000/api/server/all

func main() {
	_, err := http.Get("http://localhost:8000/api/server/all")
	if err != nil {
		return
	}

	// put the code here to request concurrently
	pool := make(chan bool, 100)
	for i := 0; i < 1000; i++ {
		pool <- true
		go func() {
			_, err := http.Get("http://localhost:8000/api/server/all")
			if err != nil {
				return
			}
			<-pool
		}()
	}
}
