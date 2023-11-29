package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/atotto/clipboard"
)

func readClipboard() (string, error) {
	content, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}
	return content, nil
}

func getCurrentDay() string {
	currentTime := time.Now()
	day := currentTime.Format("02")
	return day
}

func main() {
	day := getCurrentDay()
	args := os.Args
	if len(args) > 1 {
		dayArg := args[1]
		if len(dayArg) == 1 {
			dayArg = "0" + dayArg
		}
		day = dayArg
	}

	err := os.Mkdir(day, 0755)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return
	}

	input, err := readClipboard()
	if err != nil {
		fmt.Println("Failed to read clipboard:", err)
		return
	}	
	
	filePath := day + "/input.txt"
	err = os.WriteFile(filePath, []byte(input), 0644)
	if err != nil {
		fmt.Println("Failed to write input to file:", err)
		return
	}

	content := `package main
	
	import (
		"fmt"
	)
	
	func main() {
		
	}`
	filePath = day + "/run.go"
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Failed to write go file:", err)
		return
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && go mod init aoc2023/day%s", day, day))
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to initialize go module:", err)
		return
	}
}
