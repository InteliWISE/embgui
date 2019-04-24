package embgui

// H1 generates <h1> tag
func (n *EmbNode) H1(text string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "h1", Class: "title is-1"})
}

// H2 generates <h2> tag
func (n *EmbNode) H2(text string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "h2", Class: "title is-2"})
}

// H3 generates <h3> tag
func (n *EmbNode) H3(text string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "h3", Class: "title is-3"})
}

// H4 generates <h4> tag
func (n *EmbNode) H4(text string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "h4", Class: "title is-4"})
}

// H5 generates <h5> tag
func (n *EmbNode) H5(text string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "h5", Class: "title is-5"})
}

// Pre generates <pre><code> tags
func (n *EmbNode) Pre(text string, class string) *EmbNode {
	return n.add(&EmbNode{HTMLTag: "pre", Class: class, Text: text})
}

// P generates <p> tag
func (n *EmbNode) P(text string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "p"})
}

// Div generates <div> with custom id and styling
func (n *EmbNode) Div(id string, style string, text string) *EmbNode {
	return n.add(&EmbNode{ID: id, HTMLTag: "div", Style: style, Text: text})
}

// Box generates div-box element
func (n *EmbNode) Box() *EmbNode {
	return n.add(&EmbNode{HTMLTag: "div", Class: "box"})
}

// TwoColumns generates 2 column layout and returns 2 EmbNode pointers
//
//		col1, col2 := TwoColumns()
//		col1.H1("hello")
//		col2.H2("world")
func (n *EmbNode) TwoColumns() (*EmbNode, *EmbNode) {
	parent := n.add(&EmbNode{HTMLTag: "div", Class: "columns"})
	return parent.add(&EmbNode{HTMLTag: "div", Class: "column"}),
		parent.add(&EmbNode{HTMLTag: "div", Class: "column"})
}

// GenTiles generates set of tiles
//
//		page.GenTiles(embgui.Tile{Title: "7", Subtitle: "new users"},
//		embgui.Tile{Title: "71", Subtitle: "new sales"},
//		embgui.Tile{Title: "90%", Subtitle: "CPU usage"},
//		embgui.Tile{Title: "71%", Subtitle: "disk free"})
func (n *EmbNode) GenTiles(data ...Tile) *EmbNode {
	parent := n.add(&EmbNode{HTMLTag: "div", Class: "tile is-ancestor"})
	for _, n := range data {
		tile := parent.add(&EmbNode{HTMLTag: "div", Class: "tile is-parent"})
		article := tile.add(&EmbNode{HTMLTag: "article", Class: "tile is-child box"})
		article.add(&EmbNode{HTMLTag: "p", Class: "title", Text: n.Title})
		article.add(&EmbNode{HTMLTag: "p", Class: "subtitle", Text: n.Subtitle})
	}
	return parent
}

// A generates a link
func (n *EmbNode) A(text string, href string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "a", Href: href})
}

// Buttons generates <div> for button group
func (n *EmbNode) Buttons() *EmbNode {
	return n.add(&EmbNode{HTMLTag: "div", Class: "buttons"})
}

// LinkButton generates a link styled as button
func (n *EmbNode) LinkButton(text string, href string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "a", Href: href, Class: "button is-link", Style: "margin: .25rem"})
}

// ActionButton generates POST button wrapped into a hidden form
// It's like LinkButton, but uses POST instead of GET
func (n *EmbNode) ActionButton(text string, action string) *EmbNode {
	form := &EmbNode{HTMLTag: "form", Action: action, Method: "POST"}
	form.add(&EmbNode{HTMLTag: "button", Type: "submit", Text: text, Class: "button is-primary", Style: "margin: .25rem"})
	return n.add(form)
}

// DelButton generates DEL button wrapped into a hidden form
// your framework should support hidden _method tag
func (n *EmbNode) DelButton(text string, action string) *EmbNode {
	form := &EmbNode{HTMLTag: "form", Action: action, Method: "POST"}
	form.add(&EmbNode{HTMLTag: "input", Type: "hidden", Name: "_method", Value: "DELETE"})
	form.add(&EmbNode{HTMLTag: "button", Type: "submit", Text: text, Class: "button is-danger", Style: "margin: .25rem"})
	return n.add(form)
}

// MiniLinkButton generates a link styled as button, just like LinkButton(), but smaller
// it's good for links inside tables
func (n *EmbNode) MiniLinkButton(text string, href string) *EmbNode {
	return n.add(&EmbNode{Text: text, HTMLTag: "a", Href: href, Class: "button is-link is-small", Style: "margin: .25rem"})
}

// MiniActionButton generates POST button wrapped into a hidden form
// it's good for buttons inside tables
func (n *EmbNode) MiniActionButton(text string, action string) *EmbNode {
	form := &EmbNode{HTMLTag: "form", Action: action, Method: "POST"}
	form.add(&EmbNode{HTMLTag: "button", Type: "submit", Text: text, Class: "button is-primary is-small", Style: "margin: .25rem"})
	return n.add(form)
}

// MiniDelButton generates DEL button wrapped into a hidden form
// it's good for buttons inside tables
// your framework should support hidden _method tag
func (n *EmbNode) MiniDelButton(text string, action string) *EmbNode {
	form := &EmbNode{HTMLTag: "form", Action: action, Method: "POST"}
	form.add(&EmbNode{HTMLTag: "input", Type: "hidden", Name: "_method", Value: "DELETE"})
	form.add(&EmbNode{HTMLTag: "button", Type: "submit", Text: text, Class: "button is-danger is-small", Style: "margin: .25rem"})
	return n.add(form)
}

// Message generates pre-styled message/tip div - alerts etc
// color can be one of bulma's colors (see https://bulma.io/documentation/elements/button/#colors)
func (n *EmbNode) Message(text string, color string) *EmbNode {
	msg := n.add(&EmbNode{HTMLTag: "div", Class: "message " + color})
	return msg.add(&EmbNode{HTMLTag: "div", Text: text, Class: "message-body"})
}

// GenTableBody helps generate html table
// it returns tbody from inside the table, not the table itself, so you can add Tr/Td to it
//
//		tbody := page.GenTableBody([]string{"id", "name", "surname"})
// 		row := tbody.Tr()
// 		row.Td("1")
// 		row.Td("john")
// 		row.Td("smith")
func (n *EmbNode) GenTableBody(header []string) *EmbNode {
	table := &EmbNode{HTMLTag: "table", Class: "table is-narrow is-hoverable is-fullwidth"}
	head := table.add(&EmbNode{HTMLTag: "thead"})
	headRow := head.add(&EmbNode{HTMLTag: "tr"})
	for _, h := range header {
		headRow.add(&EmbNode{HTMLTag: "th", Text: h})
	}
	tbody := table.add(&EmbNode{HTMLTag: "tbody"})
	n.add(table)
	return tbody
}

// Tr table element
func (n *EmbNode) Tr() *EmbNode {
	return n.add(&EmbNode{HTMLTag: "tr"})
}

// Td table element
func (n *EmbNode) Td(text string) *EmbNode {
	return n.add(&EmbNode{HTMLTag: "td", Text: text})
}

// Ul starts a list
func (n *EmbNode) Ul() *EmbNode {
	return n.add(&EmbNode{HTMLTag: "ul"})
}

// Li adds element to list
func (n *EmbNode) Li(text string) *EmbNode {
	return n.add(&EmbNode{HTMLTag: "li", Text: text})
}

// Form generates simple form
func (n *EmbNode) Form(action string, method string) *EmbNode {
	return n.add(&EmbNode{HTMLTag: "form", Action: action, Method: method})
}

// FormInput generates a text input inside form
func (n *EmbNode) FormInput(label string, hideLabel bool, name string, value string) *EmbNode {
	fieldWrapper := n.add(&EmbNode{HTMLTag: "div", Class: "field"})
	if hideLabel == false {
		fieldWrapper.add(&EmbNode{HTMLTag: "label", Class: "label", Text: label})
		fieldWrapper.add(&EmbNode{HTMLTag: "div", Class: "control"}).
			add(&EmbNode{HTMLTag: "input", Type: "text", Class: "input", Name: name, Value: value})
	} else {
		fieldWrapper.add(&EmbNode{HTMLTag: "div", Class: "control"}).
			add(&EmbNode{HTMLTag: "input", Type: "text", Class: "input", Name: name, Placeholder: label, Value: value})
	}
	return fieldWrapper
}

// FileUpload generates a file upload form
func (n *EmbNode) FileUpload(action string, label string, name string) *EmbNode {
	realForm := n.add(&EmbNode{HTMLTag: "form", Action: action, Method: "POST", Enctype: "multipart/form-data"})
	fieldWrapper := realForm.add(&EmbNode{HTMLTag: "div", Class: "field"})
	fieldWrapper.add(&EmbNode{HTMLTag: "label", Class: "label", Text: label})
	fieldWrapper.add(&EmbNode{HTMLTag: "div", Class: "control"}).
		add(&EmbNode{HTMLTag: "input", Type: "file", Class: "input", Name: name})
	return realForm
}

// FormButton generates a submit button inside a form
func (n *EmbNode) FormButton(text string) *EmbNode {
	return n.add(&EmbNode{HTMLTag: "div", Class: "control"}).
		add(&EmbNode{HTMLTag: "button", Type: "sumbit", Class: "button", Text: text})
}

// FormTextArea generates a textarea inside a form
func (n *EmbNode) FormTextArea(name string, rows int, text string) *EmbNode {
	return n.add(&EmbNode{HTMLTag: "div", Class: "field"}).
		add(&EmbNode{HTMLTag: "textarea", Class: "textarea", Name: name, Rows: rows, Text: text})
}

// SearchForm generates predefined search form
// action is a URL we use to do a GET request with a "search" param
func (n *EmbNode) SearchForm(action string, value string) *EmbNode {
	realForm := n.add(&EmbNode{HTMLTag: "form", Action: action, Method: "GET"})
	wrapper := realForm.add(&EmbNode{HTMLTag: "div", Class: "field has-addons"})
	wrapper.
		add(&EmbNode{HTMLTag: "div", Class: "control"}).
		add(&EmbNode{HTMLTag: "input", Type: "text", Class: "input", Name: "search", Value: value})
	wrapper.
		add(&EmbNode{HTMLTag: "div", Class: "control"}).
		add(&EmbNode{HTMLTag: "button", Type: "sumbit", Class: "button is-info", Text: "Search"})
	return realForm
}

// RawHTML generates a div with unsafe HTML
func (n *EmbNode) RawHTML(html string) *EmbNode {
	return n.add(&EmbNode{HTMLTag: "div", Text: html, Unsafe: true})
}

// Hr generates a horizontal rule
func (n *EmbNode) Hr() *EmbNode {
	return n.add(&EmbNode{HTMLTag: "hr", Class: "hr"})
}
