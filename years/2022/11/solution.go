package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/olofenglund96/aoc-go/helpers"
)

type Item struct {
	currVal    int
	startVal   int
	monkeyMods map[*Monkey]int
}

type Monkey struct {
	index                int
	items                []*Item
	operation            func(old int) int
	operationv2          func(old int, monkeyMod int) int
	nextMonkeyFunc       func(opResult int) int
	nextMonkeyFuncv2     func(opResult int) int
	inspections          int
	inspectionsThisRound int
	monkeyMod            int
}

func createOperationFunc(opString string) func(old int) int {
	opSplit := strings.Split(opString, " ")

	op := opSplit[1]
	opR := opSplit[2]

	if op == "+" {
		if opR == "old" {
			return func(old int) int {
				return old + old
			}
		} else {
			return func(old int) int {
				return old + helpers.StrToI(opR)
			}
		}
	} else {
		if opR == "old" {
			return func(old int) int {
				return old * old
			}
		} else {
			return func(old int) int {
				return old * helpers.StrToI(opR)
			}
		}
	}
}

func createOperationFuncv2(opString string) func(old int, monkeyMod int) int {
	opSplit := strings.Split(opString, " ")

	op := opSplit[1]
	opR := opSplit[2]

	if op == "+" {
		if opR == "old" {
			return func(old int, monkeyMod int) int {
				//shelpers.Println("[monkeyMod: ", monkeyMod, "] old + old: ", old, " + ", old, " = ", old%monkeyMod+old%monkeyMod)
				return old%monkeyMod + old%monkeyMod
			}
		} else {
			return func(old int, monkeyMod int) int {
				//helpers.Println("[monkeyMod: ", monkeyMod, "] old + ", opR, ": ", old, " + ", opR, " = ", old%monkeyMod+helpers.StrToI(opR)%monkeyMod)
				return old%monkeyMod + helpers.StrToI(opR)%monkeyMod
			}
		}
	} else {
		if opR == "old" {
			return func(old int, monkeyMod int) int {
				//helpers.Println("[monkeyMod: ", monkeyMod, "] old * old: ", old, " * ", old, " = ", old%monkeyMod*old%monkeyMod)
				return old % monkeyMod * old % monkeyMod
			}
		} else {
			return func(old int, monkeyMod int) int {
				//helpers.Println("[monkeyMod: ", monkeyMod, "] old * ", opR, ": ", old, " * ", old, " = ", old%monkeyMod*helpers.StrToI(opR)%monkeyMod)
				return old % monkeyMod * helpers.StrToI(opR) % monkeyMod
			}
		}
	}
}

func createNextMonkeyFunc(testInfo []string) (func(opResult int) int, int) {
	divRow := strings.Split(testInfo[0], " ")
	divisor := helpers.StrToI(divRow[len(divRow)-1])

	trueThrowRow := strings.Split(testInfo[1], " ")
	trueThrowMonkey := helpers.StrToI(trueThrowRow[len(trueThrowRow)-1])

	falseThrowRow := strings.Split(testInfo[2], " ")
	falseThrowMonkey := helpers.StrToI(falseThrowRow[len(falseThrowRow)-1])

	return func(opResult int) int {
		if opResult%divisor == 0 {
			return trueThrowMonkey
		} else {
			return falseThrowMonkey
		}
	}, divisor
}

func createNextMonkeyFuncv2(testInfo []string) (func(opResult int) int, int) {
	divRow := strings.Split(testInfo[0], " ")
	divisor := helpers.StrToI(divRow[len(divRow)-1])

	trueThrowRow := strings.Split(testInfo[1], " ")
	trueThrowMonkey := helpers.StrToI(trueThrowRow[len(trueThrowRow)-1])

	falseThrowRow := strings.Split(testInfo[2], " ")
	falseThrowMonkey := helpers.StrToI(falseThrowRow[len(falseThrowRow)-1])

	return func(opResult int) int {
		if opResult == 0 {
			return trueThrowMonkey
		} else {
			return falseThrowMonkey
		}
	}, divisor
}

func createMonkey(rows []string) Monkey {
	monkey := Monkey{}
	r1 := strings.Split(rows[0], " ")
	monkey.index = helpers.StrToI(strings.TrimSuffix(r1[1], ":"))

	r2 := strings.Split(rows[1], ": ")
	r2 = strings.Split(r2[1], ", ")

	l2 := helpers.StrSliceToIntSlice(r2)

	for _, v := range l2 {
		monkey.items = append(monkey.items, &Item{
			currVal:    v,
			startVal:   v,
			monkeyMods: map[*Monkey]int{},
		})
	}

	r3 := strings.Split(rows[2], "= ")
	monkey.nextMonkeyFunc, monkey.monkeyMod = createNextMonkeyFunc(rows[3:])
	monkey.operation = createOperationFunc(r3[1])
	monkey.operationv2 = createOperationFuncv2(r3[1])

	monkey.nextMonkeyFuncv2, monkey.monkeyMod = createNextMonkeyFuncv2(rows[3:])
	monkey.inspections = 0

	return monkey
}

func parseMonkeys(rows []string) []*Monkey {
	monkeys := []*Monkey{}
	items := []*Item{}
	for i := 0; i < len(rows); i += 7 {
		m := createMonkey(rows[i : i+6])
		monkeys = append(monkeys, &m)
		items = append(items, m.items...)
	}

	for _, i := range items {
		for _, m := range monkeys {
			i.monkeyMods[m] = i.currVal
		}
	}

	return monkeys
}

func (i *Item) String() string {
	prStr := fmt.Sprintf("Startval: %d, ", i.startVal)
	for m, mm := range i.monkeyMods {
		prStr += fmt.Sprintf("[%d, mod=%d]=>%+v, ", m.index, m.monkeyMod, mm)
	}

	return prStr
}

func printMonkey(m *Monkey) {
	println("----------")
	helpers.Println("Monkey ", m.index, ": numItems: ", len(m.items), ", inspections: ", m.inspections, ", monkeyMod: ", m.monkeyMod)
	fmt.Printf("Items: [")
	for _, i := range m.items {
		fmt.Printf(string(i.String()))
	}
	fmt.Printf("]")
	println("----------")

}

func printMonkeys(monkeys []*Monkey) {
	for _, m := range monkeys {
		printMonkey(m)
	}
}

func takeTurn(monkeys []*Monkey, monkey *Monkey) {
	for _, item := range monkey.items {
		//printMonkey(monkey)
		if item.startVal == 65 {
			helpers.Println("== PreMods == for item ", item.startVal)
			for _, m := range monkeys {
				fmt.Printf("[[%d]=>%+v, ", m.index, item.currVal%m.monkeyMod)
			}
			println()
			helpers.Println("==========")
		}
		wl := monkey.operation(item.currVal)
		nextMonkeyIx := monkey.nextMonkeyFunc(wl)

		if item.startVal == 65 {
			helpers.Println("== Mods == for item ", item.startVal)
			for _, m := range monkeys {
				fmt.Printf("[[%d]=>%+v, ", m.index, wl%m.monkeyMod)
			}
			println()
			helpers.Println("==========")
		}

		monkeys[nextMonkeyIx].items = append(monkeys[nextMonkeyIx].items, item)
		//helpers.Println("item: ", item, ", worry: ", wl, ", nextMonkey: ", nextMonkeyIx)
		monkey.inspections += 1
	}

	monkey.items = []*Item{}
}

func updateItemMods(monkeys []*Monkey, item *Item, wl int) {
	for _, m := range monkeys {
		item.monkeyMods[m] = wl % m.monkeyMod
	}
}

func takeTurnv2(monkeys []*Monkey, monkey *Monkey) {
	for _, item := range monkey.items {
		//printMonkey(monkey)
		if item.startVal == 1203 {
			helpers.Println("== PreMods == for item ", item.startVal)
			helpers.Println(item.String())
			helpers.Println("==========")
		}
		for m, v := range item.monkeyMods {
			item.monkeyMods[m] = monkey.operationv2(v, m.monkeyMod) % m.monkeyMod
		}
		nextMonkeyIx := monkey.nextMonkeyFuncv2(item.monkeyMods[monkey])

		if item.startVal == 1203102 {
			helpers.Println("== Mods == for item ", item.startVal)
			helpers.Println(item.String())
			helpers.Println("==========")
		}

		monkeys[nextMonkeyIx].items = append(monkeys[nextMonkeyIx].items, item)
		//helpers.Println("item: ", item, ", nextMonkey: ", nextMonkeyIx)
		monkey.inspections += 1
	}

	monkey.items = []*Item{}
}

func sol1(rows []string) string {
	monkeys := parseMonkeys(rows)

	helpers.Println("monkeys: ", monkeys)

	for i := 0; i < 20; i++ {
		helpers.Println("===== Round ", i, " =====")
		for _, monkey := range monkeys {
			takeTurn(monkeys, monkey)
		}
		printMonkeys(monkeys)
		//helpers.WaitForInput()
	}

	helpers.Println("=== Done ===")

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return fmt.Sprint(monkeys[0].inspections * monkeys[1].inspections)
}

func sol2(rows []string) string {
	monkeys := parseMonkeys(rows)

	printMonkeys(monkeys)

	for i := 0; i < 10000; i++ {
		//helpers.Println("===== Round ", i, " =====")
		for _, monkey := range monkeys {
			takeTurnv2(monkeys, monkey)
		}
		//printMonkeys(monkeys)

		//helpers.WaitForInput()
	}

	helpers.Println("=== Done ===")

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return fmt.Sprint(monkeys[0].inspections * monkeys[1].inspections)
}

func main() {
	rows := helpers.ReadFileLines(fmt.Sprintf("years/2022/11/%s.dat", os.Args[2]))

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
