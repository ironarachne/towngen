package towngen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/climategen"
	"github.com/ironarachne/utility"
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
	categoryName := utility.RandomItemFromThresholdMap(townCategoryOptions)

	category := townCategories[categoryName]

	return category
}

func generateRandomExports(climate climategen.Climate, category TownCategory) map[string]int {
	exports := map[string]int{}

	numberOfExports := rand.Intn(category.MaxExports+1-category.MinExports) + category.MinExports
	exportAmount := 0
	newExport := ""

	for i := 0; i < numberOfExports; i++ {
		newExport = utility.RandomItem(climate.Resources)
		exportAmount = rand.Intn(3) + 1
		exports[newExport] = exportAmount
	}

	return exports
}

func generateRandomImports(climate climategen.Climate, category TownCategory) map[string]int {
	imports := map[string]int{}

	numberOfImports := rand.Intn(category.MaxImports+1-category.MinImports) + category.MinImports
	importAmount := 0
	newImport := ""

	for i := 0; i < numberOfImports; i++ {
		newImport = utility.RandomItem(climate.Needs)
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

	prefix := utility.RandomItem(townNamePattern.Prefixes)
	suffix := utility.RandomItem(townNamePattern.Suffixes)

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

	town.Exports = generateRandomExports(town.Climate, town.Category)
	town.Imports = generateRandomImports(town.Climate, town.Category)
	town.Population = generateRandomPopulation(town.Category)

	return town
}
