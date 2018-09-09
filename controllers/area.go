package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"jeedev-api/lib/myhttp"
	"jeedev-api/models"
	"jeedev-api/mymysql"
	"jeedev-api/units"
	"strconv"
)

//  AreaController operations for Area
type AreaController struct {
	beego.Controller
}
type Respon struct {
	Status int         `json:"status"`
	Mesage string      `json:"mesage"`
	Data   interface{} `json:"data"`
}

// URLMapping ...
func (c *AreaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Area
// @Param	body		body 	models.Area	true		"body for Area content"
// @Success 201 {int} models.Area
// @Failure 403 body is empty
// @router / [post]
func (c *AreaController) Post() {
	var v models.Area
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddArea(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Area by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Area
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AreaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAreaById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Area
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Area
// @Failure 403
// @router / [get]
func (c *AreaController) GetAll() {
	params := map[string]interface{}{}
	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		params["fields"] = v
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		params["limit"] = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		params["offset"] = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		params["sortby"] = v
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		params["order"] = v
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		params["query"] = v
	}
	mymysql.TableName(models.Area{})

	l, err := mymysql.FindAll(params, models.Area{})
	if err != nil {
		c.Data["json"] = err.Error()
	} else {

		c.Data["json"] = Respon{200, "sucess", l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Area
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Area	true		"body for Area content"
// @Success 200 {object} models.Area
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AreaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Area{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateAreaById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Area
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AreaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteArea(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @router /test [get]
func (c *AreaController) Test() {
	m := map[string]interface{}{
		"a": "aa",
		"b": 1,
	}
	url := "http://127.0.0.1:8080/v1/area"
	resp, _ := myhttp.HttpPost(url, m)
	body, _ := resp.Body() //得到的是byte
	data := units.Bytes2Intaface(body)
	fmt.Println(data)
	c.Data["json"] = data

	c.ServeJSON()
}
