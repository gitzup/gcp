package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var stderr = make(chan int)
	var stdout = make(chan int)

	go func() {
		for {
			os.Stdout.WriteString("Stdout...\n")
			delay := rand.Intn(3)
			time.Sleep(time.Duration(delay) * time.Second)
			stdout <- delay
		}
	}()
	go func() {
		for {
			os.Stderr.WriteString("Stderr...\n")
			delay := rand.Intn(3)
			time.Sleep(time.Duration(delay) * time.Second)
			stderr <- delay
		}
	}()

	stderrWait := 0
	stdoutWait := 0
	done := false
	for !done {
		select {
		case delay := <-stderr:
			stderrWait += delay
		case delay := <-stdout:
			stdoutWait += delay
		}
		if stderrWait >= 1 || stdoutWait >= 1 {
			done = true
		}
	}

	err := os.MkdirAll("/gitzup", os.ModeDir)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/gitzup/result.json", []byte(`{"test":1}`), 0644)
	if err != nil {
		panic(err)
	}

	return
}
