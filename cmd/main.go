package main

import (
	"github.com/hallelujah-shih/cover_order/config"
	"github.com/hallelujah-shih/cover_order/gen_task"
)

func main() {
	fname := "config.yaml"
	cfg, err := config.New(fname)
	if err != nil {
		panic(err)
	}
	tsk := gen_task.New(cfg.Files)
	tsk.Run()
}
