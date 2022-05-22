package main

import "errors"

type HashMap struct {
	m map[int]string
}

func NewHashMap() *HashMap {
	return &HashMap{
		m: make(map[int]string),
	}
}

func (h *HashMap) Get(key int) (string, error) {
	val, ok := h.m[key]
	if !ok {
		return "", errors.New("not found")
	}
	return val, nil
}

func (h *HashMap) Put(key int, value string) error {
	if _, ok := h.m[key]; ok {
		return errors.New("key exists in the hashmap")
	}
	h.m[key] = value
	return nil
}

func (h *HashMap) GetRange(from, to int) ([]string, error) {
	res := make([]string, 0)
	for i := from; i <= to; i++ {
		val, ok := h.m[i]
		if ok {
			res = append(res, val)
		}
	}

	if len(res) < 1 {
		return nil, errors.New("no data")
	}
	return res, nil // TODO: replace this
}
