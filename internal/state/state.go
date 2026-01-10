package state

import "fmt"

type State struct {
	Registers 	map[string]int 	// 16 Registers, 64 or 32-bit each, depending on target architecture
	Memory 		map[uint]int 	// Simple RAM
	Labels 		map[string]int	// Labels for jump commands
	Pc 			int 			// Program Counter
}

func Initialize() *State {
	s := State{
		Registers: map[string]int{},
		Memory: map[uint]int{},
		Labels: map[string]int{},
		Pc: 1,
	}
	
	// Initialize Registers
	for i := range 15 {
		s.Registers[fmt.Sprintf("R%d", i)] = 0 
	}

	return &s
}

func (s *State) HasReg(reg string) bool {
	_, ok := s.Registers[reg]
	return ok
}
