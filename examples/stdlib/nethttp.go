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
	form := page.Form("/newuser", "POST")
	form.FormInput("First Name", false, "first_name", "")
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
		embgui.Tile{Title: "90%", Subtitle: "CPU usage"},
		embgui.Tile{Title: "71%", Subtitle: "disk free"})
	html, _ := page.RenderPage()
	w.Write([]byte(html))
}

// return embbeded CSS data
// this will return gzipped bulma.io CSS that is embbeded into embgui sources
func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Content-Encoding", "gzip")
	w.Write([]byte(ui.CSS))
}

func main() {
	// tamplate, shared among all views
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
