# PDFBox Go Wrapper

This Go package provides a wrapper around the PDFBox Java library, enabling Go applications to extract text from PDF files.

## Requirements

- Java 11 or newer.
- Go 1.15 or newer.

This package uses the PDFBox utility (`pdfbox-app-3.0.2.jar`) embedded within the Go code to ensure that the extraction process is handled seamlessly.

## Installation

To use the `pdfbox` package in your Go project, add it as a dependency:

```bash
go get github.com/devalexandre/pdfbox
```

## Usage
The main functionality of this package is to extract text from a given PDF file. Here is how you can use it:

```go
package main

import (
	"fmt"
	"log"
	"github.com/devalexandre/pdfbox"
)

func main() {
	pdfFilePath := "path/to/your/file.pdf"
	text, err := pdfbox.ExtractTextFromPdf(pdfFilePath)
	if err != nil {
		log.Fatalf("Failed to extract text from PDF: %v", err)
	}
	fmt.Println(*text)
}
```

## How It Works

-  Extracts the embedded pdfbox-app-3.0.2.jar file from the Go binary.
-  Creates a temporary JAR file in the system's temporary directory.
-  Executes a Java command using the temporary JAR to extract text from the specified PDF file.
-  Cleans up by removing the generated text and temporary JAR files after extraction.

Note: The extracted text is returned as a string pointer. If an error occurs during the extraction process, the error is returned.

## Contributing
Contributions to this package are welcome. Please ensure that your code complies with the project's coding standards and includes appropriate tests.

## License
This project is licensed under the MIT License. See the LICENSE file for details