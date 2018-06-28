package models

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"fmt"
	"errors"
	"strings"
)

//Table标签
type XMLDataTable struct {
	XMLName     xml.Name `xml:"dataTable"̀`
	Id          string   `xml:"id,attr"̀`
	Name        string   `xml:"name,attr"̀`
	Table       string   `xml:"table,attr"̀`
	Column      string   `xml:"column"̀`
	Label       string   ` xml:"label"̀`
	ColumnArray []string //字段项
	LabelArray  []string //字段名字
}

//Table标签
type XMLDataChart struct {
	XMLName xml.Name `xml:"dataChart"̀`
	Id      string   `xml:"id,attr"̀`
	Name    string   `xml:"name,attr"̀`
	Table   string   `xml:"table,attr"̀`
	Type    string   `xml:"type,attr"̀`
	Column  string   `xml:"column"̀`
	Label   string   ` xml:"label"̀`
}

//Page标签
type XMLPage struct {
	XMLName    xml.Name       `xml:"page"̀`
	Id         string         `xml:"id,attr"̀`
	Name       string         `xml:"name,attr"̀`
	DataTables []XMLDataTable `xml:"dataTable"̀`
	DataCharts []XMLDataChart `xml:"dataChart"̀`
}

//Pages标签
type XMLPages struct {
	XMLName xml.Name  `xml:"pages"̀`
	Pages   []XMLPage `xml:"page"̀`
}

var ConstantXmlPages XMLPages

//参考代码：
// https://tutorialedge.net/golang/parsing-xml-with-golang/
func ReadXMLConfig(file string) {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Open XML : " + file)
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &ConstantXmlPages)
	println(ConstantXmlPages.Pages)

	for _, page := range ConstantXmlPages.Pages {
		println("page:", page.Name)
		for i, table := range page.DataTables {
			println("table:", table.Name)
			//转换 column为数组字段。
			if table.Column != "" {
				tmpStr := strings.Replace(table.Column, "\"", "", -1)
				tmpStr = strings.Replace(tmpStr, " ", "", -1)
				tmpStr = strings.Replace(tmpStr, "\r", "", -1)
				tmpStr = strings.Replace(tmpStr, "\n", "", -1)
				table.ColumnArray = strings.Split(tmpStr, ",")
			}
			//转换 label 为数组字段。
			if table.Label != "" {
				tmpStr := strings.Replace(table.Label, "\"", "", -1)
				tmpStr = strings.Replace(tmpStr, " ", "", -1)
				tmpStr = strings.Replace(tmpStr, "\r", "", -1)
				tmpStr = strings.Replace(tmpStr, "\n", "", -1)
				table.LabelArray = strings.Split(tmpStr, ",")
			}
			//fmt.Printf("table %v", table)
			//在将数据放到数组里面。
			page.DataTables[i] = table
		}
	}

	return
}

//循环pages的全部数据，找到id，返回page。
func GetPage(id string) (page XMLPage, err error) {
	for _, page := range ConstantXmlPages.Pages {
		if page.Id == id {
			return page, nil
		}
	}
	return page, errors.New("no page .")
}

//循环Tables的全部数据，找到id，返回Table。
func GetTable(id string) (table XMLDataTable, err error) {
	for _, page := range ConstantXmlPages.Pages {
		for _, table2 := range page.DataTables {
			if table2.Id == id {
				return table2, nil
			}
		}
	}
	return table, errors.New("no page .")
}
