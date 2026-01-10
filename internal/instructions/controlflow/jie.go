package controlflow

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
	"strconv"
)

func Jie(parts []string, state *state.State) error {
	if len(parts) != 4 {
		return fmt.Errorf("invalid argument count, expected 4, got %d", len(parts))
	}

	if !state.HasReg(parts[1]) || !state.HasReg(parts[2]) {
		return fmt.Errorf("invalid register(s) %s or %s", parts[1], parts[2])
	}

	idx, err := strconv.Atoi(parts[3])
	if err != nil {
		if i, ok := state.Labels[parts[3]]; ok {
			idx = i
		} else {
			return fmt.Errorf("could not parse %s to int", parts[1])
		}
	}

	if state.Registers[parts[1]] == state.Registers[parts[2]] {
		state.Pc = idx
	}

	return nil
}
