package deal

import (
	"fmt"
	"os/user"
	"strings"
	"syscall"
	"unsafe"
)
func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func Isadmin(path string) {
	//file_op.Write_bat(file_op.Start_menu+"1.text",[]string{"ddddd","d123"})
	ddl,e := syscall.LoadLibrary("shell32")
	defer syscall.FreeLibrary(ddl)
	if e != nil{
		fmt.Println(e)
	}
	myfunc,e_f:=syscall.GetProcAddress(ddl,"IsUserAnAdmin")
	if e_f != nil{
		fmt.Println("e-f fail  IsUserAnAdmin")
		fmt.Println(e_f)
	}
	r1,_,e_f_c := syscall.Syscall(myfunc,0, 0, 0, 0)
	//erron不是error
	if e_f_c != 0{
		fmt.Println("调用失败")
		fmt.Println(e_f_c)
		panic(e_f_c)
	}else {
		//return r1 != 0
	}
	fmt.Println(r1)
	//不是管理员
	if r1 == 0{
		fmt.Println("--------")
		shell_exe,shell_exe_e:=syscall.GetProcAddress(ddl,"ShellExecuteW")
		if shell_exe_e != nil {
			fmt.Println("get_shell_exe fail   ")
			fmt.Println(shell_exe_e)
		}
		ret,_,eee:=syscall.Syscall6(shell_exe,0,StrPtr("runas"),
			0,
			StrPtr(path),
			0,1,0)
		if eee == 0{
			fmt.Println("***success zhixing")
			fmt.Println(ret)
		}

	}else {
		fmt.Println("we get you")
		fmt.Println(r1)
	}
}

func GetUserName() string {
	compterUser,_ :=  user.Current()
	index := strings.Index(compterUser.Username, "\\")
	if index == -1{
		return compterUser.Username
	}
	temp := compterUser.Username[index+1:]
	return temp

}
