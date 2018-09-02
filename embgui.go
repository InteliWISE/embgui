package embgui

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"html"
	"strings"
)

// Tile represents single tile
// https://bulma.io/documentation/layout/tiles/
// see GenTiles()
type Tile struct {
	Title    string
	Subtitle string
}

// EmbNode is a HTML element
type EmbNode struct {
	Text        string
	Href        string
	Action      string
	Method      string
	HTMLTag     string
	Type        string
	Class       string
	Style       string
	Name        string
	ID          string
	Placeholder string
	Value       string
	Unsafe      bool
	Root        bool
	menuOption  string
	GUIConfig   *EmbGUI
	Children    []*EmbNode
}

// EmbGUI is a HTML page, with one root EmbNode with many children
type EmbGUI struct {
	CSS      string
	Size     string
	NavTheme string
	NavLink  string
	title    string
	cssLink  string
	menu     []MenuItem
}

// MenuItem is an singe item in the top menu
type MenuItem struct {
	Name string
	Link string
}

// New creates a new page with a tile and a side menu
// usually it's defined once as a project-wide template
// you may customize NavTheme and NavLink after creation
// NavTheme uses bulma's colors
// (see https://bulma.io/documentation/elements/button/#colors)
// NavLink is a URL for navbar's title
func New(title string, cssLink string, menu []MenuItem) (*EmbGUI, error) {
	gui := EmbGUI{title: title,
		menu:     menu,
		NavTheme: "is-white",
		Size:     "extra-large",
		NavLink:  "/",
		cssLink:  cssLink}
	sDec, _ := b64.StdEncoding.DecodeString(cssFile)
	gui.CSS = string(sDec)
	return &gui, nil
}

// NewRoot creates a new EmbNode root node that belongs to a EmbGUI Page
// it's injected into template, so usually you need to create one root per view
// menuOption will be made active on navbar
func (gui *EmbGUI) NewRoot(menuOption string) *EmbNode {
	return &EmbNode{Root: true, GUIConfig: gui, menuOption: menuOption}
}

// attr renders HTML HTMLTag attribute
func attr(name string, value string, buffer *strings.Builder) {
	if value == "" {
		return
	}
	buffer.WriteString(" ")
	buffer.WriteString(name)
	buffer.WriteString("='")
	buffer.WriteString(html.EscapeString(value))
	buffer.WriteString("'")
}

// startHTMLTag generates begging of HTML HTMLTag
func (n *EmbNode) startHTMLTag() string {
	var buffer strings.Builder
	buffer.WriteString("<")
	buffer.WriteString(n.HTMLTag)
	attr("class", n.Class, &buffer)
	attr("href", n.Href, &buffer)
	attr("id", n.ID, &buffer)
	attr("action", n.Action, &buffer)
	attr("method", n.Method, &buffer)
	attr("type", n.Type, &buffer)
	attr("style", n.Style, &buffer)
	attr("name", n.Name, &buffer)
	attr("value", n.Value, &buffer)
	attr("placeholder", n.Placeholder, &buffer)
	buffer.WriteString(">")
	return buffer.String()
}

// endHTMLTag ends HTML HTMLTag
func (n *EmbNode) endHTMLTag() string {
	var buffer strings.Builder
	buffer.WriteString("</")
	buffer.WriteString(n.HTMLTag)
	buffer.WriteString(">")
	return buffer.String()
}

// genMenu renders top menu inside the navbar
func (n *EmbNode) genMenu(menuOption string) string {
	var buffer strings.Builder
	for _, item := range n.GUIConfig.menu {
		if menuOption == item.Name {
			buffer.WriteString(`<a class="navbar-item is-active" href="`)
		} else {
			buffer.WriteString(`<a class="navbar-item" href="`)
		}
		buffer.WriteString(html.EscapeString(item.Link))
		buffer.WriteString(`">`)
		buffer.WriteString(html.EscapeString(item.Name))
		buffer.WriteString(`</a>`)
	}
	return buffer.String()
}

// renderRoot HTML root element and its child nodes
func (n *EmbNode) renderRoot() string {
	var buffer strings.Builder
	for _, child := range n.Children {
		buffer.WriteString(child.render())
	}
	return buffer.String()
}

// add adds a child to a node
func (n *EmbNode) add(node *EmbNode) *EmbNode {
	n.Children = append(n.Children, node)
	return node
}

// Render HTML element and its child nodes
// if you don't want to escape text inside the tag, set Unsafe to true
func (n *EmbNode) render() string {
	var buffer strings.Builder
	buffer.WriteString(n.startHTMLTag())
	if n.Unsafe == true {
		buffer.WriteString(n.Text)
	} else {
		buffer.WriteString(html.EscapeString(n.Text))
	}
	for _, child := range n.Children {
		buffer.WriteString(child.render())
	}
	buffer.WriteString(n.endHTMLTag())
	return buffer.String()
}

// RenderPage renders template with top-menu, root EmbNode element and its children
func (n *EmbNode) RenderPage() (string, error) {
	if n.Root == false {
		return "", errors.New("can't render page at non-root element")
	}
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<title>%s</title>
			<link rel="stylesheet" href="%s">
		</head>
		<body>
			<nav class="navbar %s">
				<div class="container">
					<div class="navbar-brand">
						<a class="navbar-item brand-text" href="%s">
							%s
						</a>
					</div>
					<div id="navMenu" class="navbar-menu is-active">
						<div class="navbar-start">
							%s
						</div>
					</div>
				</div>
			</nav>
			<section class="section">
				<div class="container">
					<div class="content">
						%s
					</div>
				</div>
			</section>
		</body>
	</html>
	`, n.GUIConfig.title,
		n.GUIConfig.cssLink,
		n.GUIConfig.NavTheme,
		n.GUIConfig.NavLink,
		n.GUIConfig.title,
		n.genMenu(n.menuOption),
		n.renderRoot()), nil
}
