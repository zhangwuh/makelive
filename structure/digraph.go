package structure

type Digraph interface {
	Graph
	Reverse() Digraph
}

type DiAdjGraph struct {
	vertex  int
	edge    int
	adjList []*adjNode
}

func NewDiAdjGraph(v int, edges []edge) *DiAdjGraph {
	g := &DiAdjGraph{
		vertex:  v,
		adjList: newAdjNodes(v),
	}
	for _, e := range edges {
		g.AddEdge(e.v, e.w)
	}
	return g
}

func (g *DiAdjGraph) V() int {
	return g.vertex
}

func (g *DiAdjGraph) E() int {
	return g.edge
}

func (g *DiAdjGraph) Adj(v interface{}) []interface{} {
	return g.adjList[v.(int)].edges
}

func (g *DiAdjGraph) AddEdge(v, w interface{}) {
	if v.(int) >= g.V() || w.(int) >= g.V() {
		return
	}
	defer func() {
		g.edge++
	}()
	g.adjList[v.(int)].edges = append(g.adjList[v.(int)].edges, w.(int))
}

func (g *DiAdjGraph) Reverse() Digraph {
	var edges []edge
	for i := 0; i < g.V(); i++ {
		for _, j := range g.Adj(i) {
			edges = append(edges, edge{j.(int), i})
		}
	}
	return NewDiAdjGraph(g.V(), edges)
}

type DirectedDfs struct {
	marked []bool
	g      Digraph
}

func NewDirectedDfs(g Digraph, sources []int) *DirectedDfs {
	d := &DirectedDfs{g: g, marked: make([]bool, g.V())}
	for _, s := range sources {
		d.dfs(s)
	}
	return d
}

func (dfs *DirectedDfs) dfs(s int) {
	dfs.marked[s] = true
	for _, n := range dfs.g.Adj(s) {
		if !dfs.isMarked(n.(int)) {
			dfs.marked[n.(int)] = true
			dfs.dfs(n.(int))
		}
	}
}

func (dfs *DirectedDfs) isMarked(i int) bool {
	return dfs.marked[i]
}
