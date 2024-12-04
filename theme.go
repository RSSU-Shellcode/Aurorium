package aurorium

import (
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type customTheme struct {
	fyne.Theme

	// custom fonts
	bold    fyne.Resource
	regular fyne.Resource
}

func (p *Program) initTheme() error {
	t := customTheme{
		Theme: theme.Current(),
	}
	for _, item := range []struct {
		name string
		res  *fyne.Resource
	}{
		{"NotoSansSC-Regular.ttf", &t.regular},
		{"NotoSansSC-Bold.ttf", &t.bold},
	} {
		data, err := os.ReadFile(filepath.Join("font", item.name))
		if err != nil {
			return err
		}
		*item.res = fyne.NewStaticResource(item.name, data)
	}
	p.app.Settings().SetTheme(&t)
	return nil
}

func (t *customTheme) Font(style fyne.TextStyle) fyne.Resource {
	switch {
	case style.Bold:
		return t.bold
	case style.Italic, style.Monospace, style.Symbol, style.Underline:
		return t.Theme.Font(style)
	default:
		return t.regular
	}
}

func (t *customTheme) Size(name fyne.ThemeSizeName) float32 {
	switch {
	// disable rounded corners style
	case strings.Contains(strings.ToLower(string(name)), "radius"):
		return 0
	// adjust general text size
	case string(name) == "text":
		return 14
	default:
		return t.Theme.Size(name)
	}
}
