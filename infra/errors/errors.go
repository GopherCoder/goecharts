package errors

type ErrorChart struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var ErrorCharts []ErrorChart

var (
	errTheme = &ErrorChart{400, "theme should not be nil", "theme should in (dark, light)"}
)
