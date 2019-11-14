package cms

import (
	"html/template"
)

//Tmpl is capitalized because it is an exported varible
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title   string
	Content string
}
