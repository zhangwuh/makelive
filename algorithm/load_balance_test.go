package algorithm

import "testing"

func TestBalancer_Take(t *testing.T) {
	balancer := NewBalancer(3)
	tasks := []*task{
		{"a", 10},
		{"b", 5},
		{"c", 20},
		{"d", 15},
		{"e", 11},
	}
	balancer.Take(tasks)
}
