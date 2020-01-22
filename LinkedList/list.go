package LinkedList

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
	"strings"
)

type List interface {
	Add(value interface{}) error
	Remove(value interface{}) error
	Show() error
}

func ShowHelp() {
	fmt.Println("")
	fmt.Println("a  -> Add to the list")
	fmt.Println("r -> Remove from list")
	fmt.Println("s -> Show all elements")
	fmt.Println("x | q-> Exit")
	fmt.Println("**************************************")
}

// command handler for cli context
func Run(cli *cli.Context) {
	// create a singly linked list
	singlyList := SinglyList{
		ListName: "Singly List",
	}

	reader := bufio.NewReader(os.Stdin)
	// main loop
	ShowHelp()
	for {
		fmt.Print("> ")

		byteValue, _, _ := reader.ReadLine()

		ch := strings.ToLower(string(byteValue))

		switch ch {
		case "a":
			// add to the list
			fmt.Print("Enter data: ")
			newByte, _, _ := reader.ReadLine()
			newIntValue, _ := strconv.Atoi(string(newByte))
			err := singlyList.Add(newIntValue)

			if err != nil {
				log.Fatal(err)
			}
		case "r":
			fmt.Print("Enter value to remove: ")
			byteToRemove, _, _ := reader.ReadLine()
			value, _ := strconv.Atoi(string(byteToRemove))
			err := singlyList.Remove(value)
			if err != nil {
				log.Fatal(err)
			}
		case "s":
			singlyList.Show()
		case "as":
			singlyList.Sort(0)
		case "ds":
			singlyList.Sort(1)
		case "c":
			singlyList.Count()
		case "x":
			fallthrough
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Error: Invalid Input!")
			ShowHelp()
		}
	}

}
