package main

import "fmt"

// Nah Folder ini memiliki sebuah interface `Component` didalamnya `Component` bisa saja adalah sebuah `Folder` atau `File`

type Folder struct {
	Component []Component
	Name      string
}

// Nah Folder juga mengimplement interface `Component` karena sebuah folder bisa saja memiliki `Folder` lagi didalamnya
func (f *Folder) Search(keyword string) bool {
	fmt.Printf("Serching recursively for keyword %s in folder: %s\n", keyword, f.Name)
	for _, composite := range f.Component {
		composite.Search(keyword)
	}
	return true
}

// Simple helper function to add component to a folder
func (f *Folder) add(c Component) {
	f.Component = append(f.Component, c)
}
