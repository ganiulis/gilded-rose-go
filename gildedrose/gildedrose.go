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
	if item.Name == "Aged Brie" {
		if item.Quality < 50 {
			item.Quality++
		}
	}

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
	}

	if item.Name != "Aged Brie" &&
		item.Name != "Backstage passes to a TAFKAL80ETC concert" &&
		item.Name != "Sulfuras, Hand of Ragnaros" &&
		item.Quality > 0 {
		item.Quality--
	}

	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn--
	}

	if item.SellIn >= 0 {
		return
	}

	if item.Name == "Aged Brie" {
		if item.Quality > 49 {
			return
		}

		item.Quality++

		return
	}

	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		item.Quality = 0
	}

	if item.Name != "Backstage passes to a TAFKAL80ETC concert" &&
		item.Name != "Sulfuras, Hand of Ragnaros" &&
		item.Quality > 0 {
		item.Quality--
	}
}
