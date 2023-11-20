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
	pr.Endpoint.ApiKey = "DP.xxx-api-key-xxx"
	basePath := "c:/temp/dynamicpdf-api-samples/"

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"DocumentA.pdf", "DocumentA.pdf")
	prInput := input.NewPdfWithResource(pdfResource)
	prInput.StartPage = 1
	prInput.PageCount = 1

	pr.Inputs = append(pr.Inputs, prInput)

	pdfResource2 := resource.NewPdfResourceWithResourcePath(basePath+"DocumentB.pdf", "DocumentB.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource2)
	pr.Inputs = append(pr.Inputs, prInput2)

	mergeOption := input.NewMergeOptions()
	prInput3 := input.NewPdfWithCloudPath("samples/merge-pdfs-pdf-endpoint/DocumentC.pdf", mergeOption)
	pr.Inputs = append(pr.Inputs, prInput3)

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath+"merge-pdfs-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
