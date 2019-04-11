package towngen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/climategen"
	"github.com/ironarachne/random"
)

// Town is a town
type Town struct {
	Name       string             `json:"townName"`
	Population int                `json:"population"`
	Category   TownCategory       `json:"category"`
	Climate    climategen.Climate `json:"climate"`
	Mayor      chargen.Character  `json:"mayor"`
	Exports    []TradeGood        `json:"exports"`
	Imports    []TradeGood        `json:"imports"`
}

// TownCategory is a type of town
type TownCategory struct {
	Name       string `json:"name"`
	MinSize    int    `json:"minSize"`
	MaxSize    int    `json:"maxSize"`
	MinExports int    `json:"minExports"`
	MaxExports int    `json:"maxExports"`
	MinImports int    `json:"minImports"`
	MaxImports int    `json:"maxImports"`
}

func generateMayor() chargen.Character {
	mayor := chargen.GenerateCharacter()

	return mayor
}

func generateRandomCategory() TownCategory {
	categoryName := random.ItemFromThresholdMap(townCategoryOptions)

	category := townCategories[categoryName]

	return category
}

func (town Town) generateRandomExports() []TradeGood {
	var exports []TradeGood

	exports = generateTradeGoods(town.Category.MinExports, town.Category.MaxExports, town.Climate.Resources)

	return exports
}

func (town Town) generateRandomImports() []TradeGood {
	var imports []TradeGood

	foreignClimate := climategen.GetForeignClimate(town.Climate)

	imports = generateTradeGoods(town.Category.MinImports, town.Category.MaxImports, foreignClimate.Resources)

	return imports
}

func generateRandomPopulation(category TownCategory) int {
	sizeIncrement := category.MaxSize - category.MinSize

	return rand.Intn(sizeIncrement) + category.MinSize
}

func generateTownName() string {
	townNamePattern := townNames["general"]

	prefix := random.Item(townNamePattern.Prefixes)
	suffix := random.Item(townNamePattern.Suffixes)

	return strings.Title(prefix + suffix)
}

// GenerateTown generates a random town
func GenerateTown(category string, climate string) Town {
	town := Town{}

	town.Mayor = generateMayor()
	town.Name = generateTownName()

	if category == "random" {
		town.Category = generateRandomCategory()
	} else {
		town.Category = townCategories[category]
	}
	if climate == "random" {
		town.Climate = climategen.Generate()
	} else {
		town.Climate = climategen.GetClimate(climate)
	}

	town.Exports = town.generateRandomExports()
	town.Imports = town.generateRandomImports()
	town.Population = generateRandomPopulation(town.Category)

	return town
}
