---
title: "Product Planning in a Git Repo"
date: 2026-03-01
description: "How I moved our product planning into a GitHub repo with Claude Code as a copilot, and why the thinking gets sharper when it's structured like code."
tags: ["AI", "ClaudeCode", "ProductManagement", "GitHub", "Leadership"]
---

On a small team, every hour I spend shuffling planning docs across five tools is an hour I'm not reviewing a PR, tracking down a production issue, or building the next feature. I wanted that time back.

This past week I tried something different: I moved our product planning into a GitHub repo and started using [Claude Code](https://docs.anthropic.com/en/docs/claude-code/overview) as a product management copilot. Not Notion. Not Google Docs. Not a fancy PM tool. Markdown files in git.

## The Setup

Our repo has a simple structure. Roadmap files organized by funnel stage — Get Them In, Start a Trial, Deliver Value — plus UX and AI. Each six-week cycle gets a directory: proposal, Jira snapshot, frozen roadmap, analytics references. A CEO-facing prioritization doc. A daily operating pulse.

That's it. Markdown files, version controlled.

## The Moment It Clicked

Here's the story that sold me on this approach.

I used Claude Code to pull together data from across our stack — Jira tickets, Amplitude analytics, support feedback from Zendesk, Slack threads, existing Google Docs — and built a cycle proposal. Candidates grouped by funnel area, each with a "Why Now" backed by data where possible. I shared that with our CEO.

We got on a call. We recorded the meeting. We worked through the priorities together — some items shifted, others moved up, and we got aligned on what mattered most for the cycle.

After the call, I took the meeting transcript and gave it to Claude. "Here's the proposal. Here's what the CEO said. Update it."

Claude updated the proposal, adjusted what shifted, noted *why* things moved with the CEO's reasoning, and kept everything structured. I reviewed the diff, made a few tweaks, merged it.

That turnaround — from recorded meeting to updated proposal — took minutes instead of half a day. That's time back to building features, responding to production issues, and actually shipping. And every decision now has a commit. Every shift has a diff. Traceability is built in.

## Why This Works Better Than Traditional Tools

**Pull. Don't context switch.** It connects directly to Jira, Figma, Notion, Amplitude — so I'm pulling context into the planning repo instead of living across five tabs. The repo is the source of truth. Everything else is an input.

**Iterate like an engineer.** Branching, diffs, PRs. When I propose changes, I can gather more insight, ask better questions, and pull in data to inform the next decision. Once changes are solidified, the PR lets everyone see exactly what changed and why.

**Iterate from anywhere.** One thing that's been surprisingly powerful is being able to adjust plans without being at my desk. I can review a proposal on my phone, leave feedback, and the repo updates in real time. Eventually this moves to Slack. For now, [remote control](https://docs.anthropic.com/en/docs/claude-code/remote-control) gives us the fastest feedback loop.

**Code-aware planning.** The AI reads the actual codebase, not a wiki description of it. When I say "we need to coordinate interrupts," it finds the 15 places things fire independently, the boolean flag that's the only cross-system mechanism, and the specific files that need to change. Architecture grounded in reality, not assumptions.

**Living docs that compound.** Each session builds on previous work. The audit informs the architecture. The architecture informs the roadmap. The roadmap informs the cycle proposal. Same repo. Cross-linked. Compounding.

## Walk Before You Run

I want to be clear about something: I'm not just turning on agents and letting them go. This is a walk-before-run approach. I'm understanding the value each piece brings and making sure we're getting the best decisions out of it before automating anything.

Right now we have our current plan of record, future planning, and daily pulse check-ins running through Claude Code. The goal is to eventually automate the data inputs — pull from Jira, Amplitude, crash reports, support tickets — update the plan status, and create a PR for us to review. Full transparency on what changed. A history of why. And constant measurement against the impact we intended.

That's the future state: monitoring happens automatically, freeing our small team to focus on the real question — *what's the right product to build next* — and to measure each feature against the impact we intended.

## One More Thing

Moving all of this into a repo also gives us a natural place to put our operating procedures. How we use Amplitude. How we use Braze. AppsFlyer. Jira conventions. So when new product ideas come up, we have a checklist right next to the plan: are the right analytics in place? Are we set up to measure the value of what we're adding? That confidence is there before we ship, not after.

This pattern works beyond product planning too. I recently started a company and used a repo the same way — formation docs, tax filings, banking, vendors, deadlines, expenses — all in markdown, organized by topic. Claude Code helped me work through each step, and every decision and document has the same traceability. When a tax notice comes in via email, I hand it to the AI and it gets organized into the right place, deadlines get updated, and reminders are tracked. One place. Version controlled. Clear history.

## The Pattern

The specific tools don't matter as much as the pattern:

- **One repo for product direction** (not tickets — direction)
- **An AI that reads your codebase and your planning docs**
- **Integrations that pull context in**
- **Version control on decisions**, not just code

Jira is still where bugs and tasks live. Figma is still where designs live. But the thinking — the prioritization, the "why now," the architecture — that lives in markdown, in git, where it can be diffed, reviewed, and built on.

## What I Want Next

The next thing I'd love to bring in is a way to easily capture feedback from the team right where they already are. Once this is connected to Slack, conversations about a proposal can happen naturally there and get captured back into a PR. The feedback loop stays in the place people are already comfortable, and the repo stays the record of what was decided.

If you're a PM or engineering leader and you want more time building and less time syncing docs: try a repo. The thinking gets sharper when it's structured like code.
