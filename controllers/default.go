package controllers

import (
	"Beego001/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller //匿名字段：
}
type FuncController struct {
	beego.Controller
}
/**
方法的重写：如果一个结构体包含某个方法A，在其匿名中也拥有同名方法A。
则结构外部结构体的方法A是对匿名字段结构体方法A的
 */
func (c *MainController) Get() {

	//1.获取请求数据
	userName := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("psd")
	//2.使用固定数据进行数据校验
	//admin 228947
	if userName != "pengwei" || password != "228947" {
		//代表数据错误
		c.Ctx.ResponseWriter.Write([]byte("对不起数据校验错误"))
		return
	}
	//数据校验正确
	c.Ctx.ResponseWriter.Write([]byte("数据校验成功"))
	return
	//1.获取name、age、sex
	//2.做数据对比
	//3.根据对比结果进行判断处理
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "2289476116@qq.com"
	c.TplName = "index.tpl"
}
/**
该方法用于处理POST请求
 */
//func (c*MainController)Post(){
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex := c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,sex,age)
//	//2.进行数据校验
//	if name != "pengwei" && age != "20"{
//		c.Ctx.WriteString("数据校验失败")
//		return
//	}
//	c.Ctx.WriteString("数据校验成功")
//}
//该方法用于处理post类型请求
func (c*MainController)Post(){
var person models.Person
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil{
		c.Ctx.WriteString("数据解析失败")
	}
	fmt.Println("姓名：",person.Name)
	fmt.Println("年龄：",person.Age)
	fmt.Println("性别：",person.Sex)
	c.Ctx.WriteString("数据解析成功！")

}

func (c*FuncController)Post(){
	var personto models.Personto
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误")
		return
	}
	err = json.Unmarshal(dataBytes,&personto)
	if err != nil{
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("姓名：",personto.Name)
	fmt.Println("生日：",personto.Birthday)
	fmt.Println("地址：",personto.Address)
	fmt.Println("别名:",personto.Nick)
	c.Ctx.WriteString("数据解析成功！")
}