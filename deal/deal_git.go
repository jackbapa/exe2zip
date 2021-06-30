package deal

import (
	"fmt"
	"net"
	"os"
	"speedup/conf"
	"strings"
)

func GitClone(url string)  {
	newUrl := strings.Replace(url,".git","/archive/refs/heads/master.zip",-1)
	ip_port := conf.Readconf()
	p:=Praser{
		Acction: "git clone",
		Url: newUrl,
	}
	fmt.Println("======正在连接加速服务器======")
	con,e := net.Dial("tcp",ip_port)
	if e != nil{
		fmt.Println(e)
		fmt.Println("======加速服务器连接失败=====")
	}else {
		header,body := p.Get_BytePackge()
		data:=append(header,body...)
		con.Write(data)
		path,_:= os.Getwd()
		fmt.Println("==================")
		fmt.Println(path)
		fmt.Println("==================")
		Download_from_socket(path,con,"master.zip")
	}
}

func Download(url string)  {
	ip_port := conf.Readconf()

	p:=Praser{
		Acction: "download",
		Url: url,
	}
	fmt.Println("======正在连接加速服务器======")
	con,e := net.Dial("tcp",ip_port)
	if e != nil{
		fmt.Println(e)
		fmt.Println("======加速服务器连接失败=====")
	}else {
		header,body := p.Get_BytePackge()
		data:=append(header,body...)
		con.Write(data)
		path,_:= os.Getwd()
		fmt.Println("==================")
		fmt.Println(path)
		fmt.Println("==================")
		str_array_temp:=strings.Split(url,"/")
		file_name := str_array_temp[len(str_array_temp)-1]
		Download_from_socket(path,con,file_name)
	}

}
