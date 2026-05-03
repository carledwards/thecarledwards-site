//go:build js && wasm

// Package main is the site-specific FoxPro Mode wrapper.
//
// Foxpro-go provides the windowing framework; this file decides what
// the site surfaces — menus, commands, content. The seam between Go
// and the host page is window.siteAPI, defined in FoxproOverlay.astro
// before the wasm boots.
package main

import (
	"strings"
	"syscall/js"
	"time"

	foxpro "github.com/carledwards/foxpro-go"
	"github.com/carledwards/foxpro-go/wasm"
	"github.com/gdamore/tcell/v2"
)

func main() {
	s := tcell.NewSimulationScreen("UTF-8")
	if err := s.Init(); err != nil {
		panic(err)
	}
	// Classic-DOS feel: a touch wider than 80×25 to fit the menu bar +
	// hint strip without crowding the desktop.
	s.SetSize(90, 26)
	s.EnableMouse()

	app := foxpro.NewAppWithScreen(s)

	// Browser-appropriate settings.
	app.Settings.QuitKeys = nil
	app.Settings.BackgroundDragChords = []foxpro.BackgroundDragChord{
		{Button: tcell.Button1, Mods: tcell.ModShift},
	}
	// The default status hint advertises terminal-only keys (F1, F2,
	// F10, F6). In a browser overlay the only thing users need to
	// know is how to dismiss it.
	app.Settings.StatusBarLeft = " Esc to close "

	// Centered ASCII-art "Carl Edwards" sitting behind the windows.
	// The default desktop hatch still fills the negative space — we
	// only paint non-blank cells of the art.
	app.DesktopDraw = drawBackdropArt

	// Single command window for the whole session — pre-populated with
	// the welcome echo and reused if the user hides/reopens it.
	cmdWin := foxpro.NewCommandWindow(app)
	cmdWin.Bounds = foxpro.Rect{X: 1, Y: 2, W: 42, H: 14}
	cmdWin.OnClose = func() { app.Manager.Remove(cmdWin) }
	cp := cmdWin.Content.(*foxpro.CommandProvider)
	// Welcome echo, sized to fit a 35-col window (33-col inner area).
	cp.Print(`  Welcome to Carl Edwards' site,`)
	cp.Print(`  in FoxPro Mode.`)
	cp.Print("")
	cp.Print(`  Type 'help' for commands.`)
	cp.Print(`  Or use the menus above.`)
	cp.Print("")
	cp.Print(`  Press Esc to dismiss.`)
	cp.Print("")
	cp.Print("")
	cp.AppendInputLine() // park cursor on a fresh blank line below
	app.Manager.Add(cmdWin)

	// showCmdWindow opens the command window if it's been closed and
	// always raises it to the top + focuses it. Used by the File menu
	// and by F2. Never hides — closing happens via the ■ on the window
	// or Esc on the overlay.
	showCmdWindow := func() {
		if !app.Manager.Contains(cmdWin) {
			app.Manager.Add(cmdWin)
		}
		app.Manager.Raise(cmdWin)
	}

	app.OnKey = func(ev *tcell.EventKey) bool {
		if ev.Key() == tcell.KeyF2 {
			showCmdWindow()
			return true
		}
		return false
	}

	registerCommands(app)
	setupMenus(app, showCmdWindow)

	// Drive a 500ms heartbeat so the command window's cursor blink
	// animates. Tick posts a no-op event each interval, which wakes
	// foxpro's PollEvent loop and triggers a redraw.
	app.Tick(500*time.Millisecond, nil)

	// Greeting toast on first open. Auto-dismisses after 5s; any
	// key/mouse press dismisses sooner. Theme/position/dismiss-rules
	// all come from the framework; we just supply the message + life.
	app.ShowWaitWindow(&foxpro.WaitWindow{
		Message: "Hello!",
		Timeout: 5 * time.Second,
	})

	wasm.Run(app, s)
}

// command is a site command paired with its help text. Kept as a
// table so HELP can iterate without needing access to foxpro-go's
// internal CommandRegistry sort.
type command struct {
	name    string
	help    string
	handler foxpro.CommandFunc
}

func siteCommands(a *foxpro.App) []command {
	return []command{
		{"HELP", "List available commands", func(cp *foxpro.CommandProvider, args string) {
			for _, c := range siteCommands(a) {
				cp.Print("  " + padName(c.name, 12) + c.help)
			}
		}},
		{"CLS", "Clear this window", func(cp *foxpro.CommandProvider, args string) {
			cp.Clear()
		}},
		{"ABOUT", "About FoxPro Mode", func(cp *foxpro.CommandProvider, args string) {
			openAboutFoxpro(a)
		}},
		{"CARL", "About Carl Edwards", func(cp *foxpro.CommandProvider, args string) {
			openAboutCarl(a)
		}},
		{"LINKEDIN", "Open Carl's LinkedIn", func(cp *foxpro.CommandProvider, args string) {
			cp.Print("Opening LinkedIn…")
			openURL("https://www.linkedin.com/in/the-carl-edwards/")
		}},
		{"GITHUB", "Open Carl's GitHub", func(cp *foxpro.CommandProvider, args string) {
			cp.Print("Opening GitHub…")
			openURL("https://github.com/carledwards")
		}},
		{"RSS", "Open the RSS feed", func(cp *foxpro.CommandProvider, args string) {
			cp.Print("Opening RSS feed…")
			openURL("/rss.xml")
		}},
		{"PROJECTS", "Go to the projects page", func(cp *foxpro.CommandProvider, args string) {
			cp.Print("Heading to projects…")
			navigate("/projects/")
		}},
		{"BLOG", "Go to the writing index", func(cp *foxpro.CommandProvider, args string) {
			cp.Print("Heading to the blog…")
			navigate("/blog/")
		}},
		{"WAIT", "Show a sample notification", func(cp *foxpro.CommandProvider, args string) {
			msg := strings.TrimSpace(args)
			if msg == "" {
				msg = "Hello from FoxPro Mode!"
			}
			a.ShowWaitWindow(&foxpro.WaitWindow{Message: msg})
		}},
		{"EXIT", "Close FoxPro Mode", func(cp *foxpro.CommandProvider, args string) {
			closeOverlay()
		}},
	}
}

func registerCommands(a *foxpro.App) {
	for _, c := range siteCommands(a) {
		a.Commands.Register(c.name, c.help, c.handler)
	}
}

func setupMenus(a *foxpro.App, showCmdWindow func()) {
	a.MenuBar = foxpro.NewMenuBar([]foxpro.Menu{
		{
			Label: "&File",
			Items: []foxpro.MenuItem{
				{Label: "&Command Window", Hotkey: "F2", OnSelect: showCmdWindow},
				{Separator: true},
				{Label: "Cl&ose (Esc)", OnSelect: closeOverlay},
			},
		},
		{
			Label: "&Connect",
			Items: []foxpro.MenuItem{
				{Label: "&LinkedIn", OnSelect: func() {
					openURL("https://www.linkedin.com/in/the-carl-edwards/")
				}},
				{Label: "&GitHub", OnSelect: func() {
					openURL("https://github.com/carledwards")
				}},
				{Label: "&RSS", OnSelect: func() { openURL("/rss.xml") }},
			},
		},
		{
			Label: "&Site",
			Items: []foxpro.MenuItem{
				{Label: "&Home", OnSelect: func() { navigate("/") }},
				{Label: "&Projects", OnSelect: func() { navigate("/projects/") }},
				{Label: "&Blog", OnSelect: func() { navigate("/blog/") }},
			},
		},
		{
			Label: "&Help",
			Items: []foxpro.MenuItem{
				{Label: "&About FoxPro Mode", OnSelect: func() { openAboutFoxpro(a) }},
				{Label: "About &Carl Edwards", OnSelect: func() { openAboutCarl(a) }},
			},
		},
	})
}

func openAboutCarl(a *foxpro.App) {
	body := foxpro.NewTextProvider([]string{
		"Carl Edwards",
		"",
		"CTO, architect, hands-on builder.",
		"",
		"I write about engineering leadership, AI workflows,",
		"and developer tools.",
		"",
		"Most of what I'm tinkering with lives at",
		"github.com/carledwards.",
	})
	a.Manager.Add(foxpro.NewWindow(
		"About Carl Edwards",
		foxpro.Rect{X: 14, Y: 6, W: 56, H: 13},
		body,
	))
}

func openAboutFoxpro(a *foxpro.App) {
	body := foxpro.NewTextProvider([]string{
		"What is FoxPro Mode?",
		"",
		"This is a Go program compiled to WebAssembly.",
		"The same code runs unchanged in a real terminal.",
		"It's built on foxpro-go, a small framework for",
		"DOS-style text-mode UIs.",
		"",
		"The look and feel is borrowed from FoxPro for DOS",
		"(Microsoft, 1991) — one of the few DOS applications",
		"that pulled off a real mouse-driven GUI:",
		"overlapping windows, drop shadows, drag and resize,",
		"menu bars and pop-up panels — all rendered in 80×25",
		"character mode, years before Windows took over.",
		"",
		"It's a small tribute to that era, and a fun way to",
		"explore my site.",
	})
	a.Manager.Add(foxpro.NewWindow(
		"About FoxPro Mode",
		foxpro.Rect{X: 12, Y: 4, W: 64, H: 20},
		body,
	))
}

// padName right-pads a command name to width with spaces. Used by
// HELP to align command descriptions in a column.
func padName(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}

// backdropArt is "Carl Edwards" rendered in lower-half-block (▄) glyphs.
// 80 cols wide, 12 rows tall — sits comfortably inside a 90-col canvas
// with 5 cols of margin on each side.
var backdropArt = []string{
	" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄                    ▄▄▄▄▄▄▄▄▄▄▄                                ",
	"   ▄▄         ▄▄▄                      ▄▄       ▄▄                              ",
	"   ▄▄           ▄                      ▄▄        ▄▄                             ",
	"   ▄▄                                  ▄▄        ▄▄                             ",
	"   ▄▄          ▄▄▄      ▄▄▄▄▄  ▄▄▄▄▄   ▄▄       ▄▄  ▄▄▄   ▄▄▄▄         ▄▄▄      ",
	"   ▄▄▄▄▄    ▄▄     ▄▄     ▄▄    ▄▄     ▄▄▄▄▄▄▄▄▄      ▄▄ ▄▄   ▄▄    ▄▄     ▄▄   ",
	"   ▄▄      ▄▄       ▄▄     ▄▄  ▄▄      ▄▄             ▄▄▄          ▄▄       ▄▄  ",
	"   ▄▄     ▄▄         ▄▄     ▄▄▄        ▄▄             ▄▄          ▄▄         ▄▄ ",
	"   ▄▄      ▄▄       ▄▄     ▄▄ ▄▄       ▄▄             ▄▄           ▄▄       ▄▄  ",
	"   ▄▄       ▄▄     ▄▄     ▄▄   ▄▄      ▄▄             ▄▄            ▄▄     ▄▄   ",
	" ▄▄▄▄▄▄▄       ▄▄▄      ▄▄▄▄▄  ▄▄▄▄  ▄▄▄▄▄▄         ▄▄▄▄▄▄             ▄▄▄      ",
	"                                                                                ",
}

// drawBackdropArt renders backdropArt centered in the desktop area.
// Only non-space cells are written so the theme's desktop pattern
// shows through the negative space. Style is white-on-blue — clean
// high-contrast block letters against the deep FoxPro-blue field.
func drawBackdropArt(s tcell.Screen, area foxpro.Rect, theme foxpro.Theme) {
	const artW = 80
	artH := len(backdropArt)
	x0 := area.X + (area.W-artW)/2
	y0 := area.Y + (area.H-artH)/2
	if x0 < area.X {
		x0 = area.X
	}
	if y0 < area.Y {
		y0 = area.Y
	}
	style := tcell.StyleDefault.
		Background(theme.Palette.Blue).
		Foreground(theme.Palette.White)
	for i, line := range backdropArt {
		y := y0 + i
		if y >= area.Y+area.H {
			break
		}
		x := x0
		for _, r := range line {
			if x >= area.X+area.W {
				break
			}
			if r != ' ' {
				s.SetContent(x, y, r, nil, style)
			}
			x++
		}
	}
}

// openURL asks the host page to open a URL in a new tab.
func openURL(u string) {
	js.Global().Get("siteAPI").Call("openURL", u)
}

// navigate replaces the current page (in-site destinations).
func navigate(path string) {
	js.Global().Get("siteAPI").Call("navigate", path)
}

// closeOverlay hides the foxpro canvas without terminating the wasm.
func closeOverlay() {
	js.Global().Get("siteAPI").Call("close")
}
