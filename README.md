# TauRM
An experimental and educational register-based virtual machine with interpreted opcodes written in Go.

## What is this project about?
Tau Register Machine is a small VM/Runtime for Tau, its own language.

Tau is a statement-based language which may remind you of assembly, specifically made for the TauRM to interpret.

Note that this project is not to be taken too seriously, it is more of a fun little project than a real, performant runtime.

## Examples
Here are a few examples of programs written in Tau.

### Infinite Counter
```
ECHO Hello guys this is a demonstration for a simple program that counts.

LOOP:
    INC R0
    ECHO R0
    JMP LOOP # Counts infinitely with overflow
```

### If/Else Simulated
```
LOAD R0 2
LOAD R1 2

JIE R0 R1 EQUAL_BRANCH
JIS R0 R1 SMALLER_BRANCH
JIG R0 R1 GREATER_BRANCH

GREATER_BRANCH:
    ECHO Register 0 is greater than register 1
    JMP ANYHOW

SMALLER_BRANCH:
    ECHO Register 0 is smaller than register 1
    JMP ANYHOW

EQUAL_BRANCH:
    ECHO Register 0 and 1 are equal

ANYHOW:
    ECHO That was the test, Bye!
```

### Fibonacci Sequence
```
ECHO Fibonacci Demonstration

# Initialize Important Registers
LOAD R1 0    # X
LOAD R2 1    # Y
LOAD R3 0    # Result
LOAD R4 91   # Loop-Exit Condition
LOAD R5 0    # Counter

LOOP_START:
    ADD R1 R2 R3 
    ECHO R3
    COPY R2 R1
    COPY R3 R2
    INC R5
    JIS R5 R4 LOOP_START

ECHO Done
```

