package webapi

import (
	"fmt"

	translator "github.com/Conight/go-googletrans"

	"github.com/Yuto-M/go-clean-template/internal/entity"
)

// TranslationWebAPI -.
type TranslationWebAPI struct {
	conf translator.Config
}

// New -.
func New() *TranslationWebAPI {
	conf := translator.Config{}

	return &TranslationWebAPI{
		conf: conf,
	}
}

// Translate -.
func (t *TranslationWebAPI) Translate(translation entity.Translation) (entity.Translation, error) {
	trans := translator.New(t.conf)

	result, err := trans.Translate(translation.Original, translation.Source, translation.Destination)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	translation.Translation = result.Text

	return translation, nil
}
