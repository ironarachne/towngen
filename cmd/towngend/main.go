package main

import (
	"math/rand"

	"github.com/ironarachne/towngen"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("towngend")
	})

	app.Get("/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			ctx.Writef("error while trying to parse id parameter")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		rand.Seed(id)
		town := towngen.GenerateTown("random", "random")

		ctx.JSON(town)
	})

	app.Run(iris.Addr(":7979"))
}
