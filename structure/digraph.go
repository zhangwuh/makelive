package structure

import (
	"fmt"
	"math"

	"mklive.zhangwuh.com/utils"
)

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
			dfs.dfs(n.(int))
		}
	}
}

func (dfs *DirectedDfs) isMarked(i int) bool {
	return dfs.marked[i]
}

type DirectedCycleDetector struct {
	g       Digraph
	stack   Stack
	edgeTo  []int
	marked  []bool
	onstack []bool
}

func NewDirectedCycleDetector(g Digraph) *DirectedCycleDetector {
	dc := &DirectedCycleDetector{
		g:       g,
		stack:   NewLinkedStack(nil),
		marked:  make([]bool, g.V()),
		onstack: make([]bool, g.V()),
		edgeTo:  make([]int, g.V()),
	}
	for i := 0; i < g.V(); i++ {
		dc.dfs(i)
	}
	return dc
}

func (dc *DirectedCycleDetector) HasCycle() bool {
	return !dc.stack.IsEmpty()
}
func (dc *DirectedCycleDetector) Cycle() []int {
	var c []int
	for !dc.stack.IsEmpty() {
		c = append(c, dc.stack.Pop().(int))
	}
	return c
}

func (dfs *DirectedCycleDetector) dfs(s int) {
	dfs.marked[s] = true
	dfs.onstack[s] = true
	for _, n := range dfs.g.Adj(s) {
		current := n.(int)
		dfs.edgeTo[current] = s
		if dfs.HasCycle() {
			break
		}
		if !dfs.isMarked(current) {
			dfs.dfs(current)
		} else {
			if dfs.onstack[current] {
				for x := s; x != current; x = dfs.edgeTo[x] {
					dfs.stack.Push(x)
				}
				dfs.stack.Push(current)
				dfs.stack.Push(s)
			}
		}

	}
	dfs.onstack[s] = false
}

func (dfs *DirectedCycleDetector) isMarked(i int) bool {
	return dfs.marked[i]
}

type DepthFirstOrder struct {
	g           Digraph
	pre         []int
	post        []int
	reversePost Stack
	marked      []bool
}

func NewDepthFirstOrder(g Digraph) *DepthFirstOrder {
	do := &DepthFirstOrder{
		g:           g,
		marked:      make([]bool, g.V()),
		reversePost: NewLinkedStack(nil),
	}

	for v := 0; v < g.V(); v++ {
		do.dfs(v)
	}

	return do
}

func (dfo *DepthFirstOrder) dfs(v int) {
	if dfo.marked[v] {
		return
	}
	dfo.marked[v] = true
	dfo.pre = append(dfo.pre, v)
	for _, w := range dfo.g.Adj(v) {
		current := w.(int)
		if !dfo.marked[current] {
			dfo.dfs(current)
		}
	}
	dfo.post = append(dfo.post, v)
	dfo.reversePost.Push(v)
}

func (dfo *DepthFirstOrder) reversePostQueue() []int {
	var q []int
	for !dfo.reversePost.IsEmpty() {
		q = append(q, dfo.reversePost.Pop().(int))
	}
	return q
}

type Topological interface {
	isDAG() bool
	order() []int
}

type topologicalSorter struct {
	g     Digraph
	stack Stack
}

func NewTopologicalSorter(g Digraph) *topologicalSorter {
	return &topologicalSorter{
		g:     g,
		stack: NewLinkedStack(nil),
	}
}

func (ts *topologicalSorter) isDAG() bool {
	dd := NewDirectedCycleDetector(ts.g)
	return !dd.HasCycle()
}

func (ts *topologicalSorter) order() []int {
	if !ts.isDAG() {
		return nil
	}
	return NewDepthFirstOrder(ts.g).reversePostQueue()
}

type SCC interface { //Strongly Connected Components
	Connected(v, w int) bool
	Count() int   //count of connected components
	Id(v int) int // id of the component v belongs to
}

type normalSCC struct {
	g          Digraph
	components [][]int
	marked     []bool
}

func NewNormalSCC(g Digraph) SCC {
	scc := &normalSCC{
		g:      g,
		marked: make([]bool, g.V()),
	}
	dfo := NewDepthFirstOrder(g.Reverse())
	for !dfo.reversePost.IsEmpty() {
		var coms []int
		v := dfo.reversePost.Pop().(int)
		scc.dfs(v, &coms)
		if len(coms) > 0 {
			scc.components = append(scc.components, coms)
		}
	}

	return scc
}

func (scc *normalSCC) Connected(v, w int) bool {
	return scc.Id(v) == scc.Id(w)
}

func (scc *normalSCC) Count() int {
	return len(scc.components)
}

func (scc *normalSCC) Id(v int) int {
	for i, c := range scc.components {
		if utils.ContainsInt(c, v) {
			return i
		}
	}
	return v
}

func (scc *normalSCC) dfs(v int, coms *[]int) {
	scc.marked[v] = true
	*coms = append(*coms, v)
	for _, w := range scc.g.Adj(v) {
		current := w.(int)
		if !scc.marked[current] {
			scc.dfs(current, coms)
		}
	}
}

//edge weighted digraph start
type directedEdge struct {
	edge
	weight float32
}

func NewDirectedEdg(from, to int, weight float32) *directedEdge {
	return &directedEdge{edge{from, to}, weight}
}

func (de directedEdge) from() int {
	return de.edge.v
}

func (de directedEdge) to() int {
	return de.edge.w
}

func (de directedEdge) desc() string {
	return fmt.Sprintf("directed edge from %d to %d weight:%f", de.from(), de.to(), de.weight)
}

type EdgeWeightedDigraph interface {
	V() int                     //vertex amount
	E() int                     //edge amount
	AddEdge(edge *directedEdge) //add edge between v1 and v2
	Adj(v int) []*directedEdge
	Edges() []*directedEdge
	Desc() string
}

type edgeWeightedDigraph struct {
	v   int
	e   int
	adj [][]*directedEdge
}

func NewEdgeWeightedDigraph(v int) *edgeWeightedDigraph {
	g := &edgeWeightedDigraph{
		v: v,
	}

	g.adj = make([][]*directedEdge, g.V())

	return g
}

func (ed *edgeWeightedDigraph) V() int {
	return ed.v
}

func (ed *edgeWeightedDigraph) E() int {
	return ed.e
}

func (ed *edgeWeightedDigraph) AddEdge(e *directedEdge) {
	ed.adj[e.from()] = append(ed.adj[e.from()], e)
	ed.e++
}

func (ed *edgeWeightedDigraph) Adj(v int) []*directedEdge {
	return ed.adj[v]
}

func (ed *edgeWeightedDigraph) Edges() []*directedEdge {
	var es []*directedEdge
	for _, e := range ed.adj {
		for _, ee := range e {
			es = append(es, ee)
		}
	}
	return es
}

func (ed *edgeWeightedDigraph) Desc() string {
	return fmt.Sprintf("edgeWeightedDigraph has %d vertex and %d edges", ed.V(), ed.E())
}

//shortest path calculator from vertex s
type ShortestPath interface {
	S() int                   //start vertex
	DistanceTo(v int) float32 // return maxfloat32 if no path detected
	HasPathTo(v int) bool
	PathTo(v int) []*directedEdge
	Graph() EdgeWeightedDigraph
}

type shortestPath struct {
	s       int
	edgesTo []*directedEdge //detected edges in the shortest path to s
	distTo  []float32       // detected distance to s
	g       EdgeWeightedDigraph
}

func NewShortestPath(s int, graph EdgeWeightedDigraph) ShortestPath {
	sp := &shortestPath{s: s, g: graph}
	for i := 0; i < graph.V(); i++ {
		if i == s {
			sp.distTo = append(sp.distTo, 0)
		} else {
			sp.distTo = append(sp.distTo, math.MaxFloat32)
		}
	}
	sp.edgesTo = make([]*directedEdge, graph.V())
	sp.searchPath(s)
	return sp
}

func (sp *shortestPath) searchPath(s int) {
	sp.relaxVertex(s)
	for _, v := range sp.Graph().Adj(s) {
		sp.searchPath(v.w)
	}
}

func (sp *shortestPath) Graph() EdgeWeightedDigraph {
	return sp.g
}

func (sp *shortestPath) S() int {
	return sp.s
}

func (sp *shortestPath) DistanceTo(v int) float32 {
	return sp.distTo[v]
}

func (sp *shortestPath) HasPathTo(v int) bool {
	return sp.distTo[v] != math.MaxFloat32
}

func (sp *shortestPath) PathTo(v int) []*directedEdge {
	stack := NewLinkedStack(nil)
	for {
		if v == sp.s {
			break
		}
		if sp.HasPathTo(v) {
			edge := sp.edgesTo[v]
			stack.Push(edge)
			v = edge.from()
		}
	}
	var res []*directedEdge
	for !stack.IsEmpty() {
		res = append(res, stack.Pop().(*directedEdge))
	}
	return res
}

func printPath(ds []*directedEdge) {
	for _, d := range ds {
		fmt.Println(fmt.Sprintf("%d -> %d", d.v, d.w))
	}
}

func (sp *shortestPath) relax(e *directedEdge) {
	if sp.DistanceTo(e.from())+e.weight < sp.DistanceTo(e.to()) { //do relax
		fmt.Println(fmt.Sprintf("relax %d to %d", e.v, e.w))
		sp.distTo[e.to()] = sp.DistanceTo(e.from()) + e.weight
		sp.edgesTo[e.to()] = e
	}
}

func (sp *shortestPath) relaxVertex(v int) {
	fmt.Println(fmt.Sprintf("relax vertex %d", v))
	for _, e := range sp.g.Adj(v) {
		sp.relax(e)
	}
}

//edge weighted digraph end
