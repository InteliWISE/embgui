package main

import (
	"log"
	"net/http"

	"github.com/inteliwise/embgui"
	"github.com/labstack/echo"
)

var ui *embgui.EmbGUI

func main() {
	app := echo.New()
	var err error
	ui, err = embgui.New("EMBDEMO", "/app.css", []embgui.MenuItem{
		{Name: "Hello", Link: "/hello"},
		{Name: "Status", Link: "/status"},
		{Name: "Documentation", Link: "/docs"},
	})
	if err != nil {
		panic("UI failed!")
	}
	ui.NavTheme = "is-light"
	ui.Size = "extra-large"
	ui.NavLink = "/gui"
	app.GET("/", guiIndex)
	app.GET("/app.css", guiCSS)
	log.Fatal(app.Start(":3000"))
}

func guiIndex(c echo.Context) error {
	page := ui.NewRoot("Hello")
	page.H1("Hello 1")
	page.H2("Hello 2")
	page.H3("Hello 3")
	page.H4("Hello 4")
	page.H5("Hello 5")
	page.Hr()
	x := page.P("paragraph with <i>html</i>")
	x.Unsafe = true
	page.RawHTML("<strong>hello raw HTML (or even js!)</strong>")
	page.Hr()
	page.Pre(`<p>some code</p>`, "")
	page.Hr()
	page.Div("somediv", "", "hello")
	page.Hr()
	page.LinkButton("LinkButton", "/docs")
	page.MiniLinkButton("MiniLinkButton", "/docs")
	page.ActionButton("PostButton", "/path")
	col1, col2 := page.TwoColumns()
	col1.P("Maecenas risus quam, ultricies eget ipsum convallis, pellentesque commodo sapien. Mauris orci ligula," +
		"pharetra vitae tincidunt ut, laoreet vitae ligula. Donec tincidunt ipsum leo, eu vulputate purus venenatis eu." +
		"Mauris semper et libero placerat hendrerit. Cras mattis mollis imperdiet. Nam et lectus felis.")
	col2.P("Maecenas risus quam, ultricies eget ipsum convallis, pellentesque commodo sapien. Mauris orci ligula," +
		"pharetra vitae tincidunt ut, laoreet vitae ligula. Donec tincidunt ipsum leo, eu vulputate purus venenatis eu." +
		"Mauris semper et libero placerat hendrerit. Cras mattis mollis imperdiet. Nam et lectus felis.")
	table := page.GenTableBody([]string{"name", "surname"})
	row := table.Tr()
	row.Td("hello")
	row.Td("world")
	row = table.Tr()
	row.Td("john")
	row.Td("smith")
	list := page.Ul()
	list.Li("abc 1")
	list.Li("abc 2")
	page.Message("Something gone wrong!", "is-danger")
	page.Message("OK!", "is-success")
	form := page.Form("/newuser", "POST")
	form.FormInput("First Name", false, "first_name", "")
	form.FormInput("Last Name", false, "last_name", "")
	form.FormTextArea("desc", 10, "some text")
	form.FormButton("Send")
	page.Hr()
	page.H5("Inline buttons")
	buttons := page.Buttons()
	buttons.MiniLinkButton("mini link", "#")
	buttons.MiniDelButton("mini del", "#")
	buttons.MiniActionButton("mini action", "#")
	page.Hr()
	page.SearchForm("/somesearch", "somevalue")
	html, err := page.RenderPage()
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, html)
}

func guiCSS(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentEncoding, "gzip")
	return c.Blob(http.StatusOK, "text/css", []byte(ui.CSS))
}
