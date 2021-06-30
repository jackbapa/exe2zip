package conf

import "os"

var TempPath = "D:\\speedup\\"

func Temp_path()  {
	_, err := os.Stat(TempPath)
	if os.IsNotExist(err) {
		os.Mkdir(TempPath,0777)
	} else {

	}
}

func Readconf()string  {

	return "107.173.149.35:8086"
}
