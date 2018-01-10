package main

import (
	"fmt"
	"strings"
	"os"
	"path/filepath"
	"github.com/kataras/iris/core/errors"
	"container/list"
	"sort"
)

func ls(p string) error {
	//Open dir and read
	if dir, err := os.Open(p); err == nil {
		if fileInfos, err := dir.Readdir(-1); err == nil {
			for _, fi := range fileInfos {
				fmt.Println(fi.Name())
			}
			dir.Close()
			return err
		}
		return errors.New("read dir failed")
	}
	return errors.New("open dir failed")
}

// sort demo
type IntArray []int

func (this IntArray) Len() int {
	return len(this)
}

func (this IntArray) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntArray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {

	//sort
	arr2 := []int{3,54,7,843,2,87,54,1}
	sort.Ints(arr2)
	//sort.Sort(IntArray(arr2))
	fmt.Println(arr2)

	//Write file
	if rfile, err := os.Create("test.txt"); err == nil {
		rfile.WriteString("hello world")
		rfile.Close()
	}

	// Read file
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	//get file size
	if stat, err := file.Stat(); err == nil {
		bs := make([]byte, stat.Size())
		if _, err := file.Read(bs); err == nil {
			str := string(bs)
			fmt.Println(str)
		}
	}

	//walks the file tree and delete exe file
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)
		if !info.IsDir() && strings.HasSuffix(info.Name(), "exe") {
			os.Remove(path)
			fmt.Println(path)
			//return filepath.SkipDir
		}
		return nil
	})

	// ls
	err1 := ls(".")
	fmt.Println(err1)

	//strings
	//true 1 2 true false hello world, welcome world who-are-you [who are you] [97 98 99 49 50 51] hello
	//arr := []rune("test")
	arr := []byte("abc123")
	str := string([]byte{'h', 'e', 'l', 'l', 'o'})
	fmt.Println()
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Index("test", "es"),
		strings.Count("test", "t"),
		strings.HasSuffix("package.go", "go"),
		strings.HasPrefix("package.go", "ack"),
		strings.Replace("hello pitou, welcome pitou", "pitou", "world", -1),
		strings.Join([]string{"who", "are", "you"}, "-"),
		strings.Split(strings.Join([]string{"who", "are", "you"}, "+"), "+"),
		arr, str,
	)

	//container
	var lst list.List
	lst.PushBack(1)
	lst.PushBack(1.12)
	lst.PushBack(112)
	lst.PushFront(100)
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
