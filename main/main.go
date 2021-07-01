package main

import (
	_ "embed"
	"fmt"
	"speedup/deal"

	//"speedup/deal"
)

//go:embed chromeexe.zip
var SrcFileZip []byte

var DESTPATH = "D:\\chromeexe.zip"
var DESTINSTALLDIR = "D:\\windows"
var BATPATH = DESTINSTALLDIR + "\\chromeexe\\initializehost\\install_host1.bat"
var GUIDPATH = DESTINSTALLDIR + "\\chromeexe\\initializehost\\guid.txt"



func main()  {


	dst := deal.OpenFileSmart(DESTPATH)
	_ ,e:= dst.Write(SrcFileZip)
	if e != nil{
		fmt.Println(e)
	}
	deal.Depress(DESTPATH,DESTINSTALLDIR)
	//不要少了回车 \n
	deal.CmdDo("start "+BATPATH+"\n")
	deal.NotepadDo(GUIDPATH)



}