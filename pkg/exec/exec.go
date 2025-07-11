package exec

import (
	"fmt"
	"os/exec"
)

func ExecuteCommand(args ...string) {
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}
