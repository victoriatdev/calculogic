package logic

type Node struct {
	data     string
	children []*Node
}

type Tree struct {
	root *Node
}

func add(data string) *Node {
	return &Node{
		data:     data,
		children: make([]*Node, 0),
	}
}
