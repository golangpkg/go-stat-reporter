package models

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"fmt"
	"errors"
)

//Table标签
type XMLDataTable struct {
	XMLName xml.Name `xml:"table"̀`
	Id      string   `xml:"id,attr"̀`
	Name    string   `xml:"name,attr"̀`
	Table   string   `xml:"table,attr"̀`
	Column  string   `xml:"column"̀`
	Label   string   ` xml:"label"̀`
}

//Table标签
type XMLDataChart struct {
	XMLName xml.Name `xml:"table"̀`
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
