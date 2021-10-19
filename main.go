package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
)

var mux = http.NewServeMux()

func main() {
	// mux.HandleFunc("/panic/", panicDemo)
	// mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/calc/add/{id}", addHandler)

	log.Fatal(http.ListenAndServe(":3000", recoverMw(mux, true)))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fmt.Println(params)

	//res := Add(2, 3)

	fmt.Fprintln(w, "<h1>%v</h1>", 23)
}

func recoverMw(app http.Handler, isDev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil { //check if there was an error
				log.Println(err)

				stack := debug.Stack()
				log.Println(string(stack))

				if !isDev {
					http.Error(w, "Something went wrong :(", http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, string(stack))
				}

			}
		}()

		nw := &responseWriter{ResponseWriter: w}
		app.ServeHTTP(nw, r)
		nw.flush()
	}
}

type responseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.writes = append(rw.writes, b)
	return len(b), nil
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("The responsewriter does not support the Hijacker inteface")
	}
	return hijacker.Hijack()
}

func (rw *responseWriter) Flush() {
	flusher, ok := rw.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}
	flusher.Flush()
}

func (rw *responseWriter) flush() error {
	if rw.status != 0 {
		rw.ResponseWriter.WriteHeader(rw.status)
	}
	for _, write := range rw.writes {
		_, err := rw.ResponseWriter.Write(write)
		if err != nil {
			return err
		}
	}
	return nil
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		err := recover()
		fmt.Fprint(w, "The app is panicking now, saying: "+err.(string))
	}()*/

	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

//go run .\main.go

//http://localhost:3000/
//http://localhost:3000/panic/
//http://localhost:3000/panic-after/
