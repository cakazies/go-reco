package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Beer struct {
	ID      int
	Name    string
	Type    string
	Brewery string
}

func (c App) Index() revel.Result {

	return c.Render()
}

func (c App) Try() revel.Result {
	var beers = []Beer{
		Beer{1, "La Trappe Quadrupel Oak Aged", "Ale", "Bierbrouwerij De Koningshoeven"},
		Beer{2, "Cocoa Psycho", "Porter", "BrewDog"},
		Beer{3, "American Dream", "Lager", "Mikkeller"},
	}
	// try, err := json.Marshal(data)
	// if err != nil {
	// 	log.Println("error")
	// 	return c.RenderError(err)
	// }

	return c.RenderJSON(beers)
}
