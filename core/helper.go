package core

import (
	"os"
	"os/exec"
	"sync"
	"time"
)

var rw sync.Mutex
var wg sync.WaitGroup

const (
	fct = time.Millisecond * 33
)

func cmd(out bool, comand string, args ...string) ([]byte, error) {
	v := exec.Command(comand, args...)
	v.Stdin = os.Stdin
	if out {
		v.Stdout = os.Stdout
		v.Run()
	}
	return v.Output()
}

func clear() {
	cmd(true, "clear")
	cmd(true, "tput", "civis")
	cmd(true, "stty", "-echo", "iexten", "-icanon", "min", "1")
}
func reset() {
	cmd(true, "clear")
	cmd(true, "tput", "reset")
}
