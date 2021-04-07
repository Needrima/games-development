package main

import (
	"fmt"
)

type bookPage struct {
	content  string
	nextPage *bookPage
}

func (b *bookPage) readContent() {
	fmt.Println(b.content)

	if b.nextPage == nil {
		return
	}

	b.readContent()
}

func main() {
	pg1 := bookPage{"First page", nil}
	pg2 := bookPage{"Second page", nil}
	pg3 := bookPage{"Third page", nil}

	pg1.nextPage = &pg2
	pg2.nextPage = &pg3

	pg1.readContent()
}
