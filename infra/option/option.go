package option

/*
	echarts 选项： 窗口大小、主题、数据、

*/

// the window size : Width and Height
// WindowsSize...
type WindowSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func DefaultWS() *WindowSize {
	return &WindowSize{
		Width:  800,
		Height: 600,
	}
}

// the theme of echarts: dark or light
// Theme ...
type Theme struct {
	Name string `json:"name"`
}

func DefaultTheme() *Theme {
	return &Theme{
		Name: "dark",
	}
}

func (t *Theme) SetTheme(name string) *Theme {
	if name == "" {
		panic("theme should not be nil")
	}
	return &Theme{
		Name: name,
	}
}

/******************************************************************************/

type Title struct {
	Text      string    `json:"text"`
	Show      bool      `json:"show"`
	TextStyle TextStyle `json:"textStyle"`
}

type TextStyle struct {
	Color string `json:"color"`
}

type ToolTip struct {
}

type Legend struct {
	Data []string `json:"data"`
}

type XAxis struct {
	Name string   `json:"name"`
	Data []string `json:"data"`
}

type YAxis struct {
	Name string `json:"name"`
}

type Series struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	Data []string `json:"data"`
}

type ChartBase struct {
	Title
	ToolTip
	Legend
	XAxis
	YAxis
	Series
}

func (c *ChartBase) NewChart(title string) *ChartBase {
	return &ChartBase{
		Title: Title{Text: title},
	}
}
