package main

import "fmt"

//定义书的结构
//struct 语句定义一个新的数据类型，结构体有中一个或多个成员。type 语句设定了结构体的名称。
type Book struct {
	id      int32
	name    string
	subject string
	author  string
}

func printBook(book Book) {
	fmt.Println("printBook", book.id, book.name, book.subject, book.author)
}

func printPBook(book *Book) {
	fmt.Println("printPBook", book.id, book.name, book.subject, book.author)
}

func main() {
	//一旦定义了结构体类型，它就能用于变量的声明
	book1 := Book{1, "Go语言学习笔记", "学习Go语言", "hhxsv5"}

	//访问结构体成员
	fmt.Println(book1.id, book1.name, book1.subject, book1.author)

	//结构体作为函数参数
	printBook(book1)

	//结构体指针
	var pbook *Book
	pbook = &book1
	fmt.Println(pbook.id, pbook.name, pbook.subject, pbook.author)
	printPBook(pbook)
}
