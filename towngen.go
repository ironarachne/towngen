package towngen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/utility"
)

// Town is a town
type Town struct {
	Name       string            `json:"townName"`
	Population int               `json:"population"`
	Category   TownCategory      `json:"category"`
	Region     Region            `json:"region"`
	Mayor      chargen.Character `json:"mayor"`
	Exports    map[string]int    `json:"exports"`
	Imports    map[string]int    `json:"imports"`
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

func generateGoodsMap(minAmount int, maxAmount int, region Region) (map[string]int, map[string]int) {
	var importGoods = make(map[string]int)
	var exportGoods = make(map[string]int)
	var good string

	exports, imports := getGoodsForRegion(region)

	rangeAmount := maxAmount - minAmount

	if rangeAmount == 0 {
		rangeAmount = 1
	}

	for i := 0; i < rand.Intn(rangeAmount)+minAmount; i++ {
		good = utility.RandomItemFromThresholdMap(imports)
		importGoods[good] = rand.Intn(maxAmount*6) + 1
	}

	for i := 0; i < rand.Intn(rangeAmount)+minAmount; i++ {
		good = utility.RandomItemFromThresholdMap(exports)
		exportGoods[good] = rand.Intn(maxAmount*6) + 1
	}

	return exportGoods, importGoods
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

func generateRandomRegion() Region {
	var regionNames []string

	for _, region := range regions {
		regionNames = append(regionNames, region.Name)
	}

	return regions[utility.RandomItem(regionNames)]
}

func generateTownName() string {
	townNamePattern := townNames["general"]

	prefix := utility.RandomItem(townNamePattern.Prefixes)
	suffix := utility.RandomItem(townNamePattern.Suffixes)

	return strings.Title(prefix + suffix)
}

// GenerateTown generates a random town
func GenerateTown(category string, region string) Town {
	town := Town{}
	town.Mayor = generateMayor()
	town.Name = generateTownName()
	if category == "random" {
		town.Category = generateRandomCategory()
	} else {
		town.Category = townCategories[category]
	}
	if region == "random" {
		town.Region = generateRandomRegion()
	} else {
		town.Region = regions[region]
	}

	town.Exports, town.Imports = generateGoodsMap(town.Category.MinExports, town.Category.MaxExports, town.Region)
	town.Population = generateRandomPopulation(town.Category)

	return town
}
