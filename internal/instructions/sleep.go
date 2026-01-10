package instructions

import (
	"fmt"
	"time"
	"strconv"
)

func Sleep(parts []string) error {
	if len(parts) != 2 {
		return fmt.Errorf("invalid argument count, expected 2, got %d", len(parts))
	}

	dur, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("could not parse %s to int", parts[1])
	}

	time.Sleep(time.Duration(dur * int(time.Millisecond)))
	return nil
}
