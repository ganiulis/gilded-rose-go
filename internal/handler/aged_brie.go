package handler

import "github.com/ganiulis/gilded-rose-go/pkg/model"

type AgedBrieHandler struct {
	next ItemHandler
}

func (u *AgedBrieHandler) Update(item *model.Item) {
	if item.Name != "Aged Brie" {
		u.next.Update(item)

		return
	}

	item.SellIn--
	item.Quality++

	if item.SellIn < 0 {
		item.Quality++
	}

	if item.Quality > 50 {
		item.Quality = 50
	}
}

func (u *AgedBrieHandler) SetNext(next ItemHandler) {
	u.next = next
}
