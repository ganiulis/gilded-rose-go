package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	handlers := []ItemHandler{
		&BackstagePassesHandler{},
		&SulfurasHandler{},
		&AgedBrieHandler{},
		&GenericHandler{},
	}

	for i := 0; i < len(handlers)-1; i++ {
		nextHandler := handlers[i+1]

		handlers[i].setNext(nextHandler)
	}

	for i := 0; i < len(items); i++ {
		handlers[0].update(items[i])
	}
}

type ItemHandler interface {
	update(item *Item)
	setNext(ItemHandler)
}

type GenericHandler struct {
	next ItemHandler
}

func (u *GenericHandler) update(item *Item) {
	item.SellIn--

	if item.Quality > 0 {
		item.Quality--
	}

	if item.Quality > 0 && item.SellIn < 0 {
		item.Quality--
	}
}

func (u *GenericHandler) setNext(next ItemHandler) {
	u.next = next
}

type SulfurasHandler struct {
	next ItemHandler
}

func (u *SulfurasHandler) update(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		u.next.update(item)
	}
}

func (u *SulfurasHandler) setNext(next ItemHandler) {
	u.next = next
}

type AgedBrieHandler struct {
	next ItemHandler
}

func (u *AgedBrieHandler) update(item *Item) {
	if item.Name != "Aged Brie" {
		u.next.update(item)

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

func (u *AgedBrieHandler) setNext(next ItemHandler) {
	u.next = next
}

type BackstagePassesHandler struct {
	next ItemHandler
}

func (u *BackstagePassesHandler) update(item *Item) {
	if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		u.next.update(item)

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

func (u *BackstagePassesHandler) setNext(next ItemHandler) {
	u.next = next
}
