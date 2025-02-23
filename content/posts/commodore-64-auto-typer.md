---
title: "Building a Modern Commodore 64 Auto-Typer"
date: 2025-02-23
description: "A modern take on automating Commodore 64 program entry using ESP32 and Teensy microcontrollers, bridging retro computing with contemporary hardware."
tags: ["Commodore64", "RetroComputing", "ESP32", "Teensy", "Electronics", "Programming", "6502"]
---

My journey in programming began with the Commodore 64 back in the 1980s. Like many developers of that time, I started with BASIC before diving into 6502 assembly language. That early experience led me to help design the [LCM-32 (Low Cost Microcontroller)](https://github.com/carledwards/lcm-32) in 1986, where I worked alongside Matt Gilliland to create an affordable 6502-based single-board computer.

That passion for the 6502 architecture has stayed with me through the years, leading to various emulation and simulation projects:

- A [transistor-level 6502 simulator in Zig](https://github.com/carledwards/6502-netsim-zig)
- Implementations in [Go](https://github.com/carledwards/6502-netsim-go) and [Nim](https://github.com/carledwards/6502-netsim-nim)
- A [pure Python C64 simulator](https://github.com/carledwards/pyc64)
- A [6502 CPU simulator](https://github.com/carledwards/6502-simulator) with virtual motherboard support

Now, I'm excited to share my latest project that brings together my love for retro computing and modern microcontrollers: the Commodore 64 Auto-Typer.

## The Project

The C64 Auto-Typer combines an ESP32 and Teensy microcontroller to automatically input BASIC programs and commands into an Ultimate 64. It's designed to bridge the gap between modern development workflows and the classic C64 experience.

![C64 Auto-Typer Demo](/images/commodore-64-auto-typer/demo.gif)

## How It Works

The system uses:
- An ESP32 to store and process BASIC programs
- A Teensy 3.2 for USB HID keyboard emulation
- The Ultimate 64's USB keyboard interface

Here's the prototype on a breadboard:

![C64 Auto-Typer Breadboard](/images/commodore-64-auto-typer/breadboard.jpg)

When triggered, it converts your BASIC program into the appropriate keystrokes, handling all the special PETSCII characters and control codes that make C64 programming unique.

## Features

- Support for all C64 color codes and control characters
- Function key mapping
- Configurable typing speed
- Status LED feedback
- Simple button interface

## Building Your Own

The complete project is available on [GitHub](https://github.com/carledwards/commodore-64-auto-typer). You'll need:

### Hardware
- ESP32 development board
- Teensy 3.2
- Ultimate 64
- Push button
- Some connecting wires
- USB cable

### Software
- MicroPython for the ESP32
- Arduino IDE for the Teensy
- The project source code

## Future Plans

I'm planning several enhancements:
- Web interface for program upload
- Support for saving to disk images
- Multiple program storage and selection
- Integration with popular C64 BASIC repositories

## Full Circle

It's fascinating how life comes full circle. From learning to program on a C64, to designing the LCM-32, to creating various 6502 simulators, and now back to enhancing the C64 experience with modern technology. This project represents not just a useful tool, but a bridge between the computing worlds that have shaped my journey as a developer.

Check out the full source code and documentation on [GitHub](https://github.com/carledwards/commodore-64-auto-typer), and feel free to contribute or suggest improvements! 