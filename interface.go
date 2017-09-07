package main

import "fmt"

//接口把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
type Student interface {
	hello()
}

type MiddleStudent struct {
	id   int
	name string
}

type CollegeStudent struct {
	id   int
	name string
}

func (middleStudent MiddleStudent) hello() {
	fmt.Println("I'm middle student", middleStudent.id, middleStudent.name)
}

func (collegeStudent CollegeStudent) hello() {
	fmt.Println("I'm college student", collegeStudent.id, collegeStudent.name)
}

//Go中没有构造函数的概念,对象的创建通常交由一个全局的创建函数来完成,以NewXXX来命令,表示"构造函数"
func NewMiddleStudent(id int, name string) *MiddleStudent {
	//return &MiddleStudent{id, name}//按顺序传递参数
	return &MiddleStudent{id: id, name: name} //按属性名传递参数
}

func NewCollegeStudent(id int, name string) *CollegeStudent {
	//return &CollegeStudent{id, name}//按顺序传递参数
	return &CollegeStudent{id: id, name: name} //按属性名传递参数
}

func main() {

	var pstu Student
	pstu = new(MiddleStudent) //pstu是指针变量，new(MiddleStudent) 分配了零值填充的MiddleStudent类型的内存空间，并且返回其地址，一个*MiddleStudent类型的值
	pstu.hello()
	pstu = new(CollegeStudent)
	pstu.hello()
	fmt.Println(pstu)

	pstu = &MiddleStudent{} //等价于new
	pstu.hello()
	pstu = &CollegeStudent{}
	pstu.hello()
	fmt.Println(pstu)

	var stu Student
	stu = MiddleStudent{} //stu是MiddleStudent类型的变量
	stu.hello()
	stu = CollegeStudent{}
	stu.hello()
	fmt.Println(stu)

	//通过全局构造函数创建
	stu1 := NewMiddleStudent(1, "好好先森")
	stu1.hello()
	fmt.Println(stu1)
	stu2 := NewCollegeStudent(2, "好好先森V5")
	stu2.hello()
	fmt.Println(stu2)

}
