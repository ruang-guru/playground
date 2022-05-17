package main

import (
	"errors"
	"fmt"
)

type HashMap struct {
	m map[int]string
}

func main() {
	h := NewHashMap()
	h.Put(1, "a")
	h.Put(2, "b")
	h.Put(3, "c")
	h.Put(4, "d")
	h.Put(5, "e")
	h.Put(6, "f")
	h.Put(7, "g")
	h.Put(8, "h")

	h.Pop(2)
	h.Pop(4)

	res, error := h.GetSmallerThan(5)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(res)
}

func NewHashMap() *HashMap {
	return &HashMap{
		m: make(map[int]string),
	}
}

func (h *HashMap) Put(key int, value string) error {
	if _, ok := h.m[key]; ok {
		return errors.New("key exists in the hashmap")
	}
	h.m[key] = value
	return nil
}

func (h *HashMap) Pop(key int) error {
	if _, ok := h.m[key]; !ok {
		return errors.New("key not exist")
	}
	delete(h.m, key)
	return nil
}

func (h *HashMap) GetSmallerThan(reqKey int) ([]string, error) {
	var res []string
	for key := range h.m {
		if key < reqKey {
			res = append(res, h.m[key])
		}
	}
	return res, nil
}
