package arithmetics

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
)

func Inc(parts []string, state *state.State) error {
	if len(parts) != 2 {
		return fmt.Errorf("invalid argument count, expected 2, got %d", len(parts))
	}

	if !state.HasReg(parts[1]) {
		return fmt.Errorf("invalid register %s", parts[1])
	}

	state.Registers[parts[1]]++
	return nil
}
