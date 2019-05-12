/*
reference: https://biexiang.github.io/post/permutation/
*/
package gen_task

import (
	"bytes"
	"fmt"
)

func (g *Generator) gen() chan *Task {
	rt := make(chan *Task, 10000)
	go func() {
		defer close(rt)
		rangeAll(g.allLines, 0, rt, nil)
	}()
	return rt
}

func rangeAll(slice []*SubTask, start int, tskChan chan *Task, sliceInfo map[int]int) {
	if start == 0 {
		sliceInfo = make(map[int]int)
	}

	size := len(slice)
	if start == size-1 && checkValid(slice) {
		var tsk Task
		var buf bytes.Buffer
		for _, cmd := range slice {
			buf.WriteString(fmt.Sprintf("%d: %d ", cmd.FileID, cmd.LineNum))
			tsk.Subs = append(tsk.Subs, cmd)
		}
		tsk.ID = buf.String()
		tskChan <- &tsk
	}

	for i := start; i < size; i++ {
		if i == start || slice[i] != slice[start] {
			i_k := slice[i].FileID
			i_v := slice[i].LineNum
			s_k := slice[start].FileID
			s_v := slice[start].LineNum

			if i_k == s_k && i_v > s_v {
				continue
			}

			oldValue := sliceInfo[i_k]
			if oldValue > i_v {
				continue
			}
			sliceInfo[i_k] = i_v

			slice[i], slice[start] = slice[start], slice[i]
			rangeAll(slice, start+1, tskChan, sliceInfo)
			slice[i], slice[start] = slice[start], slice[i]

			sliceInfo[i_k] = oldValue
		}
	}
}

func checkValid(s []*SubTask) bool {
	tmp := map[int]int{}
	for _, i := range s {
		key := i.FileID
		value := i.LineNum
		if tmp[key] > value {
			return false
		} else {
			tmp[key] = value
		}
	}
	return true
}
