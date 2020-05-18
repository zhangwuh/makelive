// greedy_test.go  * Created on  2020/5/11
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package practise

import (
	"fmt"
	"testing"
)

func TestActivitySelect(t *testing.T) {
	as := &ActivitySelection{
		candidates: []*Activity{
			{1, 4}, {3, 5}, {0, 6}, {5, 7}, {3, 8}, {5, 9}, {6, 10}, {8, 11}, {8, 12}, {2, 13}, {12, 14},
		},
	}

	for _, s := range as.selectActivity() {
		fmt.Println(fmt.Sprintf("%d-%d", s.start, s.end))
	}
}

func TestTaskSelector(t *testing.T) {
	ts := newTaskSelector([]*Task{
		{9, 15}, {2, 2}, {5, 18}, {7, 1}, {4, 25}, {2, 20}, {5, 8}, {7, 10}, {4, 12}, {3, 5},
	})
	var total int
	for _, t := range ts.selectTasks() {
		fmt.Println(fmt.Sprintf("%d-%d", t.deadline, t.profit))
		total += t.profit
	}
	fmt.Println(total)
}

func TestPackageSelection(t *testing.T) {
	/*[3, 4, 5] //物品重量列表
	v = [4, 5, 6] //物品价值列表
	C = 10*/
	ps := &packageSelection{
		candidates: []*stone{{3, 4}, {4, 5}, {5, 6}},
		maxWeight:  10,
	}
	ps.pack()
	stones := ps.selected
	for _, s := range stones {
		fmt.Println(fmt.Sprintf("%f-%f", s.weight, s.value))
	}
	fmt.Println(ps.weight)
	fmt.Println(ps.value)

}
