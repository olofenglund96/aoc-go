package cmd

import "fmt"

func get_day_and_year() (int, int) {
	year, err := dayCmd.PersistentFlags().GetInt("year")
	if err != nil {
		panic(fmt.Errorf("Could not find year: %e", err))
	}

	day, err := dayCmd.PersistentFlags().GetInt("day")
	if err != nil {
		panic(fmt.Errorf("Could not find day: %e", err))
	}

	return year, day
}
