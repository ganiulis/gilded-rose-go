package inventory

import (
	"github.com/ganiulis/gilded-rose-go/internal/handler"
	"github.com/ganiulis/gilded-rose-go/pkg/model"
)

func UpdateQuality(items []*model.Item) {
	handlers := []handler.ItemHandler{
		&handler.BackstagePassesHandler{},
		&handler.SulfurasHandler{},
		&handler.AgedBrieHandler{},
		&handler.DefaultHandler{},
	}

	for i := 0; i < len(handlers)-1; i++ {
		nextHandler := handlers[i+1]

		handlers[i].SetNext(nextHandler)
	}

	for i := 0; i < len(items); i++ {
		handlers[0].Update(items[i])
	}
}
