# Commodore 64 Emulator
Huge thanks to the [Ultimate Commodore 64 Reference](https://github.com/mist64/c64ref), which I used to auto-generate a lot of code.

## To run
``` bash
go run ./
```

Emulator will start in debug mode:
```
------------+
instruction | LDX #$FF (load)
registers   | PC: FCE2 S: FF A: 00 X: 00 Y: 00
status      | N: 0 V: 0 -: 0 B: 0 D: 0 I: 1 Z: 0 C: 0
memory      | FCE2: A2 FF 78 9A D8 20 02 FD
interrupts  | Pending: false Vector: FFFC
------------+
debug>
```

Use the `help` command to list commands.