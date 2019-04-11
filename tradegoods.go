package towngen

import (
	"math/rand"

	"github.com/ironarachne/climategen"
	"github.com/ironarachne/random"
)

// TradeGood is a trade good entry
type TradeGood struct {
	Name    string
	Quality string
	Amount  int
}

func generateTradeGoods(min int, max int, resources []climategen.Resource) []TradeGood {
	var good TradeGood

	goods := []TradeGood{}

	possibleGoods := GetAllTradeGoods(resources)

	numberOfGoods := rand.Intn(max+1-min) + min
	amount := 0
	newItem := ""

	for i := 0; i < numberOfGoods; i++ {
		good = TradeGood{}
		newItem = random.Item(possibleGoods)
		amount = rand.Intn(3) + 1
		good.Name = newItem
		good.Amount = amount
		good.Quality = randomQuality()
		goods = append(goods, good)
	}

	return goods
}

// GetAllTradeGoods converts a list of resources into a list of trade goods
func GetAllTradeGoods(resources []climategen.Resource) []string {
	goods := []string{}

	for _, resource := range resources {
		goods = append(goods, resource.Name)
	}

	return goods
}

func randomQuality() string {
	qualities := map[string]int{
		"exceptional":  1,
		"fine":         2,
		"":             11,
		"questionable": 2,
		"pathetic":     1,
	}

	return random.ItemFromThresholdMap(qualities)
}
