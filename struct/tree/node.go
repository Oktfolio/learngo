package tree

import "fmt"

func main() {
	var root Node
	fmt.Println(root)
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Traverse()

	root.Print()
	fmt.Println()
	root.SetValue(100)
	root.Print()

	var pRoot *Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.Print()

}

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored")
		return
	}
	node.Value = Value
}



func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}