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

func (bp *bookPage) addNewPageToEnd(text string) {
	newPage := &bookPage{text, nil}
	for bp.nextPage != nil { // the next page will only be nil if the current page is the last page
		bp = bp.nextPage // asssign the bookpage to the last page
	}

	bp.nextPage = newPage // add a new page

	//or
	//bp.nextPage = &bookPage{text, nil}
}

func (bp *bookPage) addNewPageAfter(text string) {
	newpage := &bookPage{text, bp.nextPage}
	bp.nextPage = newpage
}

func main() {
	pg := bookPage{"First page", nil}
	pg.addNewPageToEnd("Second page")
	pg.addNewPageToEnd("Third page")
	pg.addNewPageToEnd("Fourth page")
	pg.readContent()

	fmt.Println("------------")

	pg.addNewPageAfter("Before second page")
	pg.readContent()
}
