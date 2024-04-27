package date

import (
	"errors"
	"fmt"
	"os/exec"
)

const defaultFormat string = "+%A %Y-%m-%dT%H:%M:%S%Z"

func GenerateDate(args []string) ([]byte, error) {
	// cmd := exec.Command("date", formatArg)
	if len(args) == 0 {
		return []byte{}, errors.New("no args")
	}
	cmd := exec.Command("date", args...)
	return cmd.Output()
}

func printDate(args []string) {
	// fmt.Println(format)
	out, err := GenerateDate(args)
	if err != nil {
		// fmt.Println(err)
		out, _ = GenerateDate([]string{"-u", defaultFormat})
		fmt.Print(string(out))
		return
	}
	fmt.Print(string(out))
}

func Run(args []string) {
	printDate(args)
}
