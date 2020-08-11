package godirtyxmlquery

import (
	"encoding/xml"

	"github.com/antchfx/xmlquery"
)

func InsertSubjectBeforeTarget(subject, target *xmlquery.Node) error {
	subject.Parent = target.Parent

	subject.PrevSibling = target.PrevSibling

	target.PrevSibling = subject
	subject.NextSibling = target

	if subject.PrevSibling == nil {
		if subject.Parent != nil {
			subject.Parent.FirstChild = subject
		}
	}

	return nil
}

func InsertSubjectAfterTarget(subject, target *xmlquery.Node) error {
	subject.Parent = target.Parent

	subject.NextSibling = target.NextSibling

	target.NextSibling = subject
	subject.PrevSibling = target

	if subject.NextSibling == nil {
		if subject.Parent != nil {
			subject.Parent.LastChild = subject
		}
	}

	return nil
}

func InsertSubjectBeforeTargetChildren(subject, target *xmlquery.Node) error {

	if target.FirstChild != nil {
		return InsertSubjectBeforeTarget(subject, target.FirstChild)
	}

	subject.Parent = target
	target.FirstChild = subject
	target.LastChild = subject

	return nil
}

func InsertSubjectAfterTargetChildren(subject, target *xmlquery.Node) error {
	if target.LastChild != nil {
		return InsertSubjectAfterTarget(subject, target.LastChild)
	}
	return InsertSubjectBeforeTargetChildren(subject, target)
}

func RemoveSubjectFromItsTree(subject *xmlquery.Node) error {
	if subject.PrevSibling == nil {
		if subject.Parent != nil {
			subject.Parent.FirstChild = subject.NextSibling
		}
	} else {
		subject.PrevSibling.NextSibling = subject.NextSibling
	}

	if subject.NextSibling == nil {
		if subject.Parent != nil {
			subject.Parent.LastChild = subject.PrevSibling
		}
	} else {
		subject.NextSibling.PrevSibling = subject.PrevSibling
	}

	subject.Parent = nil
	subject.PrevSibling = nil
	subject.NextSibling = nil
	return nil
}

func CopyBranch(subject *xmlquery.Node) (*xmlquery.Node, error) {

	ret := &xmlquery.Node{
		Type:         subject.Type,
		Data:         subject.Data,
		Prefix:       subject.Prefix,
		NamespaceURI: subject.NamespaceURI,
	}

	attr := make([]xml.Attr, len(subject.Attr))
	copy(attr, subject.Attr)
	ret.Attr = attr

	t := subject.FirstChild
	for {
		if t == nil {
			break
		}

		lc, err := CopyBranch(t)
		if err != nil {
			return nil, err
		}

		err = InsertSubjectAfterTargetChildren(lc, ret)
		if err != nil {
			return nil, err
		}

		t = t.NextSibling
	}

	ret.Parent = nil

	return ret, nil
}

/*
You can wrap Naive calls with func() { defer func() {x:=recover()}() } to ease
null pointers pain
*/
type NaiveEditTool struct {
	Node *xmlquery.Node
}

func (self *NaiveEditTool) AppendChild(node *xmlquery.Node) *NaiveEditTool {
	InsertSubjectAfterTargetChildren(node, self.Node)
	return &NaiveEditTool{Node: node}
}

func (self *NaiveEditTool) AppendSibling(node *xmlquery.Node) *NaiveEditTool {
	InsertSubjectAfterTarget(node, self.Node)
	return &NaiveEditTool{Node: node}
}

func (self *NaiveEditTool) PrependChild(node *xmlquery.Node) *NaiveEditTool {
	InsertSubjectBeforeTargetChildren(node, self.Node)
	return &NaiveEditTool{Node: node}
}

func (self *NaiveEditTool) PrependSibling(node *xmlquery.Node) *NaiveEditTool {
	InsertSubjectBeforeTarget(node, self.Node)
	return &NaiveEditTool{Node: node}
}
