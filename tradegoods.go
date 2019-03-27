package towngen

import (
	"github.com/ironarachne/climategen"
)

// GetAllTradeGoods converts a list of resources into a list of trade goods
func GetAllTradeGoods(resources []climategen.Resource) []string {
	var goodsName string

	goods := []string{}

	for _, resource := range resources {
		for _, t := range resource.Types {
			if t == "hide" {
				goodsName = resource.Name + " hides"
			} else if t == "meat" {
				goodsName = resource.Name + " meat"
			} else if t == "feathers" {
				goodsName = resource.Name + " feathers"
			} else if t == "eggs" {
				goodsName = resource.Name + " eggs"
			} else if t == "fruit" || t == "grain" {
				goodsName = resource.Name
			} else {
				goodsName = resource.Name
			}

			if !inSlice(goodsName, goods) {
				goods = append(goods, goodsName)
			}
		}
	}

	return goods
}
