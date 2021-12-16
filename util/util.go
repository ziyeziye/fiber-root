package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ExecCmd(command string) {
	shell := "sh"
	if runtime.GOOS == "windows" {
		shell = "powershell"
	}

	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(shell, "-c", command)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// 因为结果是字节数组，需要转换成string
	fmt.Println(string(output))
}

func PrintJson(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatalln(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	out.WriteTo(os.Stdout)
}
