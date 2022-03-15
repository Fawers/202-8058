package main

import (
	"parser/wparser"
)

func main() {
	html := wparser.Parse("exemplo.html")
	wparser.PrintHtml(html)
}
