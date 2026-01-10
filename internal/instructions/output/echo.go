package output

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
)

func Echo(parts []string, state *state.State) error {
	if len(parts) < 2 {
		return fmt.Errorf("invalid argument count, expected at least 2, got only one")
	}

	for _, s := range parts[1:] {
		if !state.HasReg(s) {
			fmt.Print(s)
		} else {
			fmt.Print(state.Registers[s])
		}

		fmt.Print(" ")
	}

	fmt.Println()
	return nil
}
