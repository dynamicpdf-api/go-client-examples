package main

import (
	"fmt"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

   func RunOne(key string, basePath string) {
        
        resource := resource.NewImageResourceWithResourcePath(basePath + "/getting-started.png","")
        imageInfo := endpoint.NewImageInfo(resource)

        imageInfo.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
        imageInfo.Endpoint.ApiKey = "DP.xxx-api-key-xxx"

        resp := imageInfo.Process();
		res := <-resp
	
		if res.IsSuccessful() == true {
			
			fmt.Print(string(res.Content().Bytes()))
		}
    } 
    
    func RunTwo(key string, basePath string) {
        
        resource := resource.NewImageResourceWithResourcePath(basePath + "/multipage.tiff","")
        imageInfo := endpoint.NewImageInfo(resource)

        imageInfo.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
        imageInfo.Endpoint.ApiKey = "DP.xxx-api-key-xxx"

        resp := imageInfo.Process();
		res := <-resp
	
		if res.IsSuccessful() == true {
			
			fmt.Print(string(res.Content().Bytes()))
		}
    } 

    func main() {
      RunOne("DP.xxx-api-key-xxx",
                "C:/temp/dynamicpdf-api-samples/image-info")
      RunTwo("DP.xxx-api-key-xxx",
                "C:/temp/dynamicpdf-api-usersguide-examples/image-info")
    }
