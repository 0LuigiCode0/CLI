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

func equals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findKey(key []byte) string {
	for _, k := range KeyList {
		if equals(k.data, key) {
			return k.Name
		}
	}

	return ""
}
