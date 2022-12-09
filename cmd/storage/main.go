package main

import (
	"fmt"
	"log"

	"github.com/DeadLandscape/course_storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello world"))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)

	// file, err = st.GetFile(file.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("File found:")
	// fmt.Println(file)
}
