package system

import (
	"os/exec"
)

/*
File name    : system.go
Author       : miaoyc
Create date  : 2021/12/2 4:34 下午
Description  : 操作系统调用相关
*/

func RunSystemShell(name, args, shell string) (output string, err error) {
	cmd := exec.Command(name, args, shell)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(outputBytes), nil
}
