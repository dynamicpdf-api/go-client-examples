package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/font"
	"github.com/dynamicpdf-api/go-client/v2/input"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/users-guide/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/pdf-fonts-example-go-output.pdf"
}

func main() {

	pdfCl := endpoint.NewPdf()
	pdfCl.Endpoint.ApiKey = apiKey

	pageInput := input.NewPage()
	pageNumber := element.NewPageNumberingElement("A", "TopLeft", 0, 0)
	pageNumber.SetFontSize(42)
	pageNumber.SetFont(font.HelveticaBoldOblique())

	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	pageNumberingElementTwo := element.NewPageNumberingElement("B", "TopRight", 0, 0)
	pageNumberingElementTwo.SetColor(color.NewRgbColorDefault().DarkGreen().Color)
	myFont := font.NewFontResource(basePath+"cnr.otf", "cnr.otf")
	pageNumberingElementTwo.SetFont(*myFont)
	pageNumberingElementTwo.SetFontSize(32)

	pageInput.Elements = append(pageInput.Elements, pageNumber)
	pageInput.Elements = append(pageInput.Elements, pageNumberingElementTwo)
	pdfCl.Inputs = append(pdfCl.Inputs, pageInput)

	resp := pdfCl.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.Remove(outputPath)
		os.WriteFile(outputPath,
			res.Content().Bytes(), os.ModeType)
	}

}
