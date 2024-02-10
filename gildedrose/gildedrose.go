package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		switch items[i].Name {
		case "Backstage passes to a TAFKAL80ETC concert":
			updateBackstagePasses(items[i])
		case "Sulfuras, Hand of Ragnaros":
			updateSulfuras(items[i])
		case "Aged Brie":
			updateAgedBrie(items[i])
		default:
			update(items[i])
		}
	}
}

func updateBackstagePasses(item *Item) {
	item.SellIn--

	if item.Quality < 50 {
		item.Quality++
	}

	if item.Quality < 50 {
		if item.SellIn < 10 {
			item.Quality++
		}

		if item.SellIn < 5 {
			item.Quality++
		}
	}

	if item.SellIn < 0 {
		item.Quality = 0
	}
}

func updateSulfuras(item *Item) {

}

func updateAgedBrie(item *Item) {
	item.SellIn--

	item.Quality++

	if item.SellIn < 0 {
		item.Quality++
	}

	if item.Quality > 50 {
		item.Quality = 50
	}
}

func update(item *Item) {
	item.SellIn--

	if item.Quality > 0 {
		item.Quality--
	}

	if item.Quality > 0 && item.SellIn < 0 {
		item.Quality--
	}
}
