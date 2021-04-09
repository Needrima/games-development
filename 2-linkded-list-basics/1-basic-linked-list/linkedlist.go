package main

import (
	"fmt"
)

type bookPage struct {
	content  string
	nextPage *bookPage
}

func (bp *bookPage) readContent() {
	/*
		fmt.Println(bp.content)

		if bp.nextPage == nil {
			return
		}

		bp.nextPage.readContent()
	*/
	//recursion is not really advised especially when dealing with multiple calls
	//for loop is a better way to handle stuff

	for bp != nil {
		fmt.Println(bp.content)
		bp = bp.nextPage
	}
}

func main() {
	pg1 := bookPage{"First page", nil}
	pg2 := bookPage{"Second page", nil}
	pg3 := bookPage{"Third page", nil}

	pg1.nextPage = &pg2
	pg2.nextPage = &pg3

	pg1.readContent()
}
