package arithmetics

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
)

func Add(parts []string, state *state.State) error {
	if len(parts) != 4 {
	  	return fmt.Errorf("invalid argument count, expected 3, got %d", len(parts))
	}

	if !state.HasReg(parts[1]) || !state.HasReg(parts[2]) || !state.HasReg(parts[3]) {
	  	return fmt.Errorf("invalid register %s", parts[1])
	}

	state.Registers[parts[3]] = state.Registers[parts[1]] + state.Registers[parts[2]]
	return nil
}
