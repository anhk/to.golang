package app

import (
	"github.com/rivo/tview"
)

type App struct {
	searchPanel *tview.InputField
	mainPanel   *tview.Flex
	helpPanel   *tview.Flex

	screen *tview.Flex
	app    *tview.Application
}

func NewApp() *App {
	this := &App{}

	list := tview.NewList().ShowSecondaryText(false)
	list.AddItem("192.168.64.53   ---   (root)", "", '3', func() {
		//a, b := list.GetItemText(list.GetCurrentItem())
		//fmt.Println(a, b)
	}).AddItem("192.168.64.52   ---   (root)", "", '4', func() {
		//a, b := list.GetItemText(list.GetCurrentItem())
		//fmt.Println(a, b)
	})
	this.searchPanel = tview.NewInputField().SetAutocompletedFunc(func(text string, index int, source int) bool {
		return false
	})

	this.searchPanel.SetBorder(true).SetTitle(" Search ")

	this.mainPanel = tview.NewFlex()
	this.mainPanel.SetBorder(true).SetTitle(" Main ")
	this.mainPanel.AddItem(list, 0, 1, false)

	this.helpPanel = tview.NewFlex()
	this.helpPanel.SetBorder(true).SetTitle(" Help ")

	this.screen = tview.NewFlex().SetDirection(tview.FlexRow)
	this.screen. //AddItem(this.searchPanel, 3, 3, false).
			AddItem(this.mainPanel, 0, 8, false).
			AddItem(this.helpPanel, 0, 1, false)

	this.screen.RemoveItem(this.searchPanel)
	this.app = tview.NewApplication().SetRoot(this.screen, true).EnableMouse(true)

	return this
}

func (app *App) Run() {
	_ = app.app.Run()
}
