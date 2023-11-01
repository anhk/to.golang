package app

import (
	"github.com/rivo/tview"
	"time"
)

// ┌───────────────────────────Search────────────────────────────────┐
// │        auto complete search                                     │
// ├─────────────────────────────────────────────────────────────────┤
// ├────────────────────────────────────────────────├┤───────────────┤
// │                                                ││               │
// │   1) x.x.x.x  (root)                           ││  Help Message │
// │   2) y.y.y.y  (root)                           ││   A) Add ...  │
// │                                                ││   D) Del ...  │
// │                                                ││   U) Update   │
// │                                                ││   /) Search   │
// │                                                ││   ?) Info     │
// │                                                ││               │
// │                                                ││ History       │
// │                                                ││  !1 x.x.x.x   │
// │                                                ││  !2 y.y.y.y   │
// │                                                ││  !3 z.z.z.z   │
// │                                                ││               │
// ├────────────────────────────────────────────────┼│               │
// │     Promote                                    ││               │
// └────────────────────────────────────────────────┴┴───────────────┘

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
	this.mainPanel.AddItem(list, 0, 1, true)

	this.helpPanel = tview.NewFlex()
	this.helpPanel.SetBorder(true).SetTitle(" Help ")

	this.screen = tview.NewFlex().SetDirection(tview.FlexRow)
	this.screen. //AddItem(this.searchPanel, 3, 3, false).
			AddItem(this.mainPanel, 0, 8, true).
			AddItem(this.helpPanel, 4, 0, false)

	//this.screen.RemoveItem(this.searchPanel)
	this.app = tview.NewApplication().SetRoot(this.screen, true).EnableMouse(true)
	//this.screen.AddItem(this.searchPanel, 3, 3, true)
	//this.app.SetFocus(this.searchPanel)

	return this
}

func (app *App) Run() {

	go func() {
		time.Sleep(3 * time.Second)
		app.screen.Clear()
		app.screen.AddItem(app.searchPanel, 3, 0, true)
		app.screen.AddItem(app.mainPanel, 0, 8, false)
		app.screen.AddItem(app.helpPanel, 4, 0, false)
		app.app.SetFocus(app.searchPanel)
		app.app.Draw()
	}()
	_ = app.app.Run()
}
