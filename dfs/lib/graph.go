package dfs

// NodesArray is an alias of array of *Node
type NodesArray []*Node

// Graph represents a graph that maintains its nodes via adjacency list
type Graph struct {
	numOfNodes  int
	ListOfNodes map[*Node]NodesArray
}

// NumOfNodes returns total number of nodes in graph
func (g *Graph) NumOfNodes() int {
	return g.numOfNodes
}

// AddNode adds new edge between srcNode and destNode
func (g *Graph) AddNode(srcNode *Node, destNode *Node) *Graph {
	if g.ListOfNodes[srcNode] == nil {
		g.ListOfNodes[srcNode] = make(NodesArray, 0)
	}

	g.ListOfNodes[srcNode] = append(g.ListOfNodes[srcNode], destNode)

	return g
}

// NewGraph creates and returns a new instance of Graph
func NewGraph() Graph {
	list := make(map[*Node]NodesArray, 0)
	return Graph{
		ListOfNodes: list,
	}
}
