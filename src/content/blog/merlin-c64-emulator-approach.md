---
title: "Merlin, the CPU, and the Commodore 64"
date: 2026-06-13
description: "Running Merlin on a C64 by making the CPU work: emulate the TMS1100, validate it against the Go core, and let the original game engine run."
tags: ["Merlin", "Commodore64", "Emulation", "TMS1100", "6502", "Go"]
---

I recently put together two Merlin projects that are really one idea in
two forms.

The first is [lets-go-merlin](https://github.com/carledwards/lets-go-merlin),
a Go emulator for **Merlin - The Electronic Wizard**. It runs the original
Texas Instruments MP3404 ROM on a TMS1100 interpreter, and that same core
is available as a [browser version](https://carledwards.github.io/lets-go-merlin/).

The newer one is [merlin-c64](https://github.com/carledwards/merlin-c64),
which runs Merlin on the Commodore 64.

That sounds like a remake, but it is not the interesting part. The fun
part is that the C64 version is still running the original Merlin program.
The C64 is pretending to be the TMS1100. A 6502 assembly interpreter
fetches Merlin's real instructions, updates the emulated chip state, and
lets the original engine decide which lights turn on, which buttons mean
what, and which game logic runs next.

![Merlin running on the Commodore 64](/images/projects/merlin-c64/hero.gif)

## Make the CPU work

The thing I like about this approach is that it avoids writing a bunch of
duplicative code to make something *look* like Merlin.

I did not need to rewrite tic-tac-toe, Echo, Music Machine, Blackjack 13,
Magic Square, or Mindbender for the C64. I needed the CPU behavior to be
good enough that the original Merlin ROM could run. Once that boundary is
right, the behavior emerges from the program that was already there.

That changes the shape of the work. Instead of asking "how do I recreate
this game?", the question becomes "what does this CPU actually do, and can
I make another machine execute those rules faithfully?"

That is a much more interesting question to me.

It also makes the limitations honest. The real Merlin TMS1100 ran at about
350 kHz, which works out to roughly 58,000 instructions per second. A C64
has a faster headline clock than that, but its 6502 is not a TMS1100. It
has to spend many 6502 instructions to emulate one Merlin instruction. The
result is that the C64 version runs in slow motion.

That is not a failure of the idea. It is the tradeoff made visible. The C64
can run the real engine, but it cannot do it at full speed without help.
Some parts, like fixed startup audio, can be handled more cleverly through
the SID. Other parts stay slow because they are driven by the emulated ROM's
timing loops.

## Testing the port

The part I care about most is the validation loop.

The Go version already gives me a clean TMS1100 core. For the C64 version,
I ported that interpreter into 6502 assembly, but I did not want to trust
"it seems to work" as proof. So `merlin-c64` has a lockstep test.

The test assembles the C64 program, loads it into a neutral 6502 emulator,
and runs it one TMS1100 instruction at a time. After each emulated TMS1100
instruction, it compares the C64 interpreter's state against the Go core
running the same Merlin ROM.

Registers, program counter, page state, output latch, and the LED-driving
R lines all have to match. If they drift, the test fails at the exact step
where the port stopped behaving like the reference.

That is the payoff of separating the idea from the target machine. The Go
core is not just a web toy. It becomes a reference implementation. The
C64 version is not just hand-written assembly. It is a target backend that
has to prove it still implements the same CPU.

## The same pattern applies elsewhere

This is the same technique I like in the 6502 work too.

You can run the 6502 simulator separately in tests, validate the assembled
code against the behavior you expect, and then move that code onto a real
C64 or C64-compatible device with a lot more confidence. The simulator is
not the end product. It is a controlled environment for proving the small
parts before the hardware gets involved.

That matters because old systems are very constrained. The constraints are
not a reason to romanticize them. They are useful because they make the
machine boundary sharp. You can see where the CPU ends, where the display
starts, where timing leaks into sound, and where a shortcut stops being
faithful.

I do not think of this as being stuck in the past. I think of it as using
small, inspectable systems to practice a discipline that still matters:
define the contract, make the implementation obey it, test it against an
independent reference, and then run it somewhere less forgiving.

In this case, "somewhere less forgiving" happens to be a Commodore 64.
That is part of the charm, but not the whole point.

The point is that if the CPU works, the original program comes with it.
