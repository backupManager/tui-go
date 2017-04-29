package main

import (
	tui "github.com/marcusolsson/tui-go"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	t := tui.NewTable(0, 0)
	t.AppendRow(tui.NewLabel("First row"))
	t.AppendRow(tui.NewLabel("Second row"))

	root := tui.NewVBox(t, tui.NewSpacer())

	// This is still a rough prototype and WILL change. Use at your own risk.
	p := tui.NewPalette()
	p.SetItem("table.cell.selected", tui.PaletteItem{
		Bg: tui.Color(termbox.ColorRed),
		Fg: tui.Color(termbox.ColorWhite),
	})

	ui := tui.New(root)
	ui.SetPalette(p)
	ui.SetKeybinding(tui.KeyEsc, func() {
		ui.Quit()
	})
	if err := ui.Run(); err != nil {
		panic(err)
	}
}