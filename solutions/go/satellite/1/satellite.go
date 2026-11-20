package satellite

import (
	"errors"
	"slices"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func TreeFromTraversals(preorder, inorder []string) (*Node, error) {
	if len(preorder) != len(inorder) {
		return nil, errors.New("traversals must have the same length")
	}

	seen := make(map[string]struct{})
	return collect(seen, preorder, inorder)
}

func collect(seen map[string]struct{}, preorder, inorder []string) (*Node, error) {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil, nil
	}

	root := preorder[0]
	if _, dup := seen[root]; dup {
		return nil, errors.New("traversals must contain unique items")
	}
	seen[root] = struct{}{}

	i := slices.Index(inorder, root)
	if i < 0 {
		return nil, errors.New("traversals must have the same elements")
	}

	left, err := collect(seen, preorder[1:1+i], inorder[:i])
	if err != nil {
		return nil, err
	}

	right, err := collect(seen, preorder[1+i:], inorder[i+1:])
	if err != nil {
		return nil, err
	}

	return &Node{Value: root, Left: left, Right: right}, nil
}
