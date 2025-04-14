package collections

import "github.com/google/uuid"

type Node struct {
	Id       uuid.UUID
	// Parent   *Node
	Data     interface{}
	Children []*Node
	Rule 	 interface{}
}

func (n *Node) GetID() uuid.UUID {
	return n.Id
}

// func (n *Node) GetParent() Node {
// 	return *n.Parent
// }

// func (n *Node) SetParent(parent *Node) {
// 	if parent == nil || parent.Id == n.Id {
// 		return
// 	}

// 	n.Parent = parent
// }

func (n *Node) GetData() interface{} {
	return n.Data
}

func (n *Node) SetData(newData interface{}) {
	n.Data = newData
}

func (n *Node) GetChildren() []*Node {
	return n.Children
}

func (n *Node) AddChildren(children []*Node) {
	if n.Children == nil {
		n.Children = []*Node{}
	}
	n.Children = append(n.Children, children...)
}

func (n *Node) AddChild(children *Node) {
	if n.Children == nil {
		n.Children = []*Node{}
	}
	n.Children = append(n.Children, children)
}

func (n *Node) SetRule(rule interface{}) {
	n.Rule = rule;
}