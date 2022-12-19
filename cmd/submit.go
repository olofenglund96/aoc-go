/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/olofenglund96/aoc-go/utils"
	"github.com/spf13/cobra"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		year, day := get_day_and_year()

		aocFacade, err := utils.NewAocHttpClient()
		if err != nil {
			panic(err)
		}

		part, err := cmd.Flags().GetString("part")
		if err != nil {
			panic(err)
		}

		fs := utils.NewAocFileSystem(year, day)

		solution, err := fs.ReadSolutionFromFile(part)
		if err != nil {
			panic(err)
		}

		resp, err := aocFacade.SubmitDay(year, day, part, solution)

		if err != nil {
			panic(err)
		}

		if strings.Contains(resp, "That's the right answer") {
			fmt.Println("Correct answer!!")
		} else {
			fmt.Println(resp)
			fmt.Println("Not quite right..")
		}
	},
}

func init() {
	dayCmd.AddCommand(submitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	submitCmd.Flags().StringP("part", "p", "1", "The part of the problem to solve (1/2)")
}
