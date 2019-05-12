package gen_task

import (
	_ "fmt"
)

type Worker struct {
	wrkChan *WorkerChan
}

func (w *Worker) Run() {
	for subTsk := range w.wrkChan.Consume {
		// TODO add sql excute
		// fmt.Println("do task:", subTsk.FileID, ":", subTsk.LineNum, "Op:", subTsk.Line)
		rt := &SubRsp{
			Err: subTsk.Line,
		}
		w.wrkChan.Result <- rt
	}
}

func newWorker(w *WorkerChan) *Worker {
	return &Worker{w}
}
