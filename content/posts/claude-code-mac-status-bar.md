---
title: "Adding Claude Code Status to Your Mac Status Bar"
date: 2026-03-14
description: "How to set up a Mac menu bar indicator that shows when Claude Code is idle, working, or needs your attention using SwiftBar and hooks."
tags: ["AI", "ClaudeCode", "MacOS", "DeveloperTools"]
---

If you're anything like me, you fire off a task in Claude Code and then flip over to something else — a browser tab, Slack, another terminal. The problem is knowing when Claude is done or needs your attention without constantly checking back.

I put together a simple setup that adds a status indicator right in the Mac menu bar. Three states, zero guesswork:

- 🤖 **Idle/Done** — Claude is finished or waiting
- 🔨 **Working** — Claude is actively using tools
- 🔔 **Needs Attention** — Claude has a question or notification for you

Here's how to set it up.

## Step 1: Install SwiftBar

[SwiftBar](https://github.com/swiftbar/SwiftBar) is a lightweight, open-source menu bar customization tool. All it does is run shell scripts on an interval and display their output in your menu bar — no network calls, no telemetry. You can review the full source on [GitHub](https://github.com/swiftbar/SwiftBar) if you want to see exactly what it's doing under the hood. Install it with Homebrew:

```bash
brew install swiftbar
```

## Step 2: Create a Plugins Directory

SwiftBar needs a directory to look for plugin scripts. Create one:

```bash
mkdir -p ~/.swiftbar
```

## Step 3: Create the Status Script

Create a file at `~/.swiftbar/claude-status.3s.sh` with the following contents:

```bash
#!/bin/bash
FLAG=/tmp/claude-status-flag

if [ -f "$FLAG" ]; then
  case $(cat "$FLAG") in
    notify) echo "🔔 Claude" ;;
    working) echo "🔨Claude" ;;
    done)   echo "🤖 Claude" ;;
  esac
else
  echo "🤖"   # idle state - minimal
fi
```

The `.3s.` in the filename tells SwiftBar to refresh the script every 3 seconds. Make it executable:

```bash
chmod +x ~/.swiftbar/claude-status.3s.sh
```

## Step 4: Add Hooks to Claude Code

Claude Code supports [hooks](https://docs.anthropic.com/en/docs/claude-code/hooks) — shell commands that fire in response to lifecycle events. Open your Claude settings file at `~/.claude/settings.json` and add the following `hooks` section (if you already have content in this file, add this before the closing `}`):

```json
"hooks": {
  "Notification": [{
    "matcher": "*",
    "hooks": [{
      "type": "command",
      "command": "echo 'notify' > /tmp/claude-status-flag && open -g 'swiftbar://refreshPlugin?name=claude-status'",
      "timeout": 2
    }]
  }],
  "PreToolUse": [{
    "matcher": "*",
    "hooks": [{
      "type": "command",
      "command": "echo 'working' > /tmp/claude-status-flag && open -g 'swiftbar://refreshPlugin?name=claude-status'",
      "timeout": 2
    }]
  }],
  "Stop": [{
    "matcher": "*",
    "hooks": [{
      "type": "command",
      "command": "echo 'done' > /tmp/claude-status-flag && open -g 'swiftbar://refreshPlugin?name=claude-status'",
      "timeout": 2
    }]
  }]
}
```

Each hook writes a state to a temp file and tells SwiftBar to refresh immediately (rather than waiting for the next 3-second poll). The `open -g` flag opens the URL in the background so it doesn't steal focus.

## Step 5: Launch SwiftBar and Set the Plugin Directory

Open SwiftBar and point it to the plugin directory you created.

**Tip:** If you used `~/.swiftbar`, the directory will be hidden in Finder. When the "Choose plugin folder" dialog appears, press **⇧⌘G** (Shift + Command + G) to open the "Go to folder" input, then type `~/.swiftbar` to navigate there directly.

## Testing It Out

You can verify everything works without running Claude Code by executing one of the hook commands directly in your terminal:

```bash
echo 'notify' > /tmp/claude-status-flag && open -g 'swiftbar://refreshPlugin?name=claude-status'
```

You should see the 🔔 icon appear in your menu bar. Swap `notify` for `working` or `done` to test the other states.

Once it's wired up, you'll always know at a glance whether Claude is heads-down, finished, or waiting on you — no terminal switching required.
