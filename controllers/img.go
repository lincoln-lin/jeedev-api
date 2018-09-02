package controllers

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"jeedev-api/myredis"
	"math/rand"
	"strconv"
)

// AppController operations for App
type ImgController struct {
	beego.Controller
}




// img ...
// @Title get img
// @Description get img by sid
// @Param	sid		path 	string	true		"sid"
// @Success 200 {string}
// @Failure 403 :sid is empty
// @router /:sid [get]
func (c *ImgController) Code() {
	sid := c.Ctx.Input.Param(":sid")

	c.Ctx.Output.ContentType("jpg")
	id := strconv.Itoa(rand.Intn(3))
	imagePath := "upload/" + id + ".jpg"
	file, _ := ioutil.ReadFile(imagePath)

	rs := myredis.Conn()
	key := "img"+id;
	code, _ := redis.String(rs.Do("GET", key))
	skey := "code"+sid
	rs.Do("SET", skey,  code)
	c.Ctx.Output.Body(file)
}


// @router /set [get]
func (c *ImgController) Set() {
	key0 := "img0";
	value0 := "abcd";

	key1 := "img1";
	value1 := "efgh";

	key2 := "img2";
	value2 := "lmno";

	r1 := myredis.SetString(key0,value0,"3600")
	r2 := myredis.SetString(key1,value1,"3600")
	r3 := myredis.SetString(key2,value2,"3600")

	if r1 && r2 && r3 {
		c.Data["json"] = "sucess"
	}else{
		c.Data["json"] = "err"
	}

	c.ServeJSON()
}

// @router /get/:key [get]
func (c *ImgController) Get() {
	key := c.Ctx.Input.Param(":key")
	n := myredis.GetString(key)
	c.Data["json"] = n
	c.ServeJSON()
}

// @router /check/:sid/:code [get]
func (c *ImgController) Check() {
	sid := c.Ctx.Input.Param(":sid")
	code := c.Ctx.Input.Param(":code")
	key := "code"+sid
	checkCode := myredis.GetString(key)
	if(code == checkCode){
		c.Data["json"] = "success"
	}else{
		c.Data["json"] = "error"

	}
	c.ServeJSON()
}