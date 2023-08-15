package project

import (
	"github.com/signintech/gopdf"
	"log"
	"errors"
)

//根据templete定义和输入数据生成pdf文件
func CreatePdf(tmp *Templete,data map[string]interface{})(*gopdf.GoPdf){
	//创建pdf对象
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	//加载字体
	// 设置字体
	err := pdf.AddTTFFont("msyh", "./font/msyh.ttf")
	if err != nil {
		log.Fatalf("Error loading font: %!!v", err)
		return nil
	}

	err = pdf.SetFont("msyh", "", 12)
	if err != nil {
		log.Fatalf("Error setting font: %!!v", err)
		return nil
	}

	//根据模板定义创建pdf页面
	for _,page:=range tmp.Pages {
		err:=CreatePdfPage(&pdf,&page,data)
		if err!=nil {
			log.Fatalf("Error create pdf page : %!!v", err)
			return nil
		}
	}
	return &pdf
}

func CreatePdfPage(pdf *gopdf.GoPdf,page *TempletePage,data map[string]interface{})(error){
	//创建pdf页面
	pdf.AddPage()
	//创建pdf页面元素
	for _,element:=range page.Elements {
		err:=CreatePdfElement(pdf,&element,data)
		if err!=nil {
			return err
		}
	}
	return nil
}

func CreatePdfElement(pdf *gopdf.GoPdf,element *PageElement,data map[string]interface{})(error){
	//创建pdf元素
	switch element.Type {
	case ElementTypeText:
		err:=CreatePdfText(pdf,element,data)
		if err!=nil {
			return err
		}
	case ElementTypeLine:
		err:=CreatePdfLine(pdf,element)
		if err!=nil {
			return err
		}
	default:
		return errors.New("Unknow element type:"+element.Type)
	}
	return nil
}

func CreatePdfText(pdf *gopdf.GoPdf,element *PageElement,data map[string]interface{})(error){
	//创建pdf文本元素
	//获取文本内容
	text:=element.Content
	text=ContentFilter(text,data)
	/*if text[0]=='$' {
		//如果是变量，则从输入数据中获取变量值
		text=data[text[1:]].(string)
	}*/

	if element.MaxContentLength>0 && len(text)>element.MaxContentLength {
		text=text[0:element.MaxContentLength]
	}

	//设置字体
	if element.Font!=nil {
		//log.Println("Font:",element.Font.Family,element.Font.Style,element.Font.Size)
		pdf.SetFont(element.Font.Family, element.Font.Style, element.Font.Size)
	}
	//设置颜色
	if element.Color!=nil {
		pdf.SetTextColor(element.Color.R,element.Color.G,element.Color.B)
	}

	xPos:=element.Rect[0] 
	yPos:=element.Rect[1]
	boxWidth := element.Rect[2] - element.Rect[0]
	lineHeight:=element.LineHeight
	wordSpace:=element.WordSpace
	// 输出自动换行的文本
	return OutputWrappedText(pdf, text, lineHeight, xPos, yPos, boxWidth,wordSpace)
}

func OutputWrappedText(pdf *gopdf.GoPdf, text string, lineHeight, xPos, yPos, boxWidth,wordSpace float64)(error){
	runeSlice := []rune(text)
	left:=xPos
	top:=yPos
	for i := 0; i < len(runeSlice); i++ {
		word := string(runeSlice[i]) + "" // 加一个空格，避免中文字符连续
		wordWidth, _ := pdf.MeasureTextWidth(word)
		if left+wordWidth > boxWidth+xPos {
			left=xPos
			top=top+lineHeight
		}
		pdf.SetX(left)
		pdf.SetY(top)	
		pdf.Cell(nil, word)
		left=left+wordWidth+wordSpace
	}
	return nil
}

func CreatePdfLine(pdf *gopdf.GoPdf,element *PageElement)(error){
	//创建pdf线条元素
	//设置颜色
	if element.Color!=nil {
		pdf.SetStrokeColor(element.Color.R,element.Color.G,element.Color.B)
	}
	pdf.SetLineWidth(element.Width)
	//画线
	pdf.Line(element.Rect[0],element.Rect[1],element.Rect[2],element.Rect[3])
	return nil
}
