/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"os/exec"

	"strings"

	"github.com/olofenglund96/aoc-go/utils"
	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		year, day := get_day_and_year()
		part, err := cmd.Flags().GetString("part")
		if err != nil {
			panic(err)
		}

		solveExample, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}

		inputFileName := "input"
		if solveExample {
			inputFileName = "example"
		}

		fmt.Printf("Solving %d/%d, part %s (%s)..\n", year, day, part, inputFileName)

		// get `go` executable path
		out, err := exec.Command("go", "run", fmt.Sprintf("years/%d/%d/solution.go", year, day), part, inputFileName).CombinedOutput()

		if err != nil {
			panic(fmt.Sprintf("Error: %s, out: %v", err, string(out)))
		}

		solutionOut := string(out)
		fmt.Println(solutionOut)
		solutionOutSlice := strings.Split(solutionOut, "\n")

		aocFileSystem := utils.NewAocFileSystem(year, day)
		err = aocFileSystem.SaveSolutionToFile(part, inputFileName, string(solutionOutSlice[len(solutionOutSlice)-1]))

		if err != nil {
			panic(fmt.Sprintf("Error: %s", err))
		}
	},
}

func init() {
	dayCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	solveCmd.Flags().StringP("part", "p", "1", "The part of the problem to solve (1/2)")
	solveCmd.Flags().BoolP("example", "e", false, "Solve the example")
}
