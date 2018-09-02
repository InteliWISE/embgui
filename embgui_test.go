package embgui

import (
	"strings"
	"testing"
)

type simpleElementTest struct {
	desc           string
	renderFunc     func(page *EmbNode) string
	expectedResult string
}

var simpleElementTests = []simpleElementTest{
	{"Hr", func(page *EmbNode) string { return page.Hr().render() }, "<hr class='hr'></hr>"},
	{"H1", func(page *EmbNode) string { return page.H1("H1").render() }, "<h1 class='title is-1'>H1</h1>"},
	{"H2", func(page *EmbNode) string { return page.H2("H2").render() }, "<h2 class='title is-2'>H2</h2>"},
	{"H3", func(page *EmbNode) string { return page.H3("H3").render() }, "<h3 class='title is-3'>H3</h3>"},
	{"H4", func(page *EmbNode) string { return page.H4("H4").render() }, "<h4 class='title is-4'>H4</h4>"},
	{"H5", func(page *EmbNode) string { return page.H5("H5").render() }, "<h5 class='title is-5'>H5</h5>"},
	{"Message", func(page *EmbNode) string { return page.Message("hello", "is-primary").render() }, "<div class='message-body'>hello</div>"},
	{"Pre", func(page *EmbNode) string { return page.Pre("lorem ipsum", "").render() }, "<pre>lorem ipsum</pre>"},
	{"P", func(page *EmbNode) string { return page.P("lorem ipsum").render() }, "<p>lorem ipsum</p>"},
	{"Div", func(page *EmbNode) string { return page.Div("some_id", "border: 1px solid black;", "hello").render() }, "<div id='some_id' style='border: 1px solid black;'>hello</div>"},
	{"Box", func(page *EmbNode) string { return page.Box().render() }, `<div class='box'></div>`},
	{"A", func(page *EmbNode) string { return page.A("link text", "https://example.com").render() }, "<a href='https://example.com'>link text</a>"},
	{"LinkButton", func(page *EmbNode) string {
		return page.LinkButton("button text", "https://example.com").render()
	}, "<a class='button is-link' href='https://example.com' style='margin: .25rem'>button text</a>"},
	{"MiniLinkButton", func(page *EmbNode) string {
		return page.MiniLinkButton("mini button text", "https://example.com").render()
	}, "<a class='button is-link is-small' href='https://example.com' style='margin: .25rem'>mini button text</a>"},
	{"ActionButton", func(page *EmbNode) string {
		return page.ActionButton("link text", "https://example.com").render()
	}, "<form action='https://example.com' method='POST'><button class='button is-primary' type='submit' style='margin: .25rem'>link text</button></form>"},
}

func preparePage() *EmbNode {
	ui, err := New("EMBDEMO", "/app.css", []MenuItem{
		{Name: "Index", Link: "/"},
		{Name: "Status", Link: "/status"},
		{Name: "Documentation", Link: "/docs"},
	})
	if err != nil {
		return nil
	}
	return ui.NewRoot("Index")
}

func TestBasic(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	for _, pair := range simpleElementTests {
		v := pair.renderFunc(page)
		if v != pair.expectedResult {
			t.Error(
				"For", pair.desc,
				"expected", pair.expectedResult,
				"got", v,
			)
		}
	}
}

func TestTable(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	table := page.GenTableBody([]string{"name", "surname"})
	row := table.Tr()
	row.Td("hello")
	row.Td("world")
	row = table.Tr()
	row.Td("john")
	row.Td("smith")
	v := table.render()
	expectedResult := `<tbody><tr><td>hello</td><td>world</td></tr><tr><td>john</td><td>smith</td></tr></tbody>`
	if v != expectedResult {
		t.Error(
			"For", "TestTable",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestForm(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	form := page.Form("/newuser", "POST")
	form.FormInput("First Name", true, "first_name", "")
	form.FormInput("Last Name", false, "surname", "Smith")
	form.FormButton("Send")
	v := form.render()
	expectedResult := `<form action='/newuser' method='POST'><div class='field'><div class='control'>` +
		`<input class='input' type='text' name='first_name' placeholder='First Name'></input></div></div>` +
		`<div class='field'><label class='label'>Last Name</label><div class='control'>` +
		`<input class='input' type='text' name='surname' value='Smith'></input></div></div>` +
		`<div class='control'><button class='button' type='sumbit'>Send</button></div></form>`
	if v != expectedResult {
		t.Error(
			"For", "TestForm",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestSearchForm(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	form := page.SearchForm("/users", "")
	v := form.render()
	expectedResult := `<form action='/users' method='GET'><div class='field has-addons'><div class='control'><input class='input' ` +
		`type='text' name='search'></input></div><div class='control'><button class='button is-info' type='sumbit'>Search</button></div></div></form>`
	if v != expectedResult {
		t.Error(
			"For", "TestSearchForm",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestList(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	list := page.Ul()
	list.Li("stuff")
	list.Li("more stuff")
	v := list.render()
	expectedResult := `<ul><li>stuff</li><li>more stuff</li></ul>`
	if v != expectedResult {
		t.Error(
			"For", "TestList",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestTiles(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	tiles := page.GenTiles(Tile{Title: "7", Subtitle: "new users"},
		Tile{Title: "71", Subtitle: "new sales"},
		Tile{Title: "90%", Subtitle: "CPU usage"},
		Tile{Title: "71%", Subtitle: "disk free"})
	v := tiles.render()
	expectedResult := `<div class='tile is-ancestor'><div class='tile is-parent'><article class='tile is-child box'>` +
		`<p class='title'>7</p><p class='subtitle'>new users</p></article></div><div class='tile is-parent'>` +
		`<article class='tile is-child box'><p class='title'>71</p><p class='subtitle'>new sales</p></article></div>` +
		`<div class='tile is-parent'><article class='tile is-child box'><p class='title'>90%</p><p class='subtitle'>CPU usage</p>` +
		`</article></div><div class='tile is-parent'><article class='tile is-child box'><p class='title'>71%</p>` +
		`<p class='subtitle'>disk free</p></article></div></div>`
	if v != expectedResult {
		t.Error(
			"For", "TestList",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestUnsafeRendering(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	div := page.Div("1", "", "<p>hello</p>")
	div.Unsafe = true
	v := div.render()
	expectedResult := `<div id='1'><p>hello</p></div>`
	if v != expectedResult {
		t.Error(
			"For", "TestUnsafeRendering",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestTwoColumns(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	col1, col2 := page.TwoColumns()
	col1.P("hello")
	col2.P("world")
	v := page.render()
	expectedResult := `<><div class='columns'><div class='column'><p>hello</p></div><div class='column'><p>world</p></div></div></>`
	if v != expectedResult {
		t.Error(
			"For", "TestTwoColumns",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestRawHTML(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	page.RawHTML("<p>raw!</p>")
	v := page.render()
	expectedResult := `<><div><p>raw!</p></div></>`
	if v != expectedResult {
		t.Error(
			"For", "TestRawHTML",
			"expected", expectedResult,
			"got", v,
		)
	}
}

func TestFullRender(t *testing.T) {
	page := preparePage()
	if page == nil {
		t.Errorf("can't initialize test page")
	}
	h1 := page.H1("helloworld")
	_, err := h1.RenderPage()
	if err == nil {
		t.Error("For", "FullRender", "Should return error when Rendering non-root element")
	}
	v, err := page.RenderPage()
	if err != nil {
		t.Error("For", "TestFullRender", "Error:", err.Error())
	}
	testStrings := []string{`<!DOCTYPE html>`,
		`<title>EMBDEMO</title>`,
		`<link rel="stylesheet" href="/app.css">`,
		`<a class="navbar-item brand-text" href="/">`,
		`<div class="content">`}
	for _, str := range testStrings {
		if strings.Contains(v, str) == false {
			t.Error(
				"For", "TestFullRender",
				"expected to have", str,
			)
		}
	}
}
