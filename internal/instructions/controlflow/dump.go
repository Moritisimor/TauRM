package controlflow

import (
	"github.com/Moritisimor/TauRM/internal/state"
	"fmt"
	"os"
	"strings"
)

// Not too sure how fitting this is in the controlflow package but it does control the flow so why not
func Dump(state *state.State) error {
	var content strings.Builder

	content.WriteString("[Registers]\n")
	for k, v := range state.Registers {
		fmt.Fprintf(&content, "\t%s: %d\n", k, v)
	}

	content.WriteString("\n[Memory]\n")
	for k, v := range state.Memory {
		fmt.Fprintf(&content, "\t%d: %d\n", k, v)
	}

	fmt.Fprintf(&content, "Last PC Position: %d", state.Pc)

	os.WriteFile("LVMDump", []byte(content.String()), 0755)
	return fmt.Errorf("core dumped")
}