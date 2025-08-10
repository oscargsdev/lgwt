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
	templ, err := template.ParseFS(postTemplate, "templates/blog.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
