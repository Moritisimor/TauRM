package memory

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
)

func Copy(parts []string, state *state.State) error {
	if len(parts) != 3 {
		return fmt.Errorf("invalid argument count, expected 3, got %d", len(parts))
	}

	if !state.HasReg(parts[1]) || !state.HasReg(parts[2]) {
		return fmt.Errorf("invalid register(s) %s or %s", parts[1], parts[2])
	}

	state.Registers[parts[2]] = state.Registers[parts[1]]
	return nil
}
