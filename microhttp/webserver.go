// Package webserver static HTTP server
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"bytes"

	chromahtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func setContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Content-Language", "en-US")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	filename := r.URL.Path[len("/"):]
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Start serving files in", filename)
	}

	setContent(w)
	w.Write(body)
}

func markdownHandler(w http.ResponseWriter, r *http.Request) {

	filename := r.URL.Path[len("/md/"):]

	md, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Serving file", filename)
	}

	//extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.Mmark | SuperSubscript
	//parser := parser.NewWithExtensions(extensions)

	setContent(w)
	//html := markdown.ToHTML(md, parser, nil)
	//html := blackfriday.Run(md)

	hl := highlighting.NewHighlighting(
		highlighting.WithStyle("monokai"),
		highlighting.WithGuessLanguage(true),
		highlighting.WithFormatOptions(
			chromahtml.WithLineNumbers(true),
			chromahtml.LineNumbersInTable(true),
			//chromahtml.PreventSurroundingPre(true),
			//chromahtml.WithClasses(true),
		),
	)

	gm := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Typographer, hl),

		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := gm.Convert(md, &buf); err != nil {
		panic(err)
	}

	w.Write(buf.Bytes())
}

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Start serving files in", path)
	}

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/md/", markdownHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	print("Exiting server")
	os.Exit(0)
}
