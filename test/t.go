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
	doc, err := xmlquery.Parse(strings.NewReader(ttt))
	if err != nil {
		log.Fatalln(err)
	}

	// html, err := xmlquery.Query(doc, "/html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	body, err := xmlquery.Query(doc, "/html/body")
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

	fmt.Println(doc.OutputXML(false))

	co, err := godirtyxmlquery.CopyLeaf(body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(co.OutputXML(false))
}
