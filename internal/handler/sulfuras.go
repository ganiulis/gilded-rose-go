package handler

import "github.com/ganiulis/gilded-rose-go/pkg/model"

type SulfurasHandler struct {
	next ItemHandler
}

func (u *SulfurasHandler) Update(item *model.Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		u.next.Update(item)
	}
}

func (u *SulfurasHandler) SetNext(next ItemHandler) {
	u.next = next
}
