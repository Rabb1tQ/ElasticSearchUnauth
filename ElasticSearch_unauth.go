package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func ScanUnauth(addressTemp string) (result bool, err error) {
	client := &http.Client{}
	//生成要访问的url
	//提交请求
	reqest, err := http.NewRequest("GET","http://"+ strings.Replace(addressTemp, "\r\n", "", -1) +"/_cat", nil)
	if err != nil {
		panic(err)
	}
	response, _ := client.Do(reqest)

	str_byte,_:=ioutil.ReadAll(response.Body)
	str:=string(str_byte)
	status := response.StatusCode
	if status == 200 {
		if strings.Contains(str,"/_cat/master") {
			fmt.Print(addressTemp)
		}
	}

	return result, err
}
func banner() {
	print("██████╗ ██╗   ██╗    ██████╗  █████╗ ██████╗ ██████╗ ██╗████████╗ ██████╗\n██╔══██╗╚██╗ ██╔╝    ██╔══██╗██╔══██╗██╔══██╗██╔══██╗██║╚══██╔══╝██╔═══██╗\n██████╔╝ ╚████╔╝     ██████╔╝███████║██████╔╝██████╔╝██║   ██║   ██║   ██║\n██╔══██╗  ╚██╔╝      ██╔══██╗██╔══██║██╔══██╗██╔══██╗██║   ██║   ██║▄▄ ██║\n██████╔╝   ██║       ██║  ██║██║  ██║██████╔╝██████╔╝██║   ██║   ╚██████╔╝\n╚═════╝    ╚═╝       ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ ╚═════╝ ╚═╝   ╚═╝    ╚══▀▀═╝ ")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

var (
	h         = flag.Bool("h", false, "帮助信息")
	run         = flag.Bool("run", false, "")
)

func main() {
	banner()
	fmt.Println("开始时间: ",time.Now().Format("2006/01/02 15:04"))
	flag.Parse()
	if *h == true ||*run==false {
		fmt.Println("Usage:同级目录下放置address.txt，然后执行 ElasticSearch_unauth run")
		return
	}
	//ip := *ip
	address, addressrerr := os.Open("address.txt")
	if addressrerr != nil {
		fmt.Println(addressrerr.Error())
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(address)
	for {
		//遇到\n结束读取
		addressTemp, errR := buf.ReadBytes('\n')
		if len(addressTemp) > 0 {
			ScanUnauth(string(addressTemp))
		}
		if errR != nil {
			if errR == io.EOF {
				break
			}
			//fmt.Println(errR.Error())
		}
	}
	fmt.Print("检测完成, 结束时间：%v",time.Now())
}
