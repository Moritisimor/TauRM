package memory

import "github.com/Moritisimor/TauRM/internal/state"

import (
	"strconv"
	"fmt"
)

func Load(parts []string, state *state.State) error {
	if len(parts) != 3 {
		return fmt.Errorf("invalid argument count, expected 3, got %d", len(parts))
	}

	if !state.HasReg(parts[1]) {
		return fmt.Errorf("invalid register %s", parts[1])
	}

	i, err := strconv.Atoi(parts[2]); if err != nil {
		return fmt.Errorf("could not parse %s to int", parts[2])
	}

	state.Registers[parts[1]] = i
	return nil
}
