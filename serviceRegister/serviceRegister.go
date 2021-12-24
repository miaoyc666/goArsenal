package serviceRegister

import "fmt"

/*
File name    : serviceRegister.go
Author       : miaoyc
Create date  : 2021/12/24 2:12 下午
Description  :
*/

type daoTemplate interface {
	getA(i string) string
	getB(i string) string
}

type myDao struct {
	nameA string
	nameB string
}

func (m *myDao) getA(i string) string {
	return m.nameA + i
}

func (m *myDao) getB(i string) string {
	return m.nameB + i
}

func query(dao daoTemplate) {
	i := "test"
	fmt.Println(dao.getA(i))
	fmt.Println(dao.getB(i))
}

func main() {
	testDao := myDao{nameA: "nameA", nameB: "nameB"}
	var dao daoTemplate = &testDao
	query(dao)
}