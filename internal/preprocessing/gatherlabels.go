package preprocessing

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/TauRM/internal/state"
)

func GatherLabels(parts []string, state *state.State) error {
	for i, s := range parts {
		if before, ok := strings.CutSuffix(s, ":"); ok {
			if _, ok := state.Labels[before]; ok {
				return fmt.Errorf("duplicate label assignment %s", before)
			}

			state.Labels[before] = i + 1
		}
	}

	return nil
}
