package controllers

import (
	"go-reco/app/configs"
	"go-reco/app/models"

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

func init() {
	configs.StartViper()
	configs.Connect()
}

func (c App) Index() revel.Result {
	stud := models.StudentModels{}
	stud.GetStudent()
	return c.Render()
}

func (c App) InsertStudent() revel.Result {
	stud := models.StudentModels{}
	stud.GetStudent()
	return c.RenderJSON("opi")
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
