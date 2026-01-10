package main

import (
	"github.com/Moritisimor/TauRM/internal/interpreter"
	"github.com/Moritisimor/TauRM/internal/preprocessing"
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lambdavm <program>")
		return
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error while reading file: %s", err.Error())
	}

	program := strings.Split(string(bytes), "\n")
	for i, cmd := range program {
		if !strings.HasPrefix(strings.TrimSpace(cmd), "#") {
			program[i] = strings.TrimSpace(strings.Split(cmd, "#")[0])
		}
	} 

	state := state.Initialize()
	preprocessing.GatherLabels(program, state)
	interpreter.InterpreterLoop(program, state)
}
