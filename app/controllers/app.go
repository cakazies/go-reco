package controllers

import (
	"go-reco/app/configs"
	"go-reco/app/models"
	"net/http"
	"strconv"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func init() {
	configs.StartViper()
	configs.Connect()
}

func (c App) Health() revel.Result {
	return c.RenderJSON("this API already consume")
}

func (c App) GetStudents() revel.Result {
	stud := models.StudentModels{}

	data, err := stud.GetStudents()
	if err != nil {
		revel.AppLog.Debug("error get student", err)
		result := response(err, "error get data student", "fail")
		c.Response.ContentType = "application/json"
		c.Response.SetStatus(http.StatusInternalServerError)
		return c.RenderJSON(result)
	}

	c.Response.SetStatus(http.StatusOK)
	c.Response.ContentType = "application/json"
	result := response(data, "get data successfully", "success")
	return c.RenderJSON(result)
}

func (c App) InsertStudent() revel.Result {
	data := models.Student{}
	c.Params.BindJSON(&data)

	err := data.Insert()
	if err != nil {
		revel.AppLog.Debug("error insert student", err)
		result := response(err, "error get student", "fail")
		c.Response.ContentType = "application/json"
		c.Response.SetStatus(http.StatusInternalServerError)
		return c.RenderJSON(result)
	}

	c.Response.SetStatus(http.StatusCreated)
	c.Response.ContentType = "application/json"
	result := response(data, "insert data successfull", "success")
	return c.RenderJSON(result)
}

func (c App) GetStudent() revel.Result {
	stud := models.Student{}
	id, err := strconv.Atoi(c.Params.Get("id"))
	if err != nil {
		result := response("param id doesn't exist", "error get student", "fail")
		c.Response.SetStatus(http.StatusNotFound)
		c.Response.ContentType = "application/json"
		return c.RenderJSON(result)
	}
	stud.ID = uint(id)

	data, err := stud.GetStudent()
	if err != nil {
		if err.Error() == "record not found" {
			result := response(err.Error(), "data not found", "fail")
			c.Response.ContentType = "application/json"
			c.Response.SetStatus(http.StatusNotFound)
			return c.RenderJSON(result)
		} else {
			revel.AppLog.Debug("error insert student", err)
			result := response(err.Error(), "error get student", "fail")
			c.Response.ContentType = "application/json"
			c.Response.SetStatus(http.StatusInternalServerError)
			return c.RenderJSON(result)
		}
	}

	c.Response.SetStatus(http.StatusFound)
	c.Response.ContentType = "application/json"
	result := response(data, "get data successfull", "success")
	return c.RenderJSON(result)
}

func (c App) UpdateStudent() revel.Result {
	id, err := strconv.Atoi(c.Params.Get("id"))
	data := models.Student{}
	if err != nil {
		result := response("param id doesn't exist", "error get student", "fail")
		c.Response.SetStatus(http.StatusNotFound)
		c.Response.ContentType = "application/json"
		return c.RenderJSON(result)
	}
	c.Params.BindJSON(&data)
	data.ID = uint(id)

	revel.AppLog.Debug("print student", data)
	err = data.UpdateStudent()
	if err != nil {
		revel.AppLog.Debug("error insert student", err)
		result := response(err, "error get student", "fail")
		c.Response.ContentType = "application/json"
		c.Response.SetStatus(http.StatusInternalServerError)
		return c.RenderJSON(result)
	}

	c.Response.SetStatus(http.StatusCreated)
	c.Response.ContentType = "application/json"
	result := response(data, "update data successfull", "success")
	return c.RenderJSON(result)
}

func (c App) DeleteStudent() revel.Result {
	id, err := strconv.Atoi(c.Params.Get("id"))
	data := models.Student{}
	if err != nil {
		result := response("param id doesn't exist", "error id", "fail")
		c.Response.SetStatus(http.StatusNotFound)
		c.Response.ContentType = "application/json"
		return c.RenderJSON(result)
	}

	data.ID = uint(id)
	err = data.DeleteStudent()
	if err != nil {
		if err.Error() == "record not found" {
			result := response(err.Error(), "data not found", "fail")
			c.Response.ContentType = "application/json"
			c.Response.SetStatus(http.StatusNotFound)
			return c.RenderJSON(result)
		} else {
			revel.AppLog.Debug("error insert student", err)
			result := response(err.Error(), "error get student", "fail")
			c.Response.ContentType = "application/json"
			c.Response.SetStatus(http.StatusInternalServerError)
			return c.RenderJSON(result)
		}
	}

	c.Response.SetStatus(http.StatusCreated)
	c.Response.ContentType = "application/json"
	result := response(data, "update data successfull", "success")
	return c.RenderJSON(result)
}
