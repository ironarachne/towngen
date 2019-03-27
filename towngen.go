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
	Exports    map[string]int     `json:"exports"`
	Imports    map[string]int     `json:"imports"`
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

func (town Town) generateRandomExports() map[string]int {
	exports := map[string]int{}
	possibleExports := GetAllTradeGoods(town.Climate.Resources)

	numberOfExports := rand.Intn(town.Category.MaxExports+1-town.Category.MinExports) + town.Category.MinExports
	exportAmount := 0
	newExport := ""

	for i := 0; i < numberOfExports; i++ {
		newExport = random.Item(possibleExports)
		exportAmount = rand.Intn(3) + 1
		exports[newExport] = exportAmount
	}

	return exports
}

func (town Town) generateRandomImports() map[string]int {
	imports := map[string]int{}
	possibleImports := GetAllTradeGoods(town.Climate.Needs)

	numberOfImports := rand.Intn(town.Category.MaxImports+1-town.Category.MinImports) + town.Category.MinImports
	importAmount := 0
	newImport := ""

	for i := 0; i < numberOfImports; i++ {
		newImport = random.Item(possibleImports)
		importAmount = rand.Intn(3) + 1
		imports[newImport] = importAmount
	}

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
