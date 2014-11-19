package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
	"net/url"
	"strconv"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Println("valvv:", v)
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL) //获取请求的方法
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println(r.Form["password"][0])
	}
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	fmt.Println(v.Encode())
	getint,err:=strconv.Atoi(r.Form.Get("age"))
	if err!=nil{
		//数字转化出错了，那么可能就不是数字
	}

	//接下来就可以判断这个数字的大小范围了
	if getint >100 {
		//太大了
	}

}

func select01(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL) //获取请求的方法
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/select01.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		slice:=[]string{"apple","pear","banane"}

		for _, v := range slice {
			if v == r.Form.Get("fruit") {
				fmt.Fprintf(w,"true")
				return
			}
		}
		fmt.Fprintf(w,"false")
	}

}

func main() {
	http.HandleFunc("/", sayHelloName) 		 //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	http.HandleFunc("/select01", select01)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
