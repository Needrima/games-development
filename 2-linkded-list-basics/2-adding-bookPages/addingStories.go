package main

import (
	"fmt"
)

type bookPage struct {
	content  string    //page content
	nextPage *bookPage // link to next page
}

func (bp *bookPage) readContent() {
	for bp != nil {
		fmt.Println(bp.content)
		bp = bp.nextPage
	}
}

func (bp *bookPage) addNewPage(text string) {
	newPage := &bookPage{text, nil}
	for bp.nextPage != nil {
		bp = bp.nextPage
	}

	bp.nextPage = newPage
	//bp.nextPage = &bookPage{text, nil}
}

func main() {
	pg1 := bookPage{"First page", nil}
	pg2 := bookPage{"Second page", nil}
	pg3 := bookPage{"Third page", nil}

	pg1.nextPage = &pg2 //page 1 is a linked list containing pages 2 & 3
	pg2.nextPage = &pg3 //page 2 is a linked list containing page 3

	pg1.readContent() // reads contents from pages 1 through 3
	// pg2.readContent() // reads contents from pages 2 through 3
	// pg3.readContent() // reads content from page 3
}
