package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
	"github.com/astaxie/beego/logs"
	"strings"
	"github.com/golangpkg/go-stat-reporter/models"
)

type StatPageController struct {
	beego.Controller
}

type StatTableApiJson struct {
	Draw            int           `json:"draw"`
	RecordsTotal    int           `json:"recordsTotal"`
	RecordsFiltered int           `json:"recordsFiltered"`
	Data            []interface{} `json:"data"`
}

//html页面
func (c *StatPageController) PageHtml() {
	//获得id
	tableId := c.GetString("pageId", "")
	page, err := models.GetPage(tableId)
	logs.Info("page: ", page)
	if err == nil {
		c.Data["Page"] = page
	} else {
		c.Data["Page"] = &models.XMLPage{}
	}
	c.TplName = "reporter/page.html"
}

//Ajax分页
func (c *StatPageController) TableApi() {
	defer c.ServeJSON()

	//获得id
	tableId := c.GetString("tableId", "")
	startParam := c.GetString("start", "")
	lengthParam := c.GetString("length", "")
	table, err := models.GetTable(tableId)
	logs.Info("table: ", table)
	if err == nil {
		c.Data["Table"] = table
	} else {
		c.Data["Table"] = &models.XMLDataTable{}
	}
	//获得全部 request参数。
	paramsTmp := c.Ctx.Request.Form
	logs.Info("startParam :", startParam, ", lengthParam:", lengthParam)
	orderByColumn := []string{}
	orderByDir := []string{}

	for key, val := range paramsTmp {
		//循环找到排序字段，然后放到数组里面。
		if strings.Index(key, "order[") == 0 && strings.Index(key, "[column]") > 0 {
			logs.Info("get order column :", key, val)
			orderByColumn = append(orderByColumn, val[0])
		}
		//排序dir。
		if strings.Index(key, "order[") == 0 && strings.Index(key, "[dir]") > 0 {
			logs.Info("get order dir :", key, val)
			orderByDir = append(orderByDir, val[0])
		}
	}
	logs.Info("orderByColumn : ", orderByColumn)
	logs.Info("orderByDir : ", orderByDir)
	logs.Info("table : %v ", table)

	//#######################组装order by sql。#######################
	orderBySql := ""
	if len(orderByColumn) != 0 && len(orderByDir) != 0 {
		orderBySql = " ORDER BY "
		idx := 0
		for _, val := range orderByColumn {
			//找到排序字段和dir。
			tmpIndex, _ := strconv.Atoi(val)
			logs.Info("get val :", tmpIndex)
			colTmp := table.ColumnArray[tmpIndex]
			dirTmp := orderByDir[idx]

			logs.Info("key:", tmpIndex, "colTmp :", colTmp, ", dirTmp :", dirTmp)
			//排序sql。
			tmpOrderSql := fmt.Sprintf(" convert(`%s`, decimal) %s", colTmp, dirTmp)

			//增加order sql 排序。
			if idx != 0 {
				orderBySql += " , " + tmpOrderSql
			} else {
				orderBySql += tmpOrderSql
			}
			idx += 1
		}
	}

	sql := fmt.Sprintf(" SELECT * FROM %s %s LIMIT %s,%s ", table.Table, orderBySql, startParam, lengthParam)
	sqlCount := fmt.Sprintf(" SELECT COUNT(1) as num FROM %s ", table.Table)
	logs.Info("select sql :", sql)
	logs.Info("count sql :", sqlCount)
	//stock_web_list = self.db.query(sql)

	var maps []orm.Params
	var list []interface{}
	o := orm.NewOrm()
	dataListTmp, err1 := o.Raw(sql).Values(&maps)
	if err1 == nil && dataListTmp > 0 {
		fmt.Println(maps[0]) // slene
		for _, val := range maps {
			//logs.Info("key:", key, ", val:", val)
			list = append(list, val)
		}
	}

	//记录count数量。
	numTmp := 0
	dataListTmp2, err2 := o.Raw(sqlCount).Values(&maps)
	if err2 == nil && dataListTmp2 > 0 {
		//fmt.Println(maps[0]["num"]) // slene
		//转换字符串。
		if s, err := strconv.Atoi(maps[0]["num"].(string)); err == nil {
			fmt.Printf("%T, %v", s, s)
			numTmp = s
		}
	}

	//放入data数据：
	tmpJson := &StatTableApiJson{}
	tmpJson.RecordsTotal = numTmp
	tmpJson.RecordsFiltered = numTmp
	tmpJson.Data = list
	c.Data["json"] = tmpJson
}
