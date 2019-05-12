package gen_task

import (
	"bufio"
	"fmt"
	"github.com/hallelujah-shih/cover_order/config"
	"io"
	"os"
	"strings"
	"sync"
)

type WorkerChan struct {
	ID      int
	Consume chan *SubTask
	Result  chan *SubRsp
}

type SubTask struct {
	FileID  int    `json:"file_id"`
	LineNum int    `json:"line_num"`
	Line    string `json:"line"`
}

type SubRsp struct {
	Unexpect bool   `json:"unexpect"`
	Err      string `json:"err"`
}

type Task struct {
	ID   string     `json:"id"`
	Subs []*SubTask `json:"subs"`
}

type Generator struct {
	sync.Mutex
	cfg      *config.ConverFile
	wksChan  map[int]*WorkerChan
	wrks     map[int]*Worker
	allLines []*SubTask
}

func (g *Generator) TotalWrks() int {
	return len(g.cfg.FileList)
}

func (g *Generator) GetChans() []*WorkerChan {
	g.Lock()
	defer g.Unlock()
	var rt []*WorkerChan
	for _, wc := range g.wksChan {
		rt = append(rt, wc)
	}
	return rt
}

func (g *Generator) Run() {
	for _, wrk := range g.wrks {
		go wrk.Run()
	}

	for t := range g.gen() {
		for _, sub := range t.Subs {
			wc := g.wksChan[sub.FileID]
			if wc != nil {
				wc.Consume <- sub
			}
			rt := <-wc.Result
			if rt.Unexpect {
				fmt.Println("task_id:", t.ID, "sub_id:", sub.FileID, ":", sub.LineNum, "err:", rt.Err)
			}
		}
		fmt.Println(t.ID)
	}

	for _, w := range g.wksChan {
		close(w.Consume)
	}
}

func New(cfg *config.ConverFile) *Generator {
	g := &Generator{}
	g.wksChan = make(map[int]*WorkerChan)
	g.wrks = make(map[int]*Worker)

	for fid, fname := range cfg.FileList {
		g.wksChan[fid] = &WorkerChan{
			ID:      fid,
			Consume: make(chan *SubTask),
			Result:  make(chan *SubRsp),
		}

		g.wrks[fid] = newWorker(g.wksChan[fid])

		var line int = 1
		f, err := os.Open(fname)
		if err != nil {
			fmt.Println("open file:", fname, "err:", err)
			continue
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			tmp := &SubTask{
				FileID:  fid,
				LineNum: line,
				Line:    strings.TrimSpace(scanner.Text()),
			}
			line++
			if tmp.Line == "" {
				continue
			}
			g.allLines = append(g.allLines, tmp)
		}
		if err := scanner.Err(); err != nil {
			if err != io.EOF {
				fmt.Println("scanner error:", err)
			}
		}
		f.Close()
	}

	return g
}
