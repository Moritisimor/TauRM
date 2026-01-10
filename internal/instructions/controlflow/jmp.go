package controlflow

import (
	"fmt"
	"strconv"
	"github.com/Moritisimor/TauRM/internal/state"
)

func Jmp(parts []string, state *state.State) error {
	if len(parts) != 2 {
	   	return fmt.Errorf("invalid argument count, expected 2, got %d", len(parts))
	}

	idx, err := strconv.Atoi(parts[1])
	if err != nil {
		if i, ok := state.Labels[parts[1]]; ok {
	  		idx = i
	  	} else {
	  		return fmt.Errorf("could not parse %s to int", parts[1])
	  	}
	}

	state.Pc = idx
	return nil
}
