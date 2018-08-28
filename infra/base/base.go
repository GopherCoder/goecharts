package base

import "goecharts/infra/option"

type Base struct {
	WindowSize option.WindowSize // window size

	Theme option.Theme // theme

	Title string // title

	SubTitle string // subtitle

}

type Charts interface {
	SetTheme(string) // set theme

	SetWindowSize(size option.WindowSize) // set window size

	HTMLRender() // show html by server

	Render(string) // show html

	Config(map[string][]interface{}) // config

	Add(map[string]interface{}) // add option
}
