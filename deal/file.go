package deal

import (
	"bufio"
	"os"
	"path/filepath"
)

func OpenFileSmart(path string) *os.File {
	f,e:=os.OpenFile(path,os.O_RDWR,0777)
	if e != nil && os.IsNotExist(e){
		newf,_:=os.Create(path)
		return newf
	}else {
		return f
	}
}


func Is_contain(target string) bool {
	_, err := os.Stat(target)
	if err != nil {
		return false
	} else {
		return true
	}
}





func Get_path() (full_path string, path string) {
	currentPath := filepath.Dir(os.Args[0])
	//包含文件名的完整目录
	//fmt.Println(os.Args[0])
	full_path = os.Args[0]
	//不包含文件名
	//fmt.Println(currentPath)
	path = currentPath
	return
}

func Write_bat(taget string, bat_lines []string) {
	f, _ := os.OpenFile(taget, os.O_RDWR|os.O_CREATE, 777)
	w := bufio.NewWriter(f)
	for _, line := range bat_lines {
		w.WriteString(line + "\r\n")
	}
	w.Flush()
}
