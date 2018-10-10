package towngen

var (
	tradeGoods = map[string]GoodsList{
		"ores": GoodsList{
			"ores",
			map[string]int{
				"iron":   15,
				"copper": 20,
				"bronze": 5,
				"gold":   1,
				"silver": 2,
			},
		},
		"wood": GoodsList{
			"wood",
			map[string]int{
				"hardwood": 2,
				"softwood": 1,
			},
		},
		"livestock": GoodsList{
			"livestock",
			map[string]int{
				"cattle":   10,
				"sheep":    8,
				"chickens": 4,
			},
		},
		"hunted": GoodsList{
			"hunted",
			map[string]int{
				"furs":     10,
				"leathers": 5,
			},
		},
		"agriculture": GoodsList{
			"agriculture",
			map[string]int{
				"wheat":        8,
				"corn":         5,
				"potatoes":     5,
				"rye":          4,
				"rice":         4,
				"tea":          3,
				"tobacco":      3,
				"cotton":       4,
				"apples":       2,
				"exotic fruit": 1,
			},
		},
		"luxury": GoodsList{
			"luxury",
			map[string]int{
				"spices":     10,
				"silks":      7,
				"jewels":     3,
				"ivory":      2,
				"fragrances": 6,
				"dyes":       4,
				"carpets":    2,
				"tapestries": 2,
			},
		},
		"quarried": GoodsList{
			"quarried",
			map[string]int{
				"marble":    6,
				"granite":   10,
				"sandstone": 2,
			},
		},
		"seaproducts": GoodsList{
			"seaproducts",
			map[string]int{
				"fish":    12,
				"crab":    3,
				"lobster": 2,
				"oil":     5,
				"pearls":  1,
			},
		},
	}
)

// GoodsList is a structured list of goods
type GoodsList struct {
	Type  string
	Goods map[string]int
}
