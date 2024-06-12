package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
)

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pr.Endpoint.ApiKey = "DP--api-key--"
	basePath := "c:/temp/solutions/image-conversion/"

	//imageResource2 := resource.NewImageResourceWithResourcePath(basePath+"dynamicpdfLogo.png", "dynamicpdfLogo.png")

	imageInput1 := input.NewImageWithResourcePath(basePath + "testimage.tif")
	//imageInput2 := input.NewImageWithResourcePath(imageResource2)

	pr.Inputs = append(pr.Inputs, imageInput1)
	//pr.Inputs = append(pr.Inputs, imageInput2)

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath+"image-conversion-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
