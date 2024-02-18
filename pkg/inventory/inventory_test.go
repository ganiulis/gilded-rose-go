package inventory

import (
	"testing"

	"github.com/ganiulis/gilded-rose-go/pkg/model"
)

func TestUpdate(t *testing.T) {
	var items = []*model.Item{
		{"fixme", 0, 0},
	}

	Update(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}
