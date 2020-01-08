package main

import (
	"sync"
	"fmt"
)

type GameObject struct {
	IsActivity bool
	ObjId   int64
}

func modiMap(modi map[int]string)  {
	modi[3] = "潇潇洒洒第三代"
}

func main() {

	normalmap :=  make(map[int] string)
	map2 := normalmap


	var sm sync.Map

	if gameObj,ok := sm.LoadOrStore(2,&GameObject{true,996});ok{
		fmt.Println("%v",gameObj)
	}
	sm.Delete(2)
	gameObj,ok :=sm.Load(3)
	fmt.Println(gameObj == nil,"res =",ok)
	normalmap[1] = "小红"
	normalmap[5] = "小红"
	normalmap[9999] = "小红"
	a := normalmap[888]
	fmt.Println("a =", len(a))
	modiMap(normalmap)
	for i,k := range normalmap  {
		fmt.Println(i,"=",k)
	}

	for i,k := range map2  {
		fmt.Println(i,"=",k)
	}
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key,"=",value)
		return true
	})
}
