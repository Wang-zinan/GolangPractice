package algorithm

import (
	"bufio"
	"fmt"
	"os"
)

func ReadIn(){
	read := bufio.NewReader(os.Stdin)
	fmt.Println("获取到的输入为：")
	str1,err:= read.ReadString('\n')
	if err != nil{
		panic(err)
	}
	fmt.Printf("打印出来的结果未：%s",str1)
}