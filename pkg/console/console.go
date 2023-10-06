package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Success 打印成功信息，绿色输出
func Success(msg string) {
	output(msg, "green")
}

// Error 打印错误信息，红色输出
func Error(msg string) {
	output(msg, "red")
}

// ErrorIf 语法糖，自带 err != nil 判断
func ErrorIf(err error) {
	if err != nil {
		Error(err.Error())
	}
}

// Warning 打印一条提示消息，黄色输出
func Warning(msg string) {
	output(msg, "yellow")
}

func WarningIn(err error) {
	Warning(err.Error())
}

// Exit 打印错误信息，并退出 os.Exit(1)
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func output(msg, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}
