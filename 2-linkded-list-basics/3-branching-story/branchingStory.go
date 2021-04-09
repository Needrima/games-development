package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	content string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) Play() {
	fmt.Println(node.content)

	if node.yesPath != nil && node.noPath != nil {
		scanner := bufio.NewScanner(os.Stdin)

		for {
			scanner.Scan()
			resp := scanner.Text()

			if resp == "yes" {
				node.yesPath.Play()
				break
			} else if resp == "no" {
				node.noPath.Play()
				break
			} else {
				fmt.Println("Wrong answer, answer yes or no")
			}
		}
	}
}

func (node *storyNode) PlayAll() {
	fmt.Println(node.content)

	if node.yesPath != nil {
		node.yesPath.PlayAll()
	}

	if node.noPath != nil {
		node.noPath.PlayAll()
	}
}

func main() {
	start := storyNode{"You are at the entrance. Enter?", nil, nil}
	yesPath := storyNode{"You should not have entered, you are dead.", nil, nil}
	noPath := storyNode{"Wise Choice, enjoy the rest of your life", nil, nil}

	start.yesPath = &yesPath
	start.noPath = &noPath

	start.Play()

	//start.PlayAll()
}
