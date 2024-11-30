package aurorium

import (
	"embed"

	"fyne.io/fyne/v2/lang"
)

//go:embed lang
var langFS embed.FS

func (p *Program) initLanguage() error {
	err := lang.AddTranslationsFS(langFS, "lang")
	if err != nil {
		return err
	}
	return nil
}
