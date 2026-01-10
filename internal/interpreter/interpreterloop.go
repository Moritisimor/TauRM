package interpreter

import (
	"fmt"
	"github.com/Moritisimor/TauRM/internal/state"
)

func InterpreterLoop(program []string, state *state.State) {
	for state.Pc < len(program) {
		snap := state.Pc
		if err := interpretCommand(state, program[state.Pc - 1]); err != nil {
			fmt.Printf("Program execution halted at line %d: %s\n", state.Pc, err.Error())
			return
		}

		if state.Pc == snap {
			state.Pc++
		}
	}
}
