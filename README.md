## EmbGUI - embedded GUI for Go apps. 

[![Go Report Card](https://goreportcard.com/badge/github.com/inteliwise/embgui)](https://goreportcard.com/report/github.com/inteliwise/embgui)

Instant admin panels for Go apps. Turn plain Go code into HTML+CSS admin panel with a collection of common components and a predefined layout.

<img src="https://raw.githubusercontent.com/inteliwise/embgui/master/examples/screenshot.png" alt="screenshot" style="max-width:100%;"></a>

## Use case

* EmbGUI is aimed primarily at microservices and APIs, that need to expose some stats or basic maintenance panel via HTTP HTML interface
* it has all the assets embedded, it's very convenient when shipping a single binary with no external dependencies
* it doesn't require to do any:
	* HTML
	* CSS
	* JS
	* layouts
	* resource embedding

## Non-use cases

* advanced admin panels and websites
* JavaScript apps (there's no JS, just plain old server-side rendering)
* unique, custom styled panels (this lib is highly opinionated and styling is very limited)


```go
package main

import (
	"net/http"
	"github.com/inteliwise/embgui"
)

var ui *embgui.EmbGUI

// sample page
func index(w http.ResponseWriter, r *http.Request) {
	page := ui.NewRoot("Hello") // "Hello" option will be active on navbar
	page.H1("Hello!")
	page.P("Lorem...")
	form := page.Form("/newuser")
	form.FormInput("First Name", "first_name")
	form.FormButton("Send")
	table := page.GenTableBody([]string{"name", "surname", "action"})
	row := table.Tr()
	row.Td("hello")
	row.Td("world")
	row.Td("").LinkButton("Inspect", "#")
	html, _ := page.RenderPage() // skipping errors for a concise example
	w.Write([]byte(html))
}

// another page that shares ui template with index
func world(w http.ResponseWriter, r *http.Request) {
	page := ui.NewRoot("World") // "World" option will be active on the navbar
	page.H1("World!")
	page.GenTiles(embgui.Tile{Title: "7", Subtitle: "new users"},
		embgui.Tile{Title: "71", Subtitle: "new sales"},
		embgui.Tile{Title: "90", Subtitle: "CPU usage"},
		embgui.Tile{Title: "71", Subtitle: "disk free"})	
	tagWithUnsafeContent := page.P("<strong>hello!</strong>")
	tagWithUnsafeContent.Unsafe = true
	page.RawHTML("<p><i>hello world</i></p>")
	html, _ := page.RenderPage()
	w.Write([]byte(html))
}

// return embedded CSS data
// this will return gzipped bulma.io CSS that is embedded into embgui sources
func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Content-Encoding", "gzip")
	w.Write([]byte(ui.CSS)) 
}

func main() {
	// template, shared among all views
	// remember, that you need to pass a link to CSS assets
	ui, _ = embgui.New("DEMO", "/app.css", []embgui.MenuItem{
		{Name: "Hello", Link: "/"},
		{Name: "World", Link: "/world"},
	})
	http.HandleFunc("/", index)
	http.HandleFunc("/world", world)
	// CSS assets used by embgui.New()
	http.HandleFunc("/app.css", css)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

Check the `examples` directory for a [Echo](https://echo.labstack.com/) framework sample app.

This is still WIP! Things may change before it reaches 1.0.0!

## Developent goals

* keep the ease of use, it always has to be from 0 to GUI in minutes
* no JavaScript, CSS only, keep everything server-side
* don't introduce any external dependencies like web-fonts or APIs (it should work in environments with no public Internet access)
* maintain compatibility with text-based browsers 
* add charts (server side rendered)
* more [components](https://bulma.io/documentation/components/) and [elements](https://bulma.io/documentation/elements/)! 

## Bulma

EmbGUI uses wonderful [Bulma](https://bulma.io/) framework for layout and styling. It's gzipped and base64 encoded and will be compiled with your app. Don't worry, it's just 30kb.
Bulma Framework is released at https://github.com/jgthms/bulma on MIT License (see `css_file.go`)
