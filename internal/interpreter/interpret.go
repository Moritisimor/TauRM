package interpreter

import (
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/TauRM/internal/helpers"
	"github.com/Moritisimor/TauRM/internal/instructions"
	"github.com/Moritisimor/TauRM/internal/instructions/arithmetics"
	"github.com/Moritisimor/TauRM/internal/instructions/controlflow"
	"github.com/Moritisimor/TauRM/internal/instructions/memory"
	"github.com/Moritisimor/TauRM/internal/instructions/output"
	"github.com/Moritisimor/TauRM/internal/state"
)

func interpretCommand(state *state.State, command string) error {
	if helpers.IsIgnorable(command) {
		return nil
	}

	parts := strings.Split(command, " ")
	var err error

	switch strings.ToLower(parts[0]) {
	case "load":
		err = memory.Load(parts, state)

	case "copy":
		err = memory.Copy(parts, state)

	case "run":
		err = output.Run(parts, state)

	case "jmp":
		err = controlflow.Jmp(parts, state)

	case "jie":
		err = controlflow.Jie(parts, state)

	case "jig":
		err = controlflow.Jig(parts, state)

	case "jis":
		err = controlflow.Jis(parts, state)

	case "sleep":
		err = instructions.Sleep(parts)

	case "add":
		err = arithmetics.Add(parts, state)

	case "inc":
		err = arithmetics.Inc(parts, state)

	case "sub":
		err = arithmetics.Sub(parts, state)

	case "dec":
		err = arithmetics.Dec(parts, state)

	case "mult":
		err = arithmetics.Mult(parts, state)

	case "div":
		err = arithmetics.Div(parts, state)

	case "store":
		err = memory.Store(parts, state)

	case "lfm":
		err = memory.Lfm(parts, state)

	case "echo":
		err = output.Echo(parts, state)

	case "dump":
		err = controlflow.Dump(state)

	case "exit":
		os.Exit(0)

	default:
		return fmt.Errorf("unknown command %s", parts[0])
	}

	return err
}
