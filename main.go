package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/add/{a}/{b}", addHandler).Methods("GET")
	r.HandleFunc("/sub/{a}/{b}", subHandler).Methods("GET")
	r.HandleFunc("/mul/{a}/{b}", mulHandler).Methods("GET")
	r.HandleFunc("/div/{a}/{b}", divHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", recoverMw(r, true)))
}

type responseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
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

func addHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot add %v to %v", params["a"], params["b"])
		}
		return ""
	}()

	a, err := strconv.ParseFloat(params["a"], 64)
	if err != nil {
		w.Write([]byte("Cannot parse " + params["a"]))
		return
	}

	b, err := strconv.ParseFloat(params["b"], 64)
	if err != nil {
		w.Write([]byte("Cannot parse " + params["b"]))
		return
	}

	w.Write([]byte(strconv.FormatFloat(a+b, 'f', -1, 64)))
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot subtract %v from %v", params["b"], params["a"])
		}
		return ""
	}()

	a, err := strconv.ParseFloat(params["a"], 64)
	if err != nil {
		w.Write([]byte("Cannot parse " + params["a"]))
		return
	}

	b, err := strconv.ParseFloat(params["b"], 64)
	if err != nil {
		w.Write([]byte("Cannot parse " + params["b"]))
		return
	}

	w.Write([]byte(strconv.FormatFloat(a-b, 'f', -1, 64)))
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot multiply %v by %v", params["a"], params["b"])
		}
		return ""
	}()

	a, err := strconv.ParseFloat(params["a"], 64)
	if err != nil {
		w.Write([]byte("Cannot parse " + params["a"]))
		return
	}

	b, err := strconv.ParseFloat(params["b"], 64)
	if err != nil {
		w.Write([]byte("Cannot parse " + params["b"]))
		return
	}

	w.Write([]byte(strconv.FormatFloat(a*b, 'f', -1, 64)))
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	defer func() string {
		if err := recover(); err != nil {
			return fmt.Sprintf("Cannot divide %v by %v", params["a"], params["b"])
		}
		return ""
	}()

	a, err := strconv.ParseFloat(params["a"], 64)
	if err != nil {
		w.Write([]byte("cannot parse " + params["a"]))
		return
	}
	b, err := strconv.ParseFloat(params["b"], 64)
	if err != nil {
		w.Write([]byte("Cannot parse " + params["b"]))
		return
	}

	if b == 0 {
		w.Write([]byte("Cannot divide by 0"))
		return
	}

	w.Write([]byte(strconv.FormatFloat(a/b, 'f', -1, 64)))
}
