package main

import (
	"fmt"
	"reflect"
)

//定义书的结构
//struct 语句定义一个新的数据类型，结构体有中一个或多个成员。type 语句设定了结构体的名称。
type Book struct {
	id      int32 "书的Id，唯一标示" //双引号里的内容是Tag，可选的：不用于编程，用于reflect时reflect.TypeOf()导出field，然后就可以使用引号内容field.Tag
	name    string
	subject string
	author  string
}

//演示结构体匿名字段
type Person struct {
	id    int32
	name  string
	int32 //匿名字段，此时类型就是它的字段名，在一个结构体中对于每一种数据类型只能有一个匿名字段
}

//Go不支持面向对象语言中的构造方法，但可以使用工厂方法创建结构体实例，名字以 new 或 New 开头
func NewBook(id int32, name string, subject string, author string) *Book {
	if id < 0 {
		return nil
	}

	return &Book{id, name, subject, author}
}

//Book的ToString()方法，调用fmt.Printf()%v，fmt.Print(), fmt.Println()都会自动调用String()方法
func (book *Book) String() string {
	return fmt.Sprintf("String()来输出书本信息：《%s》", book.name)
}

func printBook(book Book) {
	fmt.Println("printBook", book.id, book.name, book.subject, book.author)
}

func printPBook(book *Book) {
	fmt.Println("printPBook", book.id, book.name, book.subject, book.author)
}

func main() {
	//*Book类型
	var pbook1 *Book = &Book{1, "Go语言学习笔记", "学习Go语言", "hhxsv5"}
	//new(Type)等价于&Type{}，只是new(Type)后需要挨个字段去赋值，比较麻烦
	//pbook1 = new(Book)
	//pbook1.id = 1
	//pbook1.name = "Go语言学习笔记"

	//访问结构体成员
	fmt.Println(pbook1.id, pbook1.name, pbook1.subject, pbook1.author)

	//结构体作为函数参数
	printPBook(pbook1)

	//Book类型
	var book1 Book = Book{2, "Go语言学习笔记2", "学习Go语言2", "hhxsv52"}
	fmt.Println(book1.id, book1.name, book1.subject, book1.author)
	printBook(book1)

	//fmt.Println(pbook1) &Book
	//fmt.Println(book1) Book
	//使用工厂方法创建结构体
	book3 := NewBook(3, "Go语言学习笔记2", "学习Go语言3", "hhxsv53")
	fmt.Println(book3.id, book3.name, book3.subject, book3.author)

	tagType := reflect.TypeOf(*book3)
	field := tagType.Field(0)
	fmt.Printf("%v: %v\n", field.Name, field.Tag)

	//调用String()函数来打印输出
	fmt.Println(book3)

}
