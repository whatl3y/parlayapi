package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kataras/iris"
	conv "github.com/whatl3y/odds/conversions"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("public/index.html", true)
	})

	app.Get("/{odds:path}", func(ctx iris.Context) {
		odds := ctx.Params().Get("odds")
		origOdds := strings.Split(odds, "/")
		oddsIntAry := make([]int, 0)
		oddsStrAry := make([]string, 0)
		for _, strOdds := range origOdds {
			intOdds, err := strconv.Atoi(strOdds)
			if err == nil {
				oddsIntAry = append(oddsIntAry, intOdds)
				oddsStrAry = append(oddsStrAry, strOdds)
			}
		}
		overallOdds := conv.CalculateOdds(oddsIntAry)

		ctx.JSON(iris.Map{
			"american":          conv.GetAmericanOddsFromOverall(overallOdds),
			"decimal":           fmt.Sprintf("%.2f", overallOdds),
			"fractional":        conv.GetFractionalOddsFromOverallOdds(overallOdds),
			"general_for":       fmt.Sprintf("%s for 1", fmt.Sprintf("%.2f", overallOdds)),
			"general_to":        fmt.Sprintf("%s to 1", fmt.Sprintf("%.2f", overallOdds-float64(1))),
			"odds":              strings.Join(oddsStrAry, ","),
			"wager100_totalwin": strconv.FormatFloat(float64(100)*overallOdds, 'f', 2, 64),
			"wager100_win":      strconv.FormatFloat((float64(100)*overallOdds)-float64(100), 'f', 2, 64),
		})
	})

	// listen and serve on http://0.0.0.0:{port}.
	app.Run(iris.Addr(fmt.Sprintf(":%s", port)))
}
