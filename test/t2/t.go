package main

import (
	"log"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	ttt := `<html><body><p></p></body></html>`

	doc, err := xmlquery.Parse(strings.NewReader(ttt))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := xmlquery.Query(doc, "/html/body")
	if err != nil {
		log.Fatalln(err)
	}

	if body == nil {
		log.Fatalln("no body")
	}

	p, err := xmlquery.Query(body, "/p")
	if err != nil {
		log.Fatalln(err)
	}

	if p == nil {
		log.Fatalln("no p")
	}

}
