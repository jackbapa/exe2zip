package deal

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

var TestUrl string =   "https://repo.anaconda.com/archive/Anaconda3-2020.11-Windows-x86_64.exe"
var TestUrl2 = "http://zjk-rfile.checkpass.net/cnkivip/2021-02-24/21439709464944958176_%E7%8E%8B%E8%80%80_%E5%9F%BA%E4%BA%8E%E8%B7%AF%E5%BE%84%E8%B7%9F%E8%B8%AA%E7%9A%84%E6%99%BA%E8%83%BD%E5%88%86%E5%B8%83%E5%BC%8F%E9%A9%B1%E5%8A%A8%E8%BD%A6%E8%BE%86%E8%BF%90%E5%8A%A8%E7%A8%B3%E5%AE%9A%E6%80%A7%E6%8E%A7%E5%88%B6%E7%AD%96%E7%95%A5%E7%A0%94%E7%A9%B6.zip"


func Deal_con(newNet net.Conn)  {
	p := Praser{}
	p.Prase_BytePackge(newNet)
	switch p.Acction {
	case "git clone":
		//GitClone(newNet,p.Url)
	}

}



func Download_from_socket(path string,con net.Conn,fileName string)  {
	f := OpenFileSmart(path+"\\"+fileName)
	defer f.Close()

	resp_reader := bufio.NewReader(con)
	onceRead:=make([]byte,1024)
	recent_time:=time.Now().UnixNano()
	i := 0
	for {
		n, err := resp_reader.Read(onceRead)
		if err == io.EOF || n == 0 {
			break
		} else {
			//fmt.Println(n)
			fmt.Print("=>")
			i+=1
			if i==10{
				now_time:=time.Now().UnixNano()
				dur := (float64(now_time-recent_time)/1000000000)
				speed := (10/dur)/1000.0
				fmt.Println(strconv.FormatFloat(speed,'f',2,64)+" Mb/s")
				i = 0
				recent_time = time.Now().UnixNano()
			}
			fmt.Print("=>>")
			f.Write(onceRead[0:n])

		}
	}

}
