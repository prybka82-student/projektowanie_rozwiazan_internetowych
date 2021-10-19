package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"

	calc "Zadanie03/calc"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/calc/{a:[-]?[0-9]+[.,]?[0-9]*}/plus/{b:[-]?[0-9]+[.,]?[0-9]*}", addHandler).Methods("GET")
	r.HandleFunc("/calc/{a:[-]?[0-9]+[.,]?[0-9]*}/minus/{b:[-]?[0-9]+[.,]?[0-9]*}", subHandler).Methods("GET")
	r.HandleFunc("/calc/{a:[-]?[0-9]+[.,]?[0-9]*}/times/{b:[-]?[0-9]+[.,]?[0-9]*}", multHandler).Methods("GET")
	r.HandleFunc("/calc/{a:[-]?[0-9]+[.,]?[0-9]*}/divide[d]?by/{b:[-]?[0-9]+[.,]?[0-9]*}", divHandler).Methods("GET")

	r.HandleFunc("/calc/{a:[-]?[0-9]+[.,]?[0-9]*}/tothepowerof/{b:[-]?[0-9]+[.,]?[0-9]*}", powHandler).Methods("GET")

	r.HandleFunc("/const/pi", piHandler).Methods("GET")
	r.HandleFunc("/const/2pi", twoPiHandler).Methods("GET")
	r.HandleFunc("/const/pisqrd", piSqrdHandler).Methods("GET")
	r.HandleFunc("/const/e", eHandler).Methods("GET")

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

func getParams(params map[string]string) (string, string) {
	return params["a"], params["b"]
}

func write(w http.ResponseWriter, res string) {
	w.Write([]byte(res))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.Add(getParams(mux.Vars(r)))

	write(w, res)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.Subtract(getParams(mux.Vars(r)))

	write(w, res)
}

func multHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.Multiply(getParams(mux.Vars(r)))

	write(w, res)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.Divide(getParams(mux.Vars(r)))

	write(w, res)
}

func powHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.Power(getParams(mux.Vars(r)))

	write(w, res)
}

func piHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.Pi()

	write(w, res)
}

func twoPiHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.TwoPi()

	write(w, res)
}

func piSqrdHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.PiSqrd()

	write(w, res)
}

func eHandler(w http.ResponseWriter, r *http.Request) {
	res := calc.E()

	write(w, res)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<style>
	page {
		font-family: Arial, Helvetica, sans-serif;
	}
	table {
		width: 100%;
		border: 1px solid #000;
	}
	th, td {
		width: 50%;
		text-align: left;
		vertical-align: top;
		border-spacing: 0;
		padding: 0.3em;
	}
	th {
		border: 1px solid #000;
		font-weight: bold;
	}
	code {
		font-family: Monaco, "Lucida Console", monospace;
	}
	</style>
	<page>
	<h1>Online calculator</h1>

	<h2>Operations</h2>

	<table>
		<tr><th>Operation</th>		<th>URL syntax</th></tr>
		<tr><td>Adding</td>			<td><code>/calc/-123.123/plus/-123.123</code></td></tr>
		<tr><td>Subtracking</td>	<td><code>/calc/-123.123/minus/-123.123</code></td></tr>
		<tr><td>Multiplying</td>	<td><code>/calc/-123.123/times/-123.123</code></td></tr>
		<tr><td>Dividing</td>		<td><code>/calc/-123.123/dividedby/-123.123</code></td></tr>
		<tr><td>Exponentiation</td>	<td><code>/calc/-123.123/tothepowerof/-123.123</code></td></tr>
	</table>

	<h2>Constants</h2>

	<table>
		<tr><th>Constant</th><th>URL syntax</th></tr>
		<tr><td>π</td><td><code>/const/pi</code></td></tr>
		<tr><td>2π</td><td><code>/const/2pi</code></td></tr>
		<tr><td>√π</td><td><code>/const/pisqrd</code></td></tr>
		<tr><td>e</td><td><code>/const/e</code></td></tr>
	</table>
	</page>
	`))
}
