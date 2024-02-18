package handler

import "github.com/ganiulis/gilded-rose-go/pkg/model"

type BackstagePassesHandler struct {
	next ItemHandler
}

func (u *BackstagePassesHandler) Update(item *model.Item) {
	if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		u.next.Update(item)

		return
	}

	item.SellIn--

	if item.SellIn < 0 {
		item.Quality = 0

		return
	}

	item.Quality++

	if item.SellIn < 10 {
		item.Quality++
	}

	if item.SellIn < 5 {
		item.Quality++
	}

	if item.Quality > 50 {
		item.Quality = 50
	}
}

func (u *BackstagePassesHandler) SetNext(next ItemHandler) {
	u.next = next
}
