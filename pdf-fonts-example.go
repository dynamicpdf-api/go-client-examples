package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/font"
	"github.com/dynamicpdf-api/go-client/input"
)

func main() {

	pdfCl := endpoint.NewPdf()
	pdfCl.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pdfCl.Endpoint.ApiKey = "DP.xxx-api-key-xxx"

	pageInput := input.NewPage()
	pageNumber := element.NewPageNumberingElement("A", "TopLeft", 0, 0)
	pageNumber.SetFontSize(42)
	pageNumber.SetFont(font.HelveticaBoldOblique())

	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	pageNumberingElementTwo := element.NewPageNumberingElement("B", "TopRight", 0, 0)
	pageNumberingElementTwo.SetColor(color.NewRgbColorDefault().DarkGreen().Color)
	myFont := font.NewFontResource("c:/temp/fonts-example/cnr.otf", "cnr.otf")
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
		os.WriteFile("C:/temp/fonts-example/pdf-fonts-example-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
