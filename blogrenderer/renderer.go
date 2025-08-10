package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

//go:embed "templates/*"
var postTemplate embed.FS

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
