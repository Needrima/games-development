package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd, description string
	nextNode         *storyNode
}

type storyNode struct {
	story   string
	choices []*choice
}

func (node *storyNode) addChoice(cmd, description string, nextNode *storyNode) {
	newChoice := choice{cmd, description, nextNode}
	node.choices = append(node.choices, &newChoice)
}

func (node *storyNode) renderNode() {
	fmt.Println(node.story)

	for _, choice := range node.choices {
		fmt.Printf("%s : %s\n", choice.cmd, choice.description)
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("No direction with said command")
	return node
}

func (node *storyNode) play() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		node.renderNode()
		scanner.Scan()

		cmd := scanner.Text()

		node = node.executeCmd(cmd)

		if node.choices == nil {
			fmt.Println(node.story)
			break
		}
	}
}

func main() {
	start := storyNode{"Starting floor", nil}
	northPath := storyNode{"You ran into a troll and it devours you; Game Over", nil}
	southPath := storyNode{"You entered a dark room", nil}
	eastPath := storyNode{"You you ran into a cyclops and it crushed you; Game over", nil}
	westPath := storyNode{"You went to the west floor", nil}
	quit := storyNode{"You quit", nil}

	//if player goes south
	roomLit := storyNode{"You switched the light on", nil}
	roomLitEast := storyNode{"You entered a room full of treasure, Congratulations, you won", nil}
	roomLitWest := storyNode{"You ran into a troll and it devours you; Game Over", nil}

	//if player goes west
	westPathWest := storyNode{"You ran into a cyclop and it crushses you; Game over", nil}
	westPathEast := storyNode{"You entered a room full of treasure, Congratulations, you won", nil}

	start.addChoice("N", "Go north", &northPath)
	start.addChoice("S", "Go south", &southPath)
	start.addChoice("E", "Go east", &eastPath)
	start.addChoice("W", "Go west", &westPath)
	start.addChoice("Q", "Quit game", &quit)

	//north path
	northPath.addChoice("S", "Start again", &start)
	northPath.addChoice("Q", "Quit game", &quit)

	//south path
	southPath.addChoice("L", "Switch light on", &roomLit)
	roomLit.addChoice("E", "Go east", &roomLitEast)
	roomLit.addChoice("W", "Go west", &roomLitWest)
	roomLitEast.addChoice("S", "Start again", &start)
	roomLitEast.addChoice("Q", "Quit game", &quit)
	roomLitWest.addChoice("S", "Start again", &start)
	roomLitWest.addChoice("Q", "Quit game", &quit)

	//west path
	westPath.addChoice("W", "Go west", &westPathWest)
	westPath.addChoice("E", "Go east", &westPathEast)
	westPathWest.addChoice("S", "Start again", &start)
	westPathWest.addChoice("Q", "Quit game", &quit)
	westPathEast.addChoice("S", "Start again", &start)
	westPathEast.addChoice("Q", "Quit game", &quit)

	//east path
	eastPath.addChoice("S", "Start again", &start)
	eastPath.addChoice("Q", "Quit game", &quit)

	start.play()
}
