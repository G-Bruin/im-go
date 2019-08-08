package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

var (
	MaxWorker = 10
)

type Payload struct {
	Num int
}

//待执行的工作
type Job struct {
	Payload Payload
}

//任务channal
var JobQueue chan Job

//执行任务的工作者单元
type Worker struct {
	WorkerPool chan chan Job //工作者池--每个元素是一个工作者的私有任务channal
	JobChannel chan Job      //每个工作者单元包含一个任务管道 用于获取任务
	quit       chan bool     //退出信号
	no         int           //编号
}

//创建一个新工作者单元
func NewWorker(workerPool chan chan Job, no int) Worker {
	fmt.Println("创建一个新工作者单元")
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		no:         no,
	}
}

//循环  监听任务和结束信号
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
			fmt.Println("w.WorkerPool <- w.JobChannel", w)
			select {
			case job := <-w.JobChannel:
				fmt.Println("job := <-w.JobChannel")
				// 收到任务
				fmt.Println(job)
				time.Sleep(100 * time.Second)
			case <-w.quit:
				// 收到退出信号
				return
			}
		}
	}()
}

// 停止信号
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//调度中心
type Dispatcher struct {
	//工作者池
	WorkerPool chan chan Job
	//工作者数量
	MaxWorkers int
}

//创建调度中心
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

//工作者池的初始化
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 1; i < d.MaxWorkers+1; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}
	go d.dispatch()
}

//调度
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("job := <-JobQueue:")
			go func(job Job) {
				fmt.Println("等待空闲worker (任务多的时候会阻塞这里")
				//等待空闲worker (任务多的时候会阻塞这里)
				jobChannel := <-d.WorkerPool
				fmt.Println("jobChannel := <-d.WorkerPool", reflect.TypeOf(jobChannel))
				// 将任务放到上述woker的私有任务channal中
				jobChannel <- job
				fmt.Println("jobChannel <- job")
			}(job)
		}
	}
}

func main() {
	JobQueue = make(chan Job, 10)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	time.Sleep(1 * time.Second)
	go addQueue()
	time.Sleep(1000 * time.Second)
}

func addQueue() {
	for i := 0; i < 100; i++ {
		// 新建一个任务
		payLoad := Payload{Num: i}
		work := Job{Payload: payLoad}
		// 任务放入任务队列channal
		JobQueue <- work
		fmt.Println("JobQueue <- work", i)
		fmt.Println("当前协程数:", runtime.NumGoroutine())
		time.Sleep(100 * time.Millisecond)
	}
}
