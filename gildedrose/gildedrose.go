package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality--
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality++

				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" && items[i].Quality < 50 {
					if items[i].SellIn < 11 {
						items[i].Quality++
					}
					if items[i].SellIn < 6 {
						items[i].Quality++
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn--
		}

		if items[i].SellIn >= 0 {
			continue
		}

		if items[i].Name == "Aged Brie" {
			if items[i].Quality > 49 {
				continue
			}

			items[i].Quality++
			continue
		}

		if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality--
				}
			}
		} else {
			items[i].Quality = 0
		}

	}

}
