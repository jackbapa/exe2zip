package deal

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

func insertKey(regeditkeyPath string,user bool,keyValue map[string]string,)  {
	var key registry.Key
	var exists bool
	if user{
		key, exists, _ = registry.CreateKey(registry.CURRENT_USER,
			regeditkeyPath, registry.ALL_ACCESS)
	}else {
		key, exists, _ = registry.CreateKey(registry.CLASSES_ROOT,
			regeditkeyPath, registry.ALL_ACCESS)
	}
	if exists {
		fmt.Println("key already exit")
	}
	defer key.Close()

	for k,v := range keyValue{
		key.SetStringValue(k, v)
	}
}

func addkey(regeditkey string,user bool,keyValue map[string]string,)  {
	var key registry.Key
	var e error
	if user{
		key, e= registry.OpenKey(registry.CURRENT_USER,
			regeditkey, registry.ALL_ACCESS)
	}else {
		key, e= registry.OpenKey(registry.CURRENT_USER,
			regeditkey, registry.ALL_ACCESS)

	}
	if e != nil{
		fmt.Println(e)
	}
	for k,v := range keyValue{
		key.SetStringValue(k, v)
	}
	for k,v := range keyValue{
		key.SetStringValue(k, v)
	}

}


