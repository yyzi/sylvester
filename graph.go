package sylvester

import ()

type Graph struct {
	id       []byte
	nodes    []*Node
	edges    []*Edge
	nodemap  map[*[]byte]*Node
	edgemap  map[*[]byte]*Edge
	IDGen    chan []byte
	ExitChan chan bool
}

func (g *Graph) Id() *[]byte {
	return &g.id
}

func (g *Graph) Activate() {
	for _, node := range g.nodes {
		node.Activate()
	}
	for _, edge := range g.edges {
		edge.Activate()
	}
}

func (g *Graph) NewEdge(anode,bnode *Node) *Edge {
	edge := NewEdge(anode, bnode)

	g.edges = append(g.edges, edge)
	g.edgemap[edge.Id()] = edge

	return edge
}

func (g *Graph) NewEdges(anode *Node, bnodes []*Node) *Edge {
	edge := NewEdges(anode, bnodes)

	g.edges = append(g.edges, edge)
	g.edgemap[edge.Id()] = edge

	return edge
}

func (g *Graph) NewNode() *Node {
	node := NewNode()

	g.nodes = append(g.nodes, node)
	g.nodemap[node.Id()] = node

	return node
}
