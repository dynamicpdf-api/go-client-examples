package main

import (
	"fmt"
	"os"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	baseUrl := "https://api.dynamicpdf.com/"
	apiKey := "DP--api-key--"
	basePath := "./resources/creating-pdf-pdf-endpoint/"

	processCloudPdf(baseUrl, apiKey, basePath)
	processLocalPdf(baseUrl, apiKey, basePath)
}

func processLocalPdf(baseUrl string, apiKey string, basePath string) {

	pr := endpoint.NewPdf()
	pr.Endpoint.BaseUrl = baseUrl
	pr.Endpoint.ApiKey = apiKey

	layoutDataResource := resource.NewLayoutDataResource(basePath + "SimpleReportWithCoverPage.json", "SimpleReportWithCoverPage.json")
	
	theDlexResource := resource.NewDlexResourceWithPath(basePath + "SimpleReportWithCoverPage.dlex", "SimpleReportWithCoverPage.dlex")

	pr.AddNewDlexWithDlexNLayoutResources(*theDlexResource, layoutDataResource)

	additionalResource := endpoint.NewDlexWithAdditionalResource(basePath + "Northwind Logo.gif", "Northwind Logo.gif")


	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath+"pdf-dlex-pdf-local-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}
}

func processCloudPdf(baseUrl string, apiKey string, basePath string) {

	pr := endpoint.NewPdf()
	pr.Endpoint.BaseUrl = baseUrl
	pr.Endpoint.ApiKey = apiKey
	

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
		os.WriteFile(basePath+"pdf-dlex-pdf-remote-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
