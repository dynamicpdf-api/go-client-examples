package main

import (
	"fmt"
	"os"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pr.Endpoint.ApiKey = "DP.xxx-api-key-xxx"
	basePath := "c:/temp/dynamicpdf-api-samples/"

	layoutDataResource := resource.NewLayoutDataResource(basePath+"SimpleReportWithCoverPage.json", "SimpleReportWithCoverPage.json")

	pr.AddDlexWithCloudResourceNLayoutData("samples/creating-pdf-pdf-endpoint/SimpleReportWithCoverPage.dlex", layoutDataResource)

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath+"pdf-dlex-pdf-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
