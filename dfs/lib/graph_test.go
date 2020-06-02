package dfs

import "testing"

func TestNewGraph(t *testing.T) {
	g := NewGraph()

	t.Run("NumOfNodes() should return 0", func(t *testing.T) {
		if g.NumOfNodes() != 0 {
			t.Error("Expected NumOfNodes() to return 0")
		}
	})

	t.Run("ListOfNodes should not be nil", func(t *testing.T) {
		if g.ListOfNodes == nil {
			t.Error("Expected ListOfNodes to be not nil")
		}
	})
}

func TestAddNode(t *testing.T) {
	g := NewGraph()
	nodeA := Node{1}

	t.Run("AddNode() should add new link between source node and destination node", func(t *testing.T) {
		nodeB := Node{3}
		g.AddNode(&nodeA, &nodeB)

		if g.ListOfNodes[&nodeA][0] != &nodeB {
			t.Error("Expected nodeB to be linked to nodeA")
		}

		nodeE := Node{4}
		g.AddNode(&nodeB, &nodeE)

		if g.ListOfNodes[&nodeB][0] != &nodeE {
			t.Error("Expected nodeB to be linked to nodeA")
		}
	})

	t.Run("AddNode() should be chainable", func(t *testing.T) {
		nodeC := Node{5}
		nodeD := Node{6}
		g.AddNode(&nodeA, &nodeC).AddNode(&nodeA, &nodeD)

		if g.ListOfNodes[&nodeA][1] != &nodeC {
			t.Error("Expected nodeC to be linked to nodeA")
		}

		if g.ListOfNodes[&nodeA][2] != &nodeD {
			t.Error("Expected nodeD to be linked to nodeA")
		}
	})
}
