package wparser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Html struct {
	Tag       string
	InnerHtml *Html
}

func Parse(filename string) *Html {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("erro: %s\n", err)
		return nil
	}

	content, err := io.ReadAll(file)

	if err != nil {
		fmt.Printf("erro: %s\n", err)
		return nil
	}

	html := strings.TrimSpace(string(content))
	return recursiveParse(html)
}

func recursiveParse(html string) *Html {
	if strings.HasPrefix(html, "<") {
		index := strings.IndexRune(html, '>')
		// index = 5

		// "<body>"
		//  012345
		//   ^--^
		tag := html[1:index]

		postTag := strings.TrimPrefix(html, "<"+tag+">")
		innerHtml := strings.TrimSuffix(postTag, "</"+tag+">")

		htmlObj := Html{
			Tag:       tag,
			InnerHtml: recursiveParse(innerHtml),
		}

		return &htmlObj
	}

	return nil
}

func PrintHtml(html *Html) {
	fmt.Printf("%#v\n", html)
	if html.InnerHtml != nil {
		PrintHtml(html.InnerHtml)
	}
}
