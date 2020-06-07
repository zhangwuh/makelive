// graph.go  * Created on  2020/6/4
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package structure

import (
	"fmt"

	"mklive.zhangwuh.com/utils"
)

type Graph interface {
	V() int             //vertex amount
	E() int             //edge amount
	AddEdge(v1, v2 int) //add edge between v1 and v2
	Adj(v int) []int    // all vertex connected to v
}

type adjNode struct {
	edges []int
}

// adjacency-lists based graph
type AdjGraph struct {
	vertex  int
	edge    int
	adjList []*adjNode
}

type edge struct {
	v, w int
}

func NewAdjGraph(v int, edges []edge) *AdjGraph {
	g := &AdjGraph{
		vertex:  v,
		adjList: newAdjNodes(v),
	}
	for _, e := range edges {
		g.AddEdge(e.v, e.w)
	}
	return g
}

func newAdjNodes(v int) []*adjNode {
	var nodes []*adjNode
	for i := 0; i < v; i++ {
		nodes = append(nodes, &adjNode{})
	}
	return nodes
}

func (g *AdjGraph) V() int {
	return g.vertex
}

func (g *AdjGraph) E() int {
	return g.edge
}

func (g *AdjGraph) Adj(v int) []int {
	return g.adjList[v].edges
}

func (g *AdjGraph) AddEdge(v, w int) {
	if v >= g.V() || w >= g.vertex {
		return
	}
	defer func() {
		g.edge++
	}()
	g.adjList[v].edges = append(g.adjList[v].edges, w)
	g.adjList[w].edges = append(g.adjList[w].edges, v)
}

type GraphResolver struct {
	g Graph
}

func NewGraphResolver(g Graph) *GraphResolver {
	if g == nil {
		return nil
	}
	return &GraphResolver{g}
}

func (u *GraphResolver) Degree(v int) int {
	return len(u.g.Adj(v))
}

func (u *GraphResolver) MaxDegree() int {
	var max int
	for i := 0; i < u.g.V(); i++ {
		if d := u.Degree(i); d > max {
			max = d
		}
	}
	return max
}

func (u *GraphResolver) AvgDegree(g Graph) float64 {
	v := g.V()
	if v == 0 {
		return 0
	}
	return 2 * float64(g.E()) / float64(g.V())
}

func (u *GraphResolver) NumberOfSelfLoops(g Graph) int {
	var amount int
	for i := 0; i < g.V(); i++ {
		for _, j := range g.Adj(i) {
			if i == j {
				amount++
			}
		}
	}
	return amount / 2
}

func (u *GraphResolver) Desc(g Graph) string {
	desc := fmt.Sprintf("vertex:%d edge:%d\n", g.V(), g.E())
	for i := 0; i < g.V(); i++ {
		desc += fmt.Sprintf("v connected to:%v\n", g.Adj(i))
	}
	return desc
}

type GraphSearcher interface {
	Marked(v, w int) bool //is v connected to w
	Count(v int) int      //amount of nodes connected to v
}

//Weighted quick-union with path compression.
type ufSearcher struct {
	g  Graph
	wf [][]int // count of each tree
	uf []int
}

func newUfSearcher(g Graph) *ufSearcher {
	us := &ufSearcher{g: g, uf: make([]int, g.V()), wf: make([][]int, g.V())}
	for i := 0; i < g.V(); i++ {
		us.uf[i] = i
		us.wf[i] = []int{i}
	}
	for v := 0; v < g.V(); v++ {
		for _, j := range g.Adj(v) {
			us.union(v, j)
		}
	}
	return us
}

func (u *ufSearcher) union(v int, j int) {
	p := u.root(v)
	q := u.root(j)

	if p == q {
		return
	}
	if len(u.wf[p]) > len(u.wf[q]) || (len(u.wf[p]) == len(u.wf[q]) && p < q) {
		u.setRoot(j, p)
	} else {
		u.setRoot(v, q)
	}

}

func (u *ufSearcher) setRoot(v int, root int) {
	if v == root {
		return
	}
	defer func() {
		u.uf[v] = root
		u.wf[v] = []int{}
		if !utils.ContainsInt(u.wf[root], v) {
			u.wf[root] = append(u.wf[root], v)
		}
	}()
	parent := u.uf[v]
	if v == parent {
		return
	}
	u.setRoot(parent, root)
}

func (u *ufSearcher) root(w int) (root int) {
	return u.uf[w]
}

func (u *ufSearcher) Marked(v, w int) bool {
	return u.root(v) == u.root(w)
}

func (u *ufSearcher) Count(v int) int {
	return len(u.wf[u.root(v)]) - 1 //exclude self
}

type dfsSearcher struct {
	g        Graph
	marked   [][]int
	edgeTo   [][]int
	hasCycle bool
	isBlack  []bool
	twoColor bool
}

func newDfsSearcher(g Graph) *dfsSearcher {
	searcher := &dfsSearcher{g: g, isBlack: make([]bool, g.V()), twoColor: true}
	for i := 0; i < g.V(); i++ {
		var l []int
		var e []int
		for i := 0; i < g.V(); i++ {
			e = append(e, -1)
		}
		searcher.dfs(&l, &e, i, i)
		searcher.marked = append(searcher.marked, l)
		searcher.edgeTo = append(searcher.edgeTo, e)
	}
	return searcher
}

func (u *dfsSearcher) dfs(marked *[]int, edgeTo *[]int, w int, parent int) {
	*marked = append(*marked, w)
	for _, n := range u.g.Adj(w) {
		if !utils.ContainsInt(*marked, n) {
			(*edgeTo)[n] = w
			u.isBlack[n] = !u.isBlack[w]
			u.dfs(marked, edgeTo, n, w)
		} else {
			if n != parent {
				u.hasCycle = true
			}
			if !u.isBlack[n] == u.isBlack[w] {
				u.twoColor = false
			}
		}
	}
}

func (u *dfsSearcher) isBipartite() bool {
	return u.twoColor
}

func (u *dfsSearcher) pathTo(s, v int) []int {
	edges := u.edgeTo[s]
	stack := &LinkedStack{}
	for v != s {
		stack.Push(v)
		path := edges[v]
		if path == -1 {
			return []int{}
		}
		v = path
	}
	stack.Push(s)
	var path []int
	for !stack.IsEmpty() {
		path = append(path, stack.Pop().(int))
	}
	return path
}

func (u *dfsSearcher) Marked(v, w int) bool {
	return utils.ContainsInt(u.marked[v], w)
}

func (u *dfsSearcher) Count(v int) int {
	return len(u.marked[v]) - 1
}

type bfsSearcher struct {
	g      Graph
	marked [][]int
	edgeTo [][]int
}

func newBfsSearcher(g Graph) *bfsSearcher {
	searcher := &bfsSearcher{g: g}
	for i := 0; i < g.V(); i++ {
		var l []int
		var e []int
		for i := 0; i < g.V(); i++ {
			e = append(e, -1)
		}
		searcher.bfs(&l, &e, i)
		searcher.marked = append(searcher.marked, l)
		searcher.edgeTo = append(searcher.edgeTo, e)
	}
	return searcher
}

func (u *bfsSearcher) bfs(marked *[]int, edgeTo *[]int, v int) {
	q := NewQueue()
	q.Enqueue(v)
	*marked = append(*marked, v)
	for {
		val := q.Dequeue()
		if val == nil {
			break
		}
		w := val.(int)
		for _, n := range u.g.Adj(w) {
			if !utils.ContainsInt(*marked, n) {
				(*edgeTo)[n] = w
				*marked = append(*marked, n)
				q.Enqueue(n)
			}
		}

	}
}

func (u *bfsSearcher) pathTo(s, v int) []int {
	edges := u.edgeTo[s]
	stack := &LinkedStack{}
	for v != s {
		stack.Push(v)
		path := edges[v]
		if path == -1 {
			return []int{}
		}
		v = path
	}
	stack.Push(s)
	var path []int
	for !stack.IsEmpty() {
		path = append(path, stack.Pop().(int))
	}
	return path
}

func (u *bfsSearcher) Marked(v, w int) bool {
	return utils.ContainsInt(u.marked[v], w)
}

func (u *bfsSearcher) Count(v int) int {
	return len(u.marked[v]) - 1
}
