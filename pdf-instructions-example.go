package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/font"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

func topLevelMetadata() *endpoint.Pdf {
	tlm := endpoint.NewPdf()
	pageInput := input.NewPage()
	tlm.Author = "John Doe"
	tlm.Creator = "John Creator"
	tlm.Keywords = "dynamicpdf api example pdf java instructions"
	tlm.Title = "Sample PDF"
	pageInput.SetPageHeight(612)
	pageInput.SetPageWidth(1008)

	tlm.Inputs = append(tlm.Inputs, pageInput)

	return tlm
}

func mergeExample(basePath string) *endpoint.Pdf {
	pr := endpoint.NewPdf()
	mergeOption := input.NewMergeOptions()
	prInput := input.NewPdfWithCloudPath("samples/merge-pdfs-pdf-endpoint/DocumentB.pdf", mergeOption)
	pr.Inputs = append(pr.Inputs, prInput)
	imageResource := resource.NewImageResourceWithResourcePath(basePath+"DPDFLogo.png", "DPDFLogo.png")
	pr.AddImage(imageResource)
	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"DocumentA.pdf", "DocumentA.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource)
	pr.Inputs = append(pr.Inputs, prInput2)
	return pr
}

func acroFormExample(basePath string) *endpoint.Pdf {

	pdfAcro := endpoint.NewPdf()
	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"simple-form-fill.pdf", "simple-form.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	pdfAcro.Inputs = append(pdfAcro.Inputs, pdfInput)

	field1 := endpoint.NewFormFieldWithValue("nameField", "DynamicPdf")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field1)
	field2 := endpoint.NewFormFieldWithValue("descriptionField", "RealTime Pdf's. Real FAST!")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field2)

	return pdfAcro
}

func outlineExample() *endpoint.Pdf {

	pdfOut := endpoint.NewPdf()

	pageInputOne := pdfOut.AddPage()
	textOne := element.NewText("Hello World 1", element.BottomCenter, 250, 0)
	pageInputOne.Elements = append(pageInputOne.Elements, textOne)

	pdfOut.Inputs = append(pdfOut.Inputs, pageInputOne)

	pageInputTwo := pdfOut.AddPage()
	textTwo := element.NewText("Hello World 2", element.BottomCenter, 250, 0)
	pageInputTwo.Elements = append(pageInputTwo.Elements, textTwo)

	pageInputThree := pdfOut.AddPage()
	textThree := element.NewText("Hello World 3", element.BottomCenter, 250, 0)
	pageInputThree.Elements = append(pageInputThree.Elements, textThree)

	outline := pdfOut.Outlines.Add("Root Outline")
	outline.Expanded = true

	out := outline.Children()

	out.AddWithInputValue("Page 1", pageInputOne.Input, 0, endpoint.FitPage)
	out.AddWithInputValue("Page 2", pageInputTwo.Input, 0, endpoint.FitPage)
	out.AddWithInputValue("Page 3", pageInputOne.Input, 0, endpoint.FitPage)

	return pdfOut
}

func addOutlinesExistingPdf(basePath string) *endpoint.Pdf {

	pdfOut := endpoint.NewPdf()
	pdfOut.Author = "John Doe"
	pdfOut.Title = "Existing Pdf Example"

	var resource1 = resource.NewPdfResourceWithResourcePath(basePath+"AllPageElements.pdf", "AllPageElements.pdf")
	var input1 = pdfOut.AddPdf(resource1, input.NewMergeOptions())
	input1.SetId("AllPageElements")

	var resource2 = resource.NewPdfResourceWithResourcePath(basePath+"OutlineExisting.pdf", "OutlineExisting.pdf")
	var input2 = pdfOut.AddPdf(resource2, input.NewMergeOptions())
	input2.SetId("outlineDoc1")

	var rootOutline = pdfOut.Outlines.Add("Imported Outline")
	rootOutline.Expanded = true

	out := rootOutline.Children()
	out.AddPdfOutlines(*input1)
	out.AddPdfOutlines(*input2)

	return pdfOut
}

func imageExample(basePath string) *endpoint.Pdf {
	prImage := endpoint.NewPdf()
	imageResource := resource.NewImageResourceWithResourcePath(basePath+"A.png", "A.png")
	prImage.AddImage(imageResource)
	prImage.AddImageCloudPath("samples/get-image-info-image-info-endpoint/dynamicpdfLogo.png")
	return prImage
}

func barcodeExample() *endpoint.Pdf {
	barcodePdf := endpoint.NewPdf()
	barcode := element.NewAztecBarcodeElement("Hello World", element.BottomRight, 50.0, 50.0)
	barcode.SetSymbolSize(element.R105xC105)
	barcode.SetXDimension(2)
	barcode.XOffset = 0
	barcode.YOffset = 10
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcode)
	barcodePdf.Inputs = append(barcodePdf.Inputs, pdfInput)

	return barcodePdf
}

func fontsExample(basePath string) *endpoint.Pdf {

	fe := endpoint.NewPdf()

	pageInput := input.NewPage()
	pageNumber := element.NewPageNumberingElement("A", "TopLeft", 0, 0)
	pageNumber.SetFontSize(42)
	pageNumber.SetFont(font.HelveticaBoldOblique())

	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	pageNumberingElementTwo := element.NewPageNumberingElement("B", "TopRight", 0, 0)
	pageNumberingElementTwo.SetColor(color.NewRgbColorDefault().DarkGreen().Color)
	myFont := font.NewFontResource(basePath+"cnr.otf", "cnr.otf")
	pageNumberingElementTwo.SetFont(*myFont)
	pageNumberingElementTwo.SetFontSize(32)

	pageInput.Elements = append(pageInput.Elements, pageNumber)
	pageInput.Elements = append(pageInput.Elements, pageNumberingElementTwo)
	fe.Inputs = append(fe.Inputs, pageInput)
	return fe
}

func securityExample(basePath string) *endpoint.Pdf {
	secPdf := endpoint.NewPdf()
	secRes := resource.NewPdfResourceWithResourcePath(basePath+"DocumentA.pdf", "DocumentA.pdf")
	secInput := input.NewPdfWithResource(secRes)

	security := endpoint.NewAes128Security("user", "owner")
	secPdf.Security = endpoint.Security(security.Security)
	secPdf.Inputs = append(secPdf.Inputs, secInput)
	return secPdf
}

func pdfHtmlExample() *endpoint.Pdf {

	pdfExample := endpoint.NewPdf()

	inputOne := input.NewHtmlInputWithString("<html>An example HTML fragment.</html>")

	inputTwo := input.NewHtmlInputWithString("<html><p>HTML with basePath.</p><img src='./images/logo.png'></html>")
	inputTwo.SetBasePath("https://www.dynamicpdf.com")

	pdfExample.Inputs = append(pdfExample.Inputs, inputOne)
	pdfExample.Inputs = append(pdfExample.Inputs, inputTwo)

	return pdfExample
}

func templateExample(basePath string) *endpoint.Pdf {
	pdfTemp := endpoint.NewPdf()
	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"DocumentA.pdf", "DocumentA.pdf")
	pdfInput1 := input.NewPdfWithResource(pdfResource)
	temp := element.NewTemplate()
	txt := element.NewText("Hello World", element.TopCenter, 20, 10)
	txt.SetColor(color.NewRgbColorDefault().Red().Color)
	txt.SetFontSize(32)
	temp.Elements = append(temp.Elements, txt)
	pdfInput1.SetTemplate(temp)
	pdfTemp.Inputs = append(pdfTemp.Inputs, pdfInput1)
	return pdfTemp
}

//func dlexResourceExample(basePath string) *endpoint.Pdf {
//	pdfD := endpoint.NewPdf()

//	imageResource := resource.NewImageResourceWithResourcePath(basePath+"DPDFLogo.png", "DPDFLogo.png")

//	pdfD.AddDlexWithCloudResourceNLayoutDataPath("samples/users-guide-resources/SimpleReportWithCoverPage.dlex",
//		basePath+"SimpleReportWithCoverPage.json")
//	return pdfD
//}

func googleFontsExample() *endpoint.Pdf {

	pdfCl := endpoint.NewPdf()

	pageInput := input.NewPage()
	pageNumber := element.NewPageNumberingElement("test", "TopLeft", 0, 0)
	pageNumber.SetFontSize(72)
	pageNumber.TextFont()
	pageNumber.SetFont(*font.Google("Borel"))

	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	myTextElement := element.NewText("myTest", element.TopLeft, 100, 200)
	myTextElement.SetFont(*font.Google("Borel"))

	pageInput.Elements = append(pageInput.Elements, myTextElement)

	pageInput.Elements = append(pageInput.Elements, pageNumber)
	pdfCl.Inputs = append(pdfCl.Inputs, pageInput)

	return pdfCl
}

func main() {

	var theBasePath = "./resources/users-guide/"
	var theBaseUrl = "https://api.dynamicpdf.com/"
	var theOutputPath = "./output/"
	var apiKey = "Dp--api-key--"

	pdfTlm := topLevelMetadata()
	process(pdfTlm, theOutputPath+"toplevelmetadata-ouput.pdf", theBaseUrl, apiKey)

	pdfAcro := acroFormExample(theBasePath)
	process(pdfAcro, theOutputPath+"pdfAcroExample-output.pdf", theBaseUrl, apiKey)

	pdfBar := barcodeExample()
	process(pdfBar, theOutputPath+"bar-output.pdf", theBaseUrl, apiKey)

	pdfFnt := fontsExample(theBasePath)
	process(pdfFnt, theOutputPath+"fnt-output.pdf", theBaseUrl, apiKey)

	pdfSec := securityExample(theBasePath)
	process(pdfSec, theOutputPath+"sec-output.pdf", theBaseUrl, apiKey)

	pdfHtmlExample := pdfHtmlExample()
	process(pdfHtmlExample, theOutputPath+"pdfHtmlExample-output.pdf", theBaseUrl, apiKey)

	pdfTemp := templateExample(theBasePath)
	process(pdfTemp, theOutputPath+"pdfTempExample-output.pdf", theBaseUrl, apiKey)

	pdfGF := googleFontsExample()
	process(pdfGF, theOutputPath+"pdfgooglefont-output.pdf", theBaseUrl, apiKey)

	//pdfdlexResourceExample := dlexResourceExample(theBasePath)
	//process(pdfdlexResourceExample, theOutputPath +"dlexExample-output.pdf", theBaseUrl, apiKey)

	pdfIm := imageExample(theBasePath)
	process(pdfIm, theOutputPath+"pdfImage-output.pdf", theBaseUrl, apiKey)

	pdfMerge := mergeExample(theBasePath)
	process(pdfMerge, theOutputPath+"go-merge-example-output.pdf", theBaseUrl, apiKey)

	pdfOut := outlineExample()
	process(pdfOut, theOutputPath+"outline-example-output.pdf", theBaseUrl, apiKey)

	pdfOut2 := addOutlinesExistingPdf(theBasePath)
	process(pdfOut2, theOutputPath+"outline-existing-example-output.pdf", theBaseUrl, apiKey)

}

func process(thePdf *endpoint.Pdf, outputFilePath string, baseUrl string, apiKey string) {

	thePdf.Endpoint.BaseUrl = baseUrl
	thePdf.Endpoint.ApiKey = apiKey

	resp := thePdf.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputFilePath)
		os.WriteFile(outputFilePath,
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("Error Message:" + res.ErrorMessage())
		fmt.Print(res.ErrorJson())
	}

}
