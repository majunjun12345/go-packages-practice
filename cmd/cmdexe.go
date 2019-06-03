package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	/*
		可以先检查命令，再执行命令
	*/
	lsPath, _ := exec.LookPath("ls")
	fmt.Println(lsPath)

	/*
		进行相关配置后，通过 run 执行命令
	*/
	cmd := exec.Command("ls")
	out := new(bytes.Buffer)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out.String())

	cmd1 := exec.Command("pwd")
	var stdout, stderr bytes.Buffer
	cmd1.Stderr = &stderr
	cmd1.Stdout = &stdout
	cmd1.Run()
	fmt.Println(stdout.String(), stderr.String())

	/*
		通过 output 执行单一命令
		通过 CombinedOutput 执行组合命令
	*/
	result, _ := exec.Command("ls").Output()
	fmt.Println(string(result))
	result, _ = exec.Command("ls", "-ahl").CombinedOutput()
	fmt.Println(string(result))
}
