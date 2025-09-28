package main

import (
	"fmt"
	"strconv"
)

type StringMap struct {
	sm map[string]string 
}

func NewStringMap() *StringMap {
	return &StringMap{
		sm: make(map[string]string),
	}
}

func (jm *StringMap) Get(k string) string {
	return jm.sm[k]
}

func (jm *StringMap) Add(k, v string) {
	jm.sm[k] = v
}

type IntToStringMapAdapter struct {
	stringMap *StringMap
}

func NewIntToStringMapAdapter() *IntToStringMapAdapter {
	return &IntToStringMapAdapter{
		stringMap: NewStringMap(),
	}
}

func (adapter *IntToStringMapAdapter) Get(k int) string {
	return adapter.stringMap.Get(strconv.Itoa(k))
}

func (jm *IntToStringMapAdapter) Add(k, v int) {
	jm.stringMap.Add(strconv.Itoa(k), strconv.Itoa(v))
}

func main() {
	adapter := NewIntToStringMapAdapter()

	adapter.Add(5, 10)
	v := adapter.Get(5)

	fmt.Printf("Got value %v of string type", v)
}
