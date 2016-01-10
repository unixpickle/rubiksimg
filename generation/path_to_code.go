package main

import (
	"fmt"
	"os"
	"strconv"
)

type Command struct {
	Name      string
	Arguments []float64
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: path_to_code <SVG path data>")
		os.Exit(1)
	}

	commands := readCommands(os.Args[1])
	for _, command := range commands {
		switch command.Name {
		case "M":
			fmt.Printf("ctx.MoveTo(scale*%f, scale*%f)\n", command.Arguments[0],
				command.Arguments[1])
		case "L":
			fmt.Printf("ctx.LineTo(scale*%f, scale*%f)\n", command.Arguments[0],
				command.Arguments[1])
		case "C":
			fmt.Printf("ctx.CubicCurveTo(scale*%f, scale*%f, scale*%f,\n\tscale*%f, scale*%f, "+
				"scale*%f)\n",
				command.Arguments[0], command.Arguments[1],
				command.Arguments[2], command.Arguments[3],
				command.Arguments[4], command.Arguments[5])
		case "z":
			fmt.Printf("ctx.Close()\n")
		default:
			fmt.Fprintln(os.Stderr, "Unknown command: "+command.Name)
			os.Exit(1)
		}
	}
}

func readCommands(s string) []Command {
	res := make([]Command, 0)
	var currentCommand *Command
	var currentNumArg string
	for _, ch := range s {
		if (ch < '0' || ch > '9') && ch != '.' && ch != '-' {
			if currentCommand != nil && len(currentNumArg) > 0 {
				num, _ := strconv.ParseFloat(currentNumArg, 64)
				currentCommand.Arguments = append(currentCommand.Arguments, num)
				currentNumArg = ""
			}
		} else if currentCommand != nil {
			currentNumArg += string(ch)
		}
		if ch == 'M' || ch == 'L' || ch == 'C' || ch == 'z' {
			if currentCommand != nil {
				res = append(res, *currentCommand)
			}
			currentCommand = &Command{string(ch), []float64{}}
			currentNumArg = ""
		}
	}
	if currentCommand != nil {
		res = append(res, *currentCommand)
	}
	return res
}
