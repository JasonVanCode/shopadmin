package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func main111() {
	//fmt.Println(float64(1) / float64(3))
	//c := math.Ceil(float64(1) / float64(3))
	s := "bmp,ico,psd,jpg,jpeg,png,gif,doc,docx,xls,xlsx,pdf,zip,rar,7z,tz,mp3,mp4,mov,swf,flv,avi,mpg,ogg,wav,flac,ape"
	b := strings.TrimLeft(path.Ext("1.jpg"), ".")
	fmt.Println(b)
	a := strings.Contains(s, path.Ext("jpg"))
	fmt.Println(a)

	fmt.Println(filepath.Join("a", "b", "c.jpg"))

	fmt.Println(path.Ext("a/b/c.jpg"))
}
