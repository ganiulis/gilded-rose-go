package handler

import "github.com/ganiulis/gilded-rose-go/pkg/model"

type ItemHandler interface {
	Update(item *model.Item)
	SetNext(ItemHandler)
}
