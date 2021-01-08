package core

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/000mrLuigi000/Library/logger"
)

//Window обьект окна
type Window struct {
	lines  int
	column int
	frame  [][]string
	log    *logger.Logger
}

func (w *Window) reSize(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			w.log.Info("Resize stoped")
			return
		default:
			newLine, newColumn := w.size()
			if newLine != w.lines && newColumn != w.column {
				clear()
				frame := make([][]string, newLine, newLine)
				for i := range frame {
					frame[i] = make([]string, newColumn, newColumn)
				}
				w.lines = newLine
				w.column = newColumn
				w.setLayout(frame)
			} else if newLine != w.lines {
				clear()
				frame := make([][]string, newLine, newLine)
				for i := range frame {
					frame[i] = make([]string, w.column, w.column)
				}
				w.lines = newLine
				w.setLayout(frame)
			} else if newColumn != w.column {
				clear()
				frame := w.frame
				for i := range frame {
					frame[i] = make([]string, newColumn, newColumn)
				}
				w.column = newColumn
				w.setLayout(frame)
			}

			time.Sleep(fct)
		}
	}
}

func (w *Window) reView(ctx context.Context) {
	defer wg.Done()
	time.Sleep(fct / 2)
	for {
		select {
		case <-ctx.Done():
			w.log.Info("Review stoped")
			return
		default:
			for i := 0; i < w.lines; i++ {
				for j := 0; j < w.column; j++ {
					w.setPX(i, j, w.frame[i][j])
				}
			}

			resp := "\033[0;0H"
			for i := 0; i < w.lines; i++ {
				for j := 0; j < w.column; j++ {
					px := w.getPX(i, j)
					if px == "" {
						resp += " "
					} else {
						resp += px
					}
				}
				if i < w.lines-1 {
					resp += "\n"
				}
				os.Stdout.WriteString(resp)
				resp = ""
			}

			time.Sleep(fct)
		}
	}
}

func (w *Window) size() (int, int) {
	var newLine, newColumn int
	res, _ := cmd(false, "tput", "lines")
	if len(res) > 0 {
		newLine, _ = strconv.Atoi(string(res[:len(res)-1]))
	}
	res, _ = cmd(false, "tput", "cols")
	if len(res) > 0 {
		newColumn, _ = strconv.Atoi(string(res[:len(res)-1]))
	}
	return newLine, newColumn
}

func (w *Window) getPX(i, j int) string {
	rw.Lock()
	defer rw.Unlock()

	return w.frame[i][j]
}

func (w *Window) setPX(i, j int, v string) {
	rw.Lock()
	defer rw.Unlock()

	w.frame[i][j] = v
	return
}

func (w *Window) setLayout(arr [][]string) {
	rw.Lock()
	defer rw.Unlock()

	w.frame = arr
	return
}
