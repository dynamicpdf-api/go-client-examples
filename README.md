![](./logo-banner2.png)

# go-client-examples

The Cloud API Client Examples uses the DynamicPDF Cloud API Go client library (`go-client`) to create, merge, split, form fill, stamp, obtain metadata, convert, and secure/encrypt PDF documents. This project contains numerous sample projects for the tutorials and examples on [cloud.dynamicpdf.com](https://cloud.dynamicpdf.com/).

The DynamicPDF Cloud API consists of the following endpoints.

- `dlex-layout`
- `image-info`
- `pdf`
- `pdf-info`
- `pdf-text`
- `pdf-xmp`

The GO client library (`go-client`) is available on Github ([go-client](https://github.com/dynamicpdf-api/go-client)). For more information, please visit [DynamicPDF Cloud API](https://cloud.dynamicpdf.com/). Support for other languages/platforms (Java, PHP, C#, Node.js, Python) is available on GitHub ([DynamicPDF Cloud API at GitHub](https://github.com/dynamicpdf-api)).

## Package

* The DynamicPDF.API library is available via the GO package library ([go-client](https://pkg.go.dev/github.com/dynamicpdf-api/go-client/v2)) or on Github ([go-client](https://github.com/dynamicpdf-api/go-client)).


## Resources

To obtain the resources for the project, login to [cloud.dynamicpdf.com](https://cloud.dynamicpdf.com/) (assuming you have an account), and go to the **Resource Manager**. You use the `samples` folder to add the resources for the tutorials and examples from this project. Local resources are available in the resources folder of this project. 

- [Resource Manager Samples](https://cloud.dynamicpdf.com/docs/usersguide/environment-manager/environment-manager-sample-resources)  

For more information on the tutorials and example code, refer to:

- https://cloud.dynamicpdf.com/docs/tutorials/tutorials-overview
- https://cloud.dynamicpdf.com/docs/usersguide/cloud-api/cloud-api-overview

## **Tutorials**

The following table lists the available tutorials.

| Tutorial Title                                     | Project/File/Class      | Tutorial Location                                            |
| -------------------------------------------------- | ----------------------- | ------------------------------------------------------------ |
| Merging PDFs                                       | `merge-pdfs.go`               | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/merging-pdfs |
| Completing an AcroForm                             | `fill-acroform.go`    | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/form-completion |
| Creating a PDF Using a DLEX and the `pdf` Endpoint | `dlex-dlex-example.go`       | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/dlex-pdf-endpoint |
| Adding Bookmarks to a PDF                          | `add-bookmarks-example.go`          | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/bookmarks |
| Creating a PDF Using the `dlex-layout` Endpoint    | `dlex-layout-example.go` | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/dlex-layout |
| Extracting Image Metadata                          | `image-info-example.go`          | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/image-info |
| Extract PDF Metadata                               | `pdf-info-example.go`            | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/pdf-info |
| Extracting PDF's Text                              | `pdf-text-example.go`        | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/pdf-text |
| Extract XMP Metadata                               | `pdf-xmp-example.go`        | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/pdf-xmp |
| Convert HTML to PDF                                | `pdf-html-example.go`   | https://cloud.dynamicpdf.com/docs/tutorials/cloud-api/pdf-html |

# Support

The primary source for the DynamicPDF Cloud API support is through [Stack Overflow](https://stackoverflow.com/questions/tagged/dynamicpdf-api). Please use the "[dynamicpdf-api](https://stackoverflow.com/questions/tagged/dynamicpdf-api)" tag to ask questions. Our support team actively monitors the tag and responds promptly to any questions.  Also, let us know you asked the question by following up with an email to [support@dynamicpdf.com](mailto:support@dynamicpdf.com). 

## Pro Plan Subscribers[#](https://cloud.dynamicpdf.com/support#pro-plan-subscribers)

Ticket support is available to Pro Plan subscribers. But we still encourage you to help the community by posting on Stack Overflow when possible. You can also email [support@dynamicpdf.com](mailto:support@dynamicpdf.com) if you need to ask something specific to your use case that may not help the DynamicPDF Cloud API community.

# License

The `dotnet-client-examples` library is licensed under the [MIT License](./LICENSE).
