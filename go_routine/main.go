package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	num  int
	name string
}

var (
	qq = rand.Intn
	ch = make(chan Task, 100)
)

func main() {
	go product()
	//启动一定数量的worker
	for i := 0; i < 10; i++ {
		go worker(ch)
	}

	select {}
}

func product() {
	for {
		//发送任务给worker
		helloTasks := getTaks()

		for _, task := range helloTasks {
			ch <- task
		}
		fmt.Println("\n正在生产6个产品...")
		time.Sleep(time.Second)
	}
}

func worker(ch chan Task) {
	for {
		//接受任务
		task := <-ch
		process(task)
	}
}

func process(task Task) {
	fmt.Println("消费中", task.num, task.name)
}

func getTaks() (tt []Task) {
	tt = []Task{
		Task{1, "noddles"},
		Task{2, "cake"},
		Task{3, "steak"},
		Task{4, "dumplings"},
		Task{5, "soup"},
		Task{qq(9248), "pizza"},
	}
	return
}
