package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/templates/"
	apiKey = "DP--api-key--"
	outputPath = "./output/pdf-page-elements-go-output.pdf"
}

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.ApiKey = apiKey
	pageInput := pdfExample.AddPageWithDimension(1008, 612)

	textElement := element.NewText("Hello PDF", element.TopCenter, 50, 100)
	textElement.SetColor(color.NewRgbColorDefault().Blue().Color)
	textElement.SetFontSize(42)


	lineElement := element.NewLine(element.TopLeft, 900, 150)
	lineElement.XOffset = 305
	lineElement.YOffset = 150
	lineElement.SetColor(color.NewRgbColorDefault().Red().Color)
	lineElement.SetWidth(4)
	
	lineStyle := lineElement.LineStyle()
	lineElement.SetLineStyle(*lineStyle.Solid())

	rectangle := element.NewRectangle(element.TopCenter, 100, 500)
	rectangle.SetCornerRadius(10)
	rectangle.SetBorderWidth(5)
	
	
	rectangle.SetBorderStyle(*lineStyle.Dots())
	
	
	rectangle.XOffset = -250
	rectangle.YOffset = -10
	rectangle.SetBorderColor(color.NewRgbColorDefault().Blue().Color)
	rectangle.SetFillColor(color.NewRgbColorDefault().Green().Color)

	imageResource := resource.NewImageResourceWithResourcePath(basePath+"dynamicPdfLogo.png", "dynamicPdfLogo.png")

	imageElement := element.NewImagewithImageResource(imageResource, element.TopLeft, 835, 75)
	

	pageInput.Elements = append(pageInput.Elements, imageElement)

	pageInput.Elements = append(pageInput.Elements, rectangle)

	pageInput.Elements = append(pageInput.Elements, textElement)
	pageInput.Elements = append(pageInput.Elements, lineElement)

	resp := pdfExample.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath)
		os.WriteFile(outputPath,
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}
}
