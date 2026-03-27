---
title: "Claude Code Usage on a TRMNL E-Ink Display"
date: 2026-03-26
description: "A small project that puts your Claude Code usage stats on a TRMNL e-ink display — session and weekly limits at a glance."
tags: ["AI", "ClaudeCode", "TRMNL", "DeveloperTools"]
---

![TRMNL showing Claude usage stats](/images/claude-usage-trmnl.png)

I wanted a passive way to keep an eye on my Claude Code usage limits without opening a terminal. TRMNL's e-ink display seemed like a perfect fit — always on, zero distraction.

The result is a small Python script that runs on your Mac every 5 minutes, scrapes the `/usage` output from Claude Code, and pushes three metrics to a TRMNL private plugin:

- **Session** — current session usage percentage and reset time
- **Week - All Models** — weekly usage across all models
- **Week - Sonnet** — weekly Sonnet-specific usage

A `launchd` job handles the scheduling. No background daemons, no network listeners — just a cron-style script and an HTTPS POST to your own TRMNL webhook.

The plugin is live in the TRMNL Recipe Gallery: [Claude Usage - TRMNL](https://trmnl.com/recipes/263932)

Source and setup instructions: [carledwards/claude-usage-trmnl](https://github.com/carledwards/claude-usage-trmnl)
