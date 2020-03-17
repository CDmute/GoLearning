package map

import (
	"fmt"

	"reflect"
)

func MapEquals(){
	m1:=make(map[string]string)
	m2:=make(map[string]string)
	m3:=m1

	m1["key"]="value"
	m2["key"]="value"

	fmt.Println(fmt.Sprintf("Is m1 equals to m2? %v",&m1==&m2))
	fmt.Println(fmt.Sprintf("Is m1 equals to m3? %v",&m1==&m3))
	fmt.Println(fmt.Sprintf("Is m1 deeply equals to m3? %v",reflect.DeepEqual(m1,m2)))
}