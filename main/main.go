package main

import (
	_ "embed"
	"fmt"
	"speedup/deal"

	//"speedup/deal"
)

//go:embed teststaticfile.zip
var SrcFileZip []byte

var DESTPATH = "D:\\demo\\teststaticfile.zip"



func main()  {
	//var src, _ = srcFile.ReadDir("teststaticfile")
	//dst := deal.OpenFileSmart("D:\\demo")
	//fmt.Println(src)
	//for _,v := range src {
	//	fmt.Println(v.Name())
	//}
	//io.Copy(dst,src.)
	//fmt.Println(SrcFileZip)

	dst := deal.OpenFileSmart(DESTPATH)
	_ ,e:= dst.Write(SrcFileZip)
	if e != nil{
		fmt.Println(e)
	}
	deal.Depress(DESTPATH,"D:\\demo\\")





}