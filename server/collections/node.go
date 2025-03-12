package collections

type Node struct {
	id       int
	parent   *Node
	data     interface{}
	children []*Node
}

func (n *Node) GetID() int {
	return n.id
}

func (n *Node) GetParent() Node {
	return *n.parent
}

func (n *Node) SetParent(parent *Node) {
	if parent == nil || parent.id == n.id {
		return
	}

	n.parent = parent
}

func (n *Node) GetData() interface{} {
	return n.data
}

func (n *Node) SetData(newData interface{}) {
	n.data = newData
}

func (n *Node) GetChildren() []*Node {
	return n.children
}

func (n *Node) AddChildren(children []*Node) {
	if n.children == nil {
		n.children = []*Node{}
	}
	n.children = append(n.children, children...)
}
