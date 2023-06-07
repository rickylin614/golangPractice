package controllers

import (
	"strconv"

	"github.com/revel/revel"
)

type UserCtrl struct {
	*revel.Controller
}

func (c UserCtrl) Index() revel.Result {
	var datas = "Hello, world! 你好"

	return c.Render(datas)
}

func (c UserCtrl) GetOne(id int) revel.Result {
	m := map[string]any{
		"datas": "userId:" + strconv.Itoa(id),
	}
	return c.RenderJSON(m)
}

func (c UserCtrl) Create() revel.Result {
	return c.Render()
}

func (c UserCtrl) Update() revel.Result {
	return c.Render()
}

func (c UserCtrl) Delete() revel.Result {
	return c.Render()
}
