package structure

import "strings"

type SymbolGraph interface {
	HasKey(key string) bool
	Index(key string) (int, bool)
	Key(index int) (string, bool)
	Graph() Graph
}

type symbolGraph struct {
	g       Graph
	indexes map[string]int
	keys    []string
	counter int
}

func NewSymbolGraph(es []string) *symbolGraph {
	sg := &symbolGraph{
		indexes: make(map[string]int),
	}

	var edges []edge
	for _, e := range es {
		keys := strings.Split(e, " ")
		key1 := keys[0]
		key2 := keys[1]

		index1, ok := sg.Index(key1)
		if !ok {
			index1 = sg.newIndex()
			sg.indexes[key1] = index1
		}
		index2, ok := sg.Index(key2)
		if !ok {
			index2 = sg.newIndex()
			sg.indexes[key2] = index2
		}

		edges = append(edges, edge{index1, index2})
	}

	sg.keys = make([]string, sg.counter)
	for k, v := range sg.indexes {
		sg.keys[v] = k
	}

	sg.g = NewAdjGraph(sg.counter, edges)
	return sg
}

func (sg *symbolGraph) V() int {
	return sg.g.V()
}

func (sg *symbolGraph) E() int {
	return sg.g.E()
}

func (sg *symbolGraph) AddEdge(v1, v2 interface{}) {
	key1 := v1.(string)
	key2 := v2.(string)
	index1, ok := sg.Index(key1)
	if !ok {
		index1 = sg.newIndex()
	}
	index2, ok := sg.Index(key2)
	if !ok {
		index2 = sg.newIndex()
	}

	sg.g.AddEdge(index1, index2)
}

func (sg *symbolGraph) newIndex() int {
	defer func() {
		sg.counter++
	}()
	return sg.counter
}

func (sg *symbolGraph) Adj(v interface{}) []interface{} {
	key := v.(string)
	index, ok := sg.Index(key)
	if !ok {
		return nil
	}
	adjs := sg.g.Adj(index)
	var res []interface{}
	for _, j := range adjs {
		k, _ := sg.Key(j.(int))
		res = append(res, k)
	}
	return res
}

func (sg *symbolGraph) HasKey(key string) bool {
	_, ok := sg.indexes[key]
	return ok
}

func (sg *symbolGraph) Index(key string) (int, bool) {
	if i, ok := sg.indexes[key]; ok {
		return i, true
	}
	return 0, false
}

func (sg *symbolGraph) Key(index int) (string, bool) {
	if index >= 0 && index < len(sg.keys) {
		return sg.keys[index], true
	}
	return "", false
}

func (sg *symbolGraph) Graph() Graph {
	return sg.g
}
