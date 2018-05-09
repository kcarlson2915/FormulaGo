package main

import "github.com/kcarlson2915/FormulaGo/ergast"

func main() {
	c := ergast.Client{ergast.DefaultURL}
	c.GetSeason(2018)
}
