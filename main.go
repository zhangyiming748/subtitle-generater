package main

import (
	"log"
	"os"
	"subtitle-generater/generater"
)

func main() {
	var (
		s, e, o string
	)

	if len(os.Args) > 1 {
		s = os.Args[1]
		e = os.Args[2]
		o = os.Args[3]
	} else {
		s = generater.GetVal("subtitle", "subject")
		e = generater.GetVal("subtitle", "event")
		o = generater.GetVal("subtitle", "otherword")
	}

	ret := generater.MarketingGenerater(s, e, o)
	writeAll("subtitle.txt", ret)

}
func writeAll(fname, content string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	n, err := f.WriteString(content)
	if err != nil {
		log.Println("写文件出错")
	} else {
		log.Printf("写入%d个字节", n)
	}
}
