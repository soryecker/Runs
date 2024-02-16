package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const help = `
Usage: runs <textfile> <command>

Iterate through each line of the text file and execute it in the command.

Parameters:
 	<textfile>  
		The path to the text file that will be processed. This file should exist on your system and be readable.
	<command>   
		The command to execute after processing the text file. This must be a recognized command that the program can perform.

Example:
	runs domain.txt "assetfinder -subs-only {*}"

	Iterate through each line in domain.txt, fill it into "{*}", and then execute.

Notes:
  - Ensure that the path to the text file is correct and the file is accessible.
  - The command should be a valid operation that this program is designed to handle.
  - Review the available commands and their syntax if necessary.

varsion:
	0.0.1 - 面包狗
`

func ExecuteCommand(command string) error {
	// 使用exec.Command创建一个表示外部命令的*Cmd
	cmd := exec.Command("cmd", "/C", command)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// 开始执行命令
	if err := cmd.Start(); err != nil {
		return err
	}

	// 创建一个新的bufio.Reader，从标准输出管道读取数据
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // 实时打印命令的输出
	}

	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func run() {
	txtPath := os.Args[1] // 第二个参数
	comd := os.Args[2]    // 第二个参数

	file, err := os.Open(txtPath)
	if err != nil {
		log.Fatalf("File opening failed: %v", err)
		return
	}
	defer file.Close()

	// 创建文件的缓冲读取器
	scanner := bufio.NewScanner(file)

	// 遍历文件的每一行
	for scanner.Scan() {
		// 获取当前行的内容
		line := scanner.Text()
		//每一行的数据
		line = strings.ReplaceAll(line, " ", "")
		comm := strings.ReplaceAll(comd, "{*}", line)
		//fmt.Println(line)
		err := ExecuteCommand(comm)
		if err != nil {
			return
		}
	}

	// 检查遍历过程中是否有错误发生
	if err := scanner.Err(); err != nil {
		log.Fatalf("An error occurred while reading the file: %v", err)
	}
}

func main() {

	helpWanted := false

	for _, arg := range os.Args[1:] { // 跳过第0个元素，从第1个开始遍历
		if arg == "-h" {
			helpWanted = true
			break // 找到 -h 后停止循环
		} else if arg == "" {
			helpWanted = true
		}
	}

	if helpWanted {
		fmt.Println(help)
	}
	if len(os.Args) < 2 {
		fmt.Println(help)
	} else {
		run()
	}

}
