package main

import (
	"os"
	"time"

	"github.com/oscargsdev/lgwt/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWritter(os.Stdout, t)
}
