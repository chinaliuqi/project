/*
一个HTTP SERVER 读取txt样例
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {

	file, err := os.Open(`/Users/liuqi/go/src/project/test/web/fileweb.txt`)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		} else {
			fmt.Println(line)
			//fmt.Fprint(writer,line,request.URL.Path[1:])
			fmt.Fprintf(writer, line)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
