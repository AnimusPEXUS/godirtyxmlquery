package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/AnimusPEXUS/godirtyxmlquery"
	"github.com/antchfx/xmlquery"
)

func main() {
	ttt := `<html><body></body></html>`
	html, err := xmlquery.Parse(strings.NewReader(ttt))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := xmlquery.Query(html, "/html/body")
	if err != nil {
		log.Fatalln(err)
	}

	tool := &godirtyxmlquery.NaiveEditTool{Node: body}

	tool.PrependChild(&xmlquery.Node{
		Type: xmlquery.ElementNode,
		Data: "b",
	}).AppendChild(&xmlquery.Node{
		Type: xmlquery.TextNode,
		Data: "Great Job!",
	})

	tool = &godirtyxmlquery.NaiveEditTool{Node: body}
	tool.AppendChild(&xmlquery.Node{
		Type: xmlquery.ElementNode,
		Data: "queen",
	}).AppendChild(&xmlquery.Node{
		Type: xmlquery.TextNode,
		Data: "I will be King! And You.. You will be Queen!",
	})

	fmt.Println(html.OutputXML(false))
}
