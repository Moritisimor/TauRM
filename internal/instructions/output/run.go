package output

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Moritisimor/TauRM/internal/state"
)

func Run(parts []string, s *state.State) error {
	if len(parts) < 3 {
		return fmt.Errorf("invalid argument count, expected more than 2.")
	}

	if !s.HasReg(parts[1]) {
		return fmt.Errorf("invalid register %s", parts[1])
	}

	cmd := exec.Command(parts[2], parts[3:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run() 
	s.Registers[parts[1]] = cmd.ProcessState.ExitCode()

	return nil
}
