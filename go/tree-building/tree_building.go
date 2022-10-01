package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
	Parent int
}

const ROOT int = 0

type MapRecord map[int]int

func Build(records []Record) (*Node, error) {
	var length int = len(records)
	if length == 0 {
		return nil, nil
	}

	// descending sort
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	var mr MapRecord = make(MapRecord, length)
	var maxId int = ROOT

	for _, record := range records {
		if record.ID == ROOT && record.Parent > ROOT {
			return nil, errors.New("root node has parent")
		}

		if record.Parent != ROOT && record.Parent >= record.ID {
			return nil, errors.New("parent id higher than id")
		}

		if record.Parent != ROOT && record.Parent == record.ID {
			return nil, errors.New("node cycle directly")
		}

		if parent, ok := mr[record.ID]; ok {
			// same record found check if 0 or other
			if parent == ROOT {
				return nil, errors.New("duplicate root")
			}
			return nil, errors.New("duplicate node")
		}
		mr[record.ID] = record.Parent
		// store max ID
		if record.ID > maxId {
			maxId = record.ID
		}
	}

	var node *Node = &Node{}
	// check if 'root' exist
	if parent, ok := mr[ROOT]; ok {
		// check if id continuos or not
		for maxId >= ROOT {
			if _, found := mr[maxId]; !found {
				return nil, errors.New("non-continuos id")
			}
			maxId--
		}

		node = &Node{
			ID:       ROOT,
			Parent:   parent,
			Children: []*Node{},
		}
	} else {
		return nil, errors.New("no root node")
	}

	for _, record := range records {
		if record.ID == ROOT && record.Parent == ROOT {
			continue
		}

		if AssignRecord(node, record) {
			continue
		}

		return nil, errors.New("parent node not found ix")
	}

	return node, nil
}

func AssignRecord(node *Node, record Record) bool {

	if node.ID == record.Parent {
		node.Children = append(node.Children, &Node{
			ID:       record.ID,
			Parent:   record.Parent,
			Children: []*Node{},
		})
		return true
	}

	if len(node.Children) > 0 {
		for _, v := range node.Children {
			if AssignRecord(v, record) {
				return true
			}
		}
	}

	return false
}
