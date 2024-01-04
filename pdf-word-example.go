package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pdfExample.Endpoint.ApiKey = "DP--api-key--"
	wordResource := resource.NewWordResourceWithResourcePath("./resources/word-pdf/Doc1.docx", "Doc1.docx")
	pdfInput := input.NewWordInputWithResource(wordResource)
	pdfExample.Inputs = append(pdfExample.Inputs, pdfInput)
	resp := pdfExample.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("./output/word-to-pdf-output-go.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}

}
