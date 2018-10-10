package towngen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/utility"
)

// Town is a town
type Town struct {
	Name       string
	Population int
	Category   TownCategory
	Mayor      chargen.Character
	Exports    map[string]int
	Imports    map[string]int
}

// TownCategory is a type of town
type TownCategory struct {
	Name       string
	MinSize    int
	MaxSize    int
	MinExports int
	MaxExports int
	MinImports int
	MaxImports int
}

func generateGoodsMap(minAmount int, maxAmount int) map[string]int {
	var goods = make(map[string]int)
	var good string
	rangeAmount := maxAmount - minAmount
	if rangeAmount == 0 {
		rangeAmount = 1
	}

	for i := 0; i < rand.Intn(rangeAmount)+minAmount; i++ {
		good = utility.RandomItem(tradeGoods)
		goods[good] = rand.Intn(maxAmount*6) + 1
	}

	return goods
}

func generateMayor() chargen.Character {
	mayor := chargen.GenerateCharacter()

	return mayor
}

func generateRandomCategory() TownCategory {
	categoryName := utility.RandomItemFromThresholdMap(townCategoryOptions)

	category := townCategories[categoryName]

	return category
}

func generateRandomPopulation(category TownCategory) int {
	sizeIncrement := category.MaxSize - category.MinSize

	return rand.Intn(sizeIncrement) + category.MinSize
}

func generateTownName() string {
	townNamePattern := townNames["general"]

	prefix := utility.RandomItem(townNamePattern.Prefixes)
	suffix := utility.RandomItem(townNamePattern.Suffixes)

	return strings.Title(prefix + suffix)
}

// GenerateTown generates a random town
func GenerateTown() Town {
	town := Town{}
	town.Mayor = generateMayor()
	town.Name = generateTownName()
	town.Category = generateRandomCategory()
	town.Exports = generateGoodsMap(town.Category.MinExports, town.Category.MaxExports)
	town.Imports = generateGoodsMap(town.Category.MinImports, town.Category.MaxExports)
	town.Population = generateRandomPopulation(town.Category)

	return town
}
