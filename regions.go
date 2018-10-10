package towngen

var (
	regions = map[string]Region{
		"forest": Region{
			"forest",
			[]string{
				"hunted",
				"wood",
			},
			[]string{
				"ores",
				"quarried",
				"luxury",
			},
		},
		"hills": Region{
			"hills",
			[]string{
				"agriculture",
				"livestock",
				"hunted",
			},
			[]string{
				"wood",
				"luxury",
			},
		},
		"mountains": Region{
			"mountains",
			[]string{
				"ores",
				"quarried",
			},
			[]string{
				"agriculture",
				"livestock",
				"luxury",
			},
		},
		"desert": Region{
			"desert",
			[]string{
				"luxury",
			},
			[]string{
				"agriculture",
				"livestock",
				"seaproducts",
				"ore",
				"quarried",
				"wood",
			},
		},
		"coastal": Region{
			"coastal",
			[]string{
				"seaproducts",
			},
			[]string{
				"wood",
				"hunted",
				"ore",
				"quarried",
				"luxury",
			},
		},
	}
)

// Region is a geographic area
type Region struct {
	Name     string
	Produces []string
	Needs    []string
}

func getGoodsForRegion(region Region) (map[string]int, map[string]int) {
	var exportGoods = make(map[string]int)
	var importGoods = make(map[string]int)
	var importTypes []string
	var exportTypes []string

	for _, goodType := range region.Produces {
		exportTypes = append(exportTypes, goodType)
	}

	for _, goodType := range region.Needs {
		importTypes = append(importTypes, goodType)
	}

	for _, goodType := range exportTypes {
		for key, value := range tradeGoods[goodType].Goods {
			exportGoods[key] = value
		}
	}

	for _, goodType := range importTypes {
		for key, value := range tradeGoods[goodType].Goods {
			importGoods[key] = value
		}
	}

	return exportGoods, importGoods
}
