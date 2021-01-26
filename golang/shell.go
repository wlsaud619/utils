package golang

import (
	"fmt"
	"os/exec"
)

func SHELL(cmd string) ([]uint8, error) {
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
	fmt.Println(cmd)
	return out, err

}
