package tree

import "fmt"

const ROOT_ID int = 0

// Define the Record type
type Record struct {
	ID, Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	positions := make([]int, len(records))
	for i, r := range records {
		if r.ID < 0 || r.ID >= len(records) {
			return nil, fmt.Errorf("out of range record id %d", r.ID)
		}

		positions[r.ID] = i
	}

	nodes := make([]Node, len(records))
	for i := range positions {
		r := records[positions[i]]
		if r.ID != i {
			return nil, fmt.Errorf("non-contiguous node %d (want: %d)", r.ID, i)
		}
		validParentForChild := (r.ID > r.Parent) || (r.ID == ROOT_ID && r.Parent == ROOT_ID)
		if !validParentForChild {
			return nil, fmt.Errorf("node %d has impossible parent %d", r.ID, r.Parent)
		}

		nodes[i].ID = i
		if i != ROOT_ID {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}

	return &nodes[ROOT_ID], nil
}
