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

import "fmt"

type Graph interface {
	V() int             //vertex amount
	E() int             //edge amount
	AddEdge(v1, v2 int) //add edge between v1 and v2
	Adj(v int) []int    // all vertex connected to v
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

func (u *GraphResolver) MaxDegree(g Graph) int {
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
