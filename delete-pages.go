package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pr.Endpoint.ApiKey = "DP--api-key--"
	basePath := "c:/temp/dynamicpdf-api-samples/delete-pages/"

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath + "pdfnumberedinput.pdf", "pagenumberedinput1.pdf")
	prInput := input.NewPdfWithResource(pdfResource)
	prInput.StartPage = 1
	prInput.PageCount = 3

	pr.Inputs = append(pr.Inputs, prInput)

	pdfResource2 := resource.NewPdfResourceWithResourcePath(basePath + "pdfnumberedinput.pdf", "pagenumberedinput2.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource2)
	prInput2.StartPage = 6
	prInput2.PageCount = 2

	pr.Inputs = append(pr.Inputs, prInput2)

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath +"delete-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
