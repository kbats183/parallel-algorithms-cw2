package main

type Graph interface {
	N() int
	Edge(n int) []int
}

type CubeGridGraph struct {
	Side int
}

func (g *CubeGridGraph) N() int {
	return g.Side * g.Side * g.Side
}

func (g *CubeGridGraph) EncodeGraphNode(i int, j int, k int) int {
	return (i*g.Side+j)*g.Side + k
}

func (g *CubeGridGraph) Edge(n int) []int {
	var nb []int

	k := n % g.Side
	j := n / g.Side % g.Side
	i := n / g.Side / g.Side

	if i > 0 {
		nb = append(nb, g.EncodeGraphNode(i-1, j, k))
	}
	if i+1 < g.Side {
		nb = append(nb, g.EncodeGraphNode(i+1, j, k))
	}
	if j > 0 {
		nb = append(nb, g.EncodeGraphNode(i, j-1, k))
	}
	if j+1 < g.Side {
		nb = append(nb, g.EncodeGraphNode(i, j+1, k))
	}
	if k > 0 {
		nb = append(nb, g.EncodeGraphNode(i, j, k-1))
	}
	if k+1 < g.Side {
		nb = append(nb, g.EncodeGraphNode(i, j, k+1))
	}
	return nb
}

func NewInMemoryCubeGridGraph(side int) *InMemoryCubeGridGraph {
	c := side * side * side
	g := InMemoryCubeGridGraph{
		side: side,
		//edges: make([][]int, c),
	}

	for n := 0; n < c; n++ {
		var nb []int
		k := n % side
		j := n / g.side % g.side
		i := n / g.side / g.side

		if i > 0 {
			nb = append(nb, g.EncodeGraphNode(i-1, j, k))
		}
		if i+1 < g.side {
			nb = append(nb, g.EncodeGraphNode(i+1, j, k))
		}
		if j > 0 {
			nb = append(nb, g.EncodeGraphNode(i, j-1, k))
		}
		if j+1 < g.side {
			nb = append(nb, g.EncodeGraphNode(i, j+1, k))
		}
		if k > 0 {
			nb = append(nb, g.EncodeGraphNode(i, j, k-1))
		}
		if k+1 < g.side {
			nb = append(nb, g.EncodeGraphNode(i, j, k+1))
		}
		g.edges = append(g.edges, nb)
	}
	return &g
}

type InMemoryCubeGridGraph struct {
	side  int
	edges [][]int
}

func (g *InMemoryCubeGridGraph) N() int {
	return g.side * g.side * g.side
}

func (g *InMemoryCubeGridGraph) EncodeGraphNode(i int, j int, k int) int {
	return (i*g.side+j)*g.side + k
}

func (g *InMemoryCubeGridGraph) Edge(n int) []int {
	return g.edges[n]
}
