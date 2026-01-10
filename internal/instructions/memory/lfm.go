package memory

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
	"strconv"
)

func Lfm(parts []string, state *state.State) error {
	if len(parts) != 3 {
		return fmt.Errorf("Invalid argument count, expected 3, got %d", len(parts))
	}

	if !state.HasReg(parts[1]) {
		return fmt.Errorf("invalid register %s", parts[1])
	}

	address, err := strconv.ParseUint(parts[2], 0, 32); if err != nil {
		return fmt.Errorf("invalid memory address %s", parts[2])
	}

	state.Registers[parts[1]] = state.Memory[uint(address)]
	return nil
}
