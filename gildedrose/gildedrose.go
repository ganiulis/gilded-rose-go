package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		updateItem(items[i])
	}
}

func updateItem(item *Item) {
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality < 50 {
			item.Quality++
		}

		if item.Quality < 50 {
			if item.SellIn < 11 {
				item.Quality++
			}

			if item.SellIn < 6 {
				item.Quality++
			}
		}

		item.SellIn--

		if item.SellIn < 0 {
			item.Quality = 0
		}

		return
	}

	if item.Name == "Sulfuras, Hand of Ragnaros" {
		return
	}

	if item.Name == "Aged Brie" {
		item.SellIn--
		item.Quality++

		if item.SellIn < 0 {
			item.Quality++
		}

		if item.Quality > 50 {
			item.Quality = 50
		}

		return
	}

	item.SellIn--

	if item.Quality > 0 {
		item.Quality--
	}

	if item.Quality > 0 &&
		item.SellIn < 0 {
		item.Quality--
	}
}
