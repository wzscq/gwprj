package project

import (
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
)

type PageSize struct {
	W float64 `json:"w,omitempty"`
	H float64 `json:"h,omitempty"`
}

type ElementColor struct {
	R uint8 `json:"r,omitempty"`
	G uint8 `json:"g,omitempty"`
	B uint8 `json:"b,omitempty"`
}

type TextFont struct {
	Size float64 `json:"size,omitempty"`
	Style string `json:"style,omitempty"`
	Family string `json:"family,omitempty"`
}

type PageElement struct {
	Type string `json:"type,omitempty"`
	Font *TextFont `json:"font,omitempty"`
	Color *ElementColor `json:"color,omitempty"`
	Rect [4]float64 `json:"rect,omitempty"`
	Content string `json:"content,omitempty"`
	Width float64 `json:"width,omitempty"`
	LineHeight float64  `json:"lineHeight,omitempty"`
	WordSpace float64 `json:"wordSpace,omitempty"`
	MaxContentLength int `json:"maxContentLength,omitempty"`
}

type TempletePage struct {
	Elements []PageElement `json:"elements,omitempty"`
}

type Templete struct {
	PageSize PageSize `json:"pageSize,omitempty"`
	Pages []TempletePage `json:"pages,omitempty"`
}

const (
	ElementTypeText = "text"
	ElementTypeLine = "line"
)

const (
	ElementAlignLeft = "left"
	ElementAlignCenter = "center"
	ElementAlignRight = "right"
)

//从json格式的配置文件中加载模板
func LoadTempleteFromJsonFile(tmpFile string)(*Templete,error){
	file, err := os.Open(tmpFile)
	if err != nil {
		log.Println("Error opening file: %!v", err)
		return nil,err
	}
	defer file.Close()

	// 读取文件内容
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading the file: %!v", err)
		return nil,err
	}

	// 反序列化 JSON 到结构体
	var tmp Templete
	err = json.Unmarshal(bytes, &tmp)
	if err != nil {
		log.Println("Error unmarshalling JSON: %!v", err)
		return nil,err
	}

	return &tmp,nil
}