package algorithm

import (
	"fmt"
	"strconv"

	"mklive.zhangwuh.com/sort"
)

type task struct {
	name        string
	timeConsume int
}

func (t *task) Compare(to sort.Comparable) int {
	if t.timeConsume == to.(*task).timeConsume {
		return 0
	}
	if t.timeConsume > to.(*task).timeConsume {
		return 1
	}
	return -1
}

type balancer struct {
	workers []*worker
	size    int
}

func NewBalancer(size int) *balancer {
	b := &balancer{size: size}
	size++
	for size > 0 {
		b.workers = append(b.workers, &worker{name: strconv.Itoa(size)})
		size--
	}
	return b
}

func (balancer *balancer) insert(worker *worker) {
	balancer.workers = append(balancer.workers, worker)
	balancer.size++
	balancer.swim()
}

type worker struct {
	workload int
	name     string
}

func (balancer *balancer) Take(tasks []*task) {
	qs := &sort.NewQuickSort{
		Desc: true,
	}
	var input []sort.Comparable
	for _, v := range tasks {
		input = append(input, v)
	}
	qs.Sort(input)
	for _, t := range input {
		tt := t.(*task)
		worker := balancer.pop()
		worker.workload += tt.timeConsume
		balancer.insert(worker)
		fmt.Println(fmt.Sprintf("task %s %d assigned to worker with load %s %d", tt.name, tt.timeConsume, worker.name, worker.workload))
	}
}

func (balancer *balancer) swap(i, j int) {
	balancer.workers[i], balancer.workers[j] = balancer.workers[j], balancer.workers[i]
}

func (balancer *balancer) pop() *worker {
	if balancer.size == 0 {
		return nil
	}
	root := balancer.workers[1]
	defer func() {
		balancer.swap(1, balancer.size)
		balancer.sink()
		balancer.size--
	}()
	return root
}

func (balancer *balancer) swim() { // swim last node
	i := balancer.size
	for i/2 > 0 {
		parent := i / 2
		if balancer.workers[i].workload < balancer.workers[parent].workload {
			balancer.swap(i, parent)
			i = parent
		} else {
			return
		}
	}
}

func (balancer *balancer) min(i, j int) int {
	if balancer.workers[i].workload < balancer.workers[j].workload {
		return i
	}
	return j
}

func (balancer *balancer) sink() {
	i := 1
	for 2*i < balancer.size {
		lc := 2 * i
		var min int
		if lc+1 < balancer.size {
			lr := lc + 1
			min = balancer.min(lc, lr)
		} else {
			min = lc
		}
		if balancer.workers[min].workload < balancer.workers[i].workload {
			balancer.swap(i, min)
			i = min
		} else {
			return
		}
	}
}
