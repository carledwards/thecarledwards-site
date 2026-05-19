---
title: "From a Single Bit to a Real CPU: The Learn Series"
date: 2026-05-19
description: "Why I built a hands-on series that takes you from one bit up to a working 8-bit CPU running real code in the browser — and why it's built AI-forward."
tags: ["Learn", "6502", "Binary", "AI", "Teaching"]
---

It's been quiet here for a bit. Not idle — heads-down. A lot has shipped
since my last post, and the part I'm most happy with is something I can
finally point people at: a [Learn series](/learn) on the site.

The one-breath version: it starts at a single bit, done by hand, and
builds up — binary, a tiny model of a computer, and then a real 8-bit
6502 CPU you assemble and run, right in the browser, watching it drive
actual hardware.

## Why I made it

Who doesn't love blinking LEDs and buttons you can click?

The first binary counter I ever saw was in 6th grade — a row of little
lights ticking up, and I had no idea what it was doing. I didn't have
the words for it. I just knew I couldn't stop watching it, and I wanted
to know *why* it did that. Hooked, basically. That feeling — not "this
is hard," but "wait, what *is* this?" — is the one I'm trying to hand to
someone else.

I also just like to teach. The thing I've noticed, over and over, is
that what actually demystifies an idea for me isn't a better
definition — it's seeing the same thing from a few different angles
until one of them lands. So this is the angle I wish I'd had back then:
not a page *about* binary, a binary counter you can poke.

The only honest way to know if that's working is feedback. More than
one person has told me, plainly, that they "never got binary" — then
went through this, and now they do. That's the win. Not that the
material is impressive; that someone who'd given up on it walks away,
an hour later, understanding more than they did.

And honestly, it was never really about binary. If poking at this
leaves someone a little more curious — a little more willing to take a
thing apart to see how it works, or to come at a problem from a new
angle — that's the part I actually care about. The 6502 is just a great
excuse.

## The approach: by hand, by doing

So the lessons run the other way around. By hand. By doing. The
smallest real version of each idea, built by you — not described to
you. You click something, change one thing, watch the machine react,
and form a guess about why. The understanding tends to show up about a
step after you stop waiting for it.

That's why the 6502 part isn't a simulator demo you watch. You write
the assembly, it assembles in your browser, it runs on a faithful 6502,
and eight LEDs light up from a byte you stored. The same program would
run unmodified on the real chip on a desk. Small enough to fully
understand, real enough to matter.

## Built AI-forward, on purpose

The other deliberate choice: this is built for how people actually
learn now. Every lab has a **Copy for AI** button. When you want to go
deeper than a lesson goes, or you're just stuck, you grab the full
context — your code, the machine, the syntax reference, what the
assembler said — and take it into whatever assistant you like, then
come back and keep going.

The series is also a springboard. The pieces underneath it — the
assembler, the simulator, the way it all runs in a browser — are the
foundation for the projects I'm building next. More on those when
they're ready. Today the series stands on its own, and that's where I'd
start.

## Go poke at it

Pick anywhere that looks interesting and start —
[the lessons are here](/learn). You don't need to have written code
before. You don't need to "be a math person." Even the trying is a
positive thing; that part is underrated.

Go break something and figure out why. That's the good part.
