package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", post.Title)
	return err
}
