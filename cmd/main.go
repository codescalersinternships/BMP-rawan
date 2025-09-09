package main

import (
	bmp "github.com/codescalersinternships/BMP-rawan/pkg"
)

func main() {
	w,h,err:= bmp.ExtractWH("testdata/sample1.bmp")
	if err!=nil{
		panic(err)
	}
	println("Width:",w,"Height:",h)
}
