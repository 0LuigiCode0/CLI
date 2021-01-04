package core

import (
	"context"
	"os"
	"time"

	"github.com/000mrLuigi000/Library/logger"
)

type window struct {
	lines  int
	column int
	frame  [][]string
	log    *logger.Logger
	fct    time.Duration

	apptowin
}

type apptowin interface {
	size() (int, int)
	clear()
}

func (w *window) reSize(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			w.log.Info("Resize stoped")
			return
		default:
			newLine, newColumn := w.size()
			if newLine != w.lines && newColumn != w.column {
				w.clear()
				frame := make([][]string, newLine, newLine)
				for i := range frame {
					frame[i] = make([]string, newColumn, newColumn)
				}
				w.lines = newLine
				w.column = newColumn
				w.setLayout(frame)
			} else if newLine != w.lines {
				w.clear()
				frame := make([][]string, newLine, newLine)
				for i := range frame {
					frame[i] = make([]string, w.column, w.column)
				}
				w.lines = newLine
				w.setLayout(frame)
			} else if newColumn != w.column {
				w.clear()
				frame := w.frame
				for i := range frame {
					frame[i] = make([]string, newColumn, newColumn)
				}
				w.column = newColumn
				w.setLayout(frame)
			}

			time.Sleep(w.fct)
		}
	}
}

func (w *window) reView(ctx context.Context) {
	defer wg.Done()
	time.Sleep(w.fct / 2)
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

			time.Sleep(w.fct)
		}
	}
}

func (w *window) getPX(i, j int) string {
	rw.Lock()
	defer rw.Unlock()

	return w.frame[i][j]
}

func (w *window) setPX(i, j int, v string) {
	rw.Lock()
	defer rw.Unlock()

	w.frame[i][j] = v
	return
}

func (w *window) setLayout(arr [][]string) {
	rw.Lock()
	defer rw.Unlock()

	w.frame = arr
	return
}
