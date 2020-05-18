// greedy.go  * Created on  2020/5/11
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package practise

import "sort"

type Activity struct {
	start int
	end   int
}

type ActivitySelection struct {
	candidates []*Activity
	selected   []*Activity
	current    *Activity
}

func (as *ActivitySelection) selectActivity() []*Activity {
	as.sortByEnd()
	as.doSelect()
	return as.selected
}

func (as *ActivitySelection) sortByEnd() {
	sort.Slice(as.candidates, func(i, j int) bool {
		return as.candidates[i].end < as.candidates[j].end
	})
}

func (as *ActivitySelection) doSelect() {
	for _, a := range as.candidates {
		if len(as.selected) == 0 {
			as.appendToSelected(a)
			continue
		}
		if a.start >= as.current.end {
			as.appendToSelected(a)
		}
	}
}

func (as *ActivitySelection) appendToSelected(a *Activity) {
	as.selected = append(as.selected, a)
	as.current = a
}

type Task struct {
	deadline int
	profit   int
}

type TaskSelector struct {
	candidates []*Task
	selected   []*Task
	slots      []int
}

func newTaskSelector(candidates []*Task) *TaskSelector {
	var maxDeadline int
	for _, c := range candidates {
		if maxDeadline < c.deadline {
			maxDeadline = c.deadline
		}
	}
	var slots []int
	for ; maxDeadline > 0; maxDeadline-- {
		slots = append(slots, 0)
	}
	return &TaskSelector{
		candidates: candidates,
		slots:      slots,
	}
}

func (ts *TaskSelector) sort() {
	sort.Slice(ts.candidates, func(i, j int) bool {
		return ts.candidates[i].profit > ts.candidates[j].profit
	})
}

func (ts *TaskSelector) pickSlot(task *Task) bool {
	for i := task.deadline - 1; i >= 0; i-- {
		if ts.slots[i] == 0 {
			ts.slots[i] = 1
			return true
		}
	}
	return false
}

func (ts *TaskSelector) selectTasks() []*Task {
	ts.sort()
	for _, t := range ts.candidates {
		if ts.pickSlot(t) {
			ts.selected = append(ts.selected, t)
		}
	}
	return ts.selected
}

type stone struct {
	weight float32
	value  float32
}

type packageSelection struct {
	candidates []*stone
	selected   []*stone
	maxWeight  float32
	weight     float32
	value      float32
}

func (ps *packageSelection) put(s *stone) {
	if ps.weight+s.weight > ps.maxWeight {
		ps.put(s.split(ps.maxWeight - ps.weight))
		return
	}
	ps.selected = append(ps.selected, s)
	ps.weight += s.weight
	ps.value += s.value
}

func (ps *packageSelection) sort() {
	sort.Slice(ps.candidates, func(i, j int) bool {
		return ps.candidates[i].value/ps.candidates[i].weight < ps.candidates[j].value/ps.candidates[j].weight
	})
}

func (s *stone) split(availableWeight float32) *stone {
	return &stone{weight: availableWeight, value: s.value * (availableWeight / s.weight)}
}
func (ps *packageSelection) pack() {
	ps.sort()
	for _, s := range ps.candidates {
		ps.put(s)
		if ps.maxWeight == ps.weight {
			return
		}
	}
}
