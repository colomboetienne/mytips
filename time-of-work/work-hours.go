package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

// On doit ajouter \n pour que le parsing se fasse ...
// TODO remove \n
const HOUR_FORMAT = "15:04"

func main() {

	nb_hours := time.Duration((7*time.Hour + 40*time.Minute))

	arrived := readTime("At what time did you arrived (%H:%M): ")
	fmt.Println(arrived)

	nb_lunch := readTimeDelta("How long did your lunch last (%Hh%Mm): ")

	fmt.Println(nb_hours)
	fmt.Println(nb_lunch)


}

func readTimeDelta(input_message string) *time.Duration {

	text := askInput(input_message)

	lunch_time, err := time.ParseDuration(text)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &lunch_time
}

func readTime(input_message string) *time.Time {

	text := askInput(input_message)

	arrived, err := time.Parse(HOUR_FORMAT, text)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &arrived
}

func askInput(input_message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input_message)
	text, _ := reader.ReadString('\n')
	re := regexp.MustCompile("\n")
    fmt.Println(re.ReplaceAllString(text, ""))
	// strings.Replace(text, "\n","",-1)

	return text

}
