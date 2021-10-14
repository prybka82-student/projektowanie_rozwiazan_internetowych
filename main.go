package main

import (
	"flag"
	"fmt"
	"gophercises/handlers"
	cyoa "gophercises/models"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the CYOA web app on.")
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()

	fmt.Printf("Using the story in %q.\n", *filename)

	file, error := os.Open(*filename)
	if error != nil {
		panic(error)
	}

	story, error := cyoa.JsonStory(file)
	if error != nil {
		panic(error)
	}

	//fmt.Printf("%+v\n", story)

	//how to create a new template:
	//tpl := template.Must(template.New("").Parse("Hello!"))
	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := handlers.NewHandler(story,
		handlers.WithTemplate(tpl),
		handlers.WithPathFunc(pathFn),
	)

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("Starting the server on port: %q\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var storyTmpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose Your Own Adventure</title>
    <link rel="icon" href="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/f77904e3-3600-4909-b6f8-53bdb832bf23/deiagj2-cb6a927e-6ba9-4790-a592-bfedb6287e91.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcL2Y3NzkwNGUzLTM2MDAtNDkwOS1iNmY4LTUzYmRiODMyYmYyM1wvZGVpYWdqMi1jYjZhOTI3ZS02YmE5LTQ3OTAtYTU5Mi1iZmVkYjYyODdlOTEucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.he2bXS9eOmRv8eODtNKkqYmvczDac4_zdtqJSa0ykmk">
</head>
<body>
    <section class="page">
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
            <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </section>
    <style>
        body {
            font-family: helvetica, arial;
        }
        h1 {
            text-align: center;
            position: relative;
        }
        .page {
            width: 80%;
            max-width: 500px;
            margin: auto;
            margin-top: 40px;
            margin-bottom: 40px;
            padding: 80px;
            background: #FFFCF6;
            border: 1px solid #eee;
            box-shadow: 0 10px 6px -6px #777;
        }
        ul {
            border-top: 1px dotted #ccc;
            padding: 10px 0 0 0;
            -webkit-padding-start: 0;
        }
        li {
            padding-top: 10px;
        }
        a, 
        a:visited {
            text-decoration: none;
            color: #6295b5;
        }
        a:active,
        a:hover {
            color: #7792a2;
        }
        p {
            text-indent: 1em;
        }
    </style>
</body>
</html>`

//go run .\main.go
//go run .\main.go --help
//go run .\main.go --file abc.json
//go build .\main.go
//go build .\main.go; .\main.exe
//go build .\main.go; .\main

//http://localhost:3000/story/

//godoc (Go Documentation Server)
//https://stackoverflow.com/questions/63442354/godoc-command-not-found
//installation
//go get golang.org/x/tools/cmd/godoc
//help
//godoc --help
//run
//godoc -http=:3030 //or any port
