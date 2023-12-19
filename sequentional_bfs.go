package main

type SequentialBFS struct{}

func (SequentialBFS) BFS(graph Graph) []int {
	d := make([]int, graph.N())
	for i := range d {
		d[i] = -1
	}

	d[0] = 0
	q := make([]int, 0)
	q = append(q, 0)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, u := range graph.Edge(v) {
			if d[u] == -1 {
				d[u] = d[v] + 1
				q = append(q, u)
			}
		}
	}

	return d
}
