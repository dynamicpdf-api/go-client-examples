package main

import (
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	layoutDataResource := resource   .NewLayoutDataResource("c:/temp/dynamicpdf-api-samples/dlex-layout-example/SimpleReportWithCoverPage.json", "SimpleReportWithCoverPage.json")

	layoutData := endpoint.NewDlexEndpoint("samples/dlex-layout/SimpleReportWithCoverPage.dlex", layoutDataResource)
	layoutData.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	layoutData.Endpoint.ApiKey = "<API-KEY>"
	resp := layoutData.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("c:/temp/dynamicpdf-api-samples/dlex-layout-example/output.pdf", res.Content().Bytes(), os.ModeType)
	}
}
