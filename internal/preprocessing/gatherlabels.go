package preprocessing

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"strings"
)

func GatherLabels(parts []string, state *state.State) {
	for i, s := range parts {
		if before, ok := strings.CutSuffix(s, ":"); ok {
			state.Labels[before] = i
		}
	}
}
