package hackerrank

type Node struct {
	val         int
	left, right *Node
}

type TreeNode struct {
	root *Node
}

func (tn *TreeNode) Insert(val int) {
	if tn.root == nil {
		tn.root = &Node{val, nil, nil}
	} else {
		if tn.root.val < val {
			tn.root.right = &Node{val, nil, nil}
		} else {
			tn.root.left = &Node{val, nil, nil}
		}
	}
}

func (n *Node) Insert(val int) {
	if n.val < val {
		n.right = &Node{val, nil, nil}
	} else {
		n.left = &Node{val, nil, nil}
	}
}

func (tn *TreeNode) IsSymmetric() bool {
	return false
}

func (n *Node) IsSymmetric() bool {
	return false
}
