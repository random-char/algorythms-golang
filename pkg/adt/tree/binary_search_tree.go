package tree

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewBST() BinarySearchTree {
	return BinarySearchTree{}
}

func (bst *BinarySearchTree) AddValue(value int) {
	newNode := Node{
		value: value,
	}

	if bst.root == nil {
		bst.root = &newNode
	} else {
		bst.root.addChild(&newNode)
	}
}

func (n *Node) addChild(newNode *Node) {
	if n.value > newNode.value {
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.addChild(newNode)
		}
	} else {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.addChild(newNode)
		}
	}
}
