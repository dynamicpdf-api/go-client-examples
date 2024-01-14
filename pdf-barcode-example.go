package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
)

func main() {

	barcodePdf := endpoint.NewPdf()
	barcodePdf.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	barcodePdf.Endpoint.ApiKey = "DP.--api-key--"

	pageInput := input.NewPage()
	pageInput.PageHeight = 621
	pageInput.PageWidth = 1008

	barcode := element.NewCode11Barcode("12345678910", element.TopCenter, 200, 50, 50)
	barcode.SetColor(color.NewRgbColorDefault().Red().Color)

	pageInput.Elements = append(pageInput.Elements, barcode)
	barcodePdf.Inputs = append(barcodePdf.Inputs, pageInput)

	resp := barcodePdf.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("./output/barcode-example-output.pdf",
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("Error Message:" + res.ErrorMessage())
		fmt.Print(res.ErrorJson())
	}
}
