package main

// make sure to use `go run .` `not go run main.go` to run this code
func main() {
	file1 := &File{
		Name: "File1",
		Byte: "rose",
	}
	file2 := &File{
		Name: "File2",
		Byte: "rasa",
	}
	file3 := &File{
		Name: "File3",
		Byte: "roso",
	}

	folder1 := &Folder{
		Name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		Name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.Search("rose")
}
