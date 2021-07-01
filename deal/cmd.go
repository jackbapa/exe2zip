package deal

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func CmdDo(command string) string {
	cmd := exec.Command("cmd.exe")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in

	in.WriteString(command)
	stdout, err := cmd.StdoutPipe();
	if  err != nil {     //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	defer stdout.Close()   // 保证关闭输出流
	starterr := cmd.Start()
	if starterr != nil {   // 运行命令
		fmt.Println(err)
	}
	return out.String()

}

func NotepadDo(command string) string {
	//exec.Command("cmd.exe")是打开一个软件，并与之用in交互，而notepad不支持打开后的交互参数
	//所以不能使用CmdDo("notepad.exe")直接打开
	cmd := exec.Command("notepad.exe",command)
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out

	starterr := cmd.Start()
	if starterr != nil {   // 运行命令
		fmt.Println(starterr)
	}
	return out.String()


}

func CmdDoRaw(exe string,args []string,command string) string {
	cmd := exec.Command(exe,args...)
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in

	in.WriteString(command)
	stdout, err := cmd.StdoutPipe();
	if  err != nil {     //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	defer stdout.Close()   // 保证关闭输出流
	starterr := cmd.Start()
	if starterr != nil {   // 运行命令
		//log.Fatal(starterr)
		fmt.Println(starterr)
	}
	return out.String()

}