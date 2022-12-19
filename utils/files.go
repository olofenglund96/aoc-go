package utils

import (
	"fmt"
	"os"
)

type aocFileSystem struct {
	targetDir string
}

func NewAocFileSystem(year int, day int) aocFileSystem {
	return aocFileSystem{
		targetDir: fmt.Sprintf("years/%d/%d", year, day),
	}
}

func (fs aocFileSystem) CreateDay(input string) error {
	if fs.DayExists() {
		return fmt.Errorf("Day already exists, refusing to overwrite!")
	}

	if err := os.Mkdir(fs.targetDir, os.ModePerm); err != nil {
		fmt.Printf("Could not create directories: %e", err)
		return err
	}

	os.Chdir(fs.targetDir)

	err := fs.createSolutionFile()
	if err != nil {
		fmt.Printf("Could not create solution file: %s", err)
	}

	err = fs.createInputFile(input)
	if err != nil {
		fmt.Printf("Could not create input file: %s", err)
	}

	return nil
}

func (fs aocFileSystem) DayExists() bool {
	_, err := os.Stat(fs.targetDir)
	return err == nil
}

func (fs aocFileSystem) createSolutionFile() error {
	file, err := os.Create("solution.go")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olofenglund96/aoc-go/helpers"
)

func sol1(rows []string) string {

	return "Solution1"
}

func sol2(rows []string) string {

	return "Solution2"
}

func main() {
	rows := helpers.ReadFileLines(fmt.Sprintf("%s/%%s.dat", os.Args[2]))
	
	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
	
`, fs.targetDir))
	return err
}

func (fs aocFileSystem) createInputFile(input string) error {
	file, err := os.Create("input.dat")
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(input)

	return nil
}

func (fs aocFileSystem) SaveSolutionToFile(part string, inputFileName string, solutionString string) error {
	os.Chdir(fs.targetDir)

	file, err := os.Create(fmt.Sprintf("result%s_%s.out", part, inputFileName))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(solutionString)
	return err
}

func (fs aocFileSystem) ReadSolutionFromFile(part string) (string, error) {
	os.Chdir(fs.targetDir)

	solutionBytes, err := os.ReadFile(fmt.Sprintf("result%s_input.out", part))
	return string(solutionBytes), err
}
