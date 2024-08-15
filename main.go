package main

import (
	"autotiktok/myservice"
)


func main() {
	flist := myservice.Collection()
	for len(flist) > 0 {
		flist = myservice.Jump(flist)
	}

}