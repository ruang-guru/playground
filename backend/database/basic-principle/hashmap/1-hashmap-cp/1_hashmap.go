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
	return nil, nil // TODO: replace this
}
