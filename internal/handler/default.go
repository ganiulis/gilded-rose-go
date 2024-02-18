package handler

import "github.com/ganiulis/gilded-rose-go/pkg/model"

type DefaultHandler struct {
	next ItemHandler
}

func (u *DefaultHandler) Update(item *model.Item) {
	item.SellIn--

	if item.Quality > 0 {
		item.Quality--
	}

	if item.Quality > 0 && item.SellIn < 0 {
		item.Quality--
	}
}

func (u *DefaultHandler) SetNext(next ItemHandler) {
	u.next = next
}
