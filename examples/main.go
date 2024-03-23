package main

import (
	"fmt"

	"github.com/devalexandre/pdfbox"
)

func main() {

	// content, err := ExtractTextFromPdf("GENERATIVE_AI_WITH_LANGCHAIN.pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	content, err := pdfbox.ExtractTextFromPdf("./GENERATIVE_AI_WITH_LANGCHAIN.pdf")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(content)
}
