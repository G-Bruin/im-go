package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//声明成游戏
type Payload struct {
	name string
}

//任务
type Job struct {
	Payload Payload
}

//任务队列
var JobQueue chan Job

//  执行者 消费者 工人
type Worker struct {
	WorkerPool chan chan Job //对象池
	JobChannel chan Job      //通道里面拿
	quit       chan bool     //
	name       string        //工人的名字
}

// 调度器
type Dispatcher struct {
	name       string        //调度的名字
	maxWorkers int           //获取 调试的大小
	WorkerPool chan chan Job //注册和工人一样的通道
}

//打游戏
func (p *Payload) Play() {
	fmt.Printf("%s 打LOL游戏...当前任务完成\n", p.name)
}

// 新建一个工人
func NewWorker(workerPool chan chan Job, name string) Worker {
	fmt.Printf("创建了一个工人,它的名字是:%s \n", name)
	//workerPool 确定woker的容量
	return Worker{
		name:       name,           //工人的名字
		WorkerPool: workerPool,     //工人在哪个对象池里工作,可以理解成部门
		JobChannel: make(chan Job), //工人的任务
		quit:       make(chan bool),
	}
}

// 工人开始工作

func (w *Worker) Start() {
	//开一个新的协程
	go func() {
		for {
			//注册到对象池中,
			w.WorkerPool <- w.JobChannel
			fmt.Printf("[%s]把自己注册到 对象池中 \n", w.name)
			select {
			//接收到了新的任务
			case job := <-w.JobChannel:
				fmt.Printf("[%s] 工人接收到了任务 当前任务的长度是[%d]\n", w.name, len(w.WorkerPool))
				job.Payload.Play()
				time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			//接收到了任务
			case <-w.quit:
				fmt.Println("结束任务", w.name)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	//容量为{maxWorkers}的channel
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,       // 将工人放到一个池中,可以理解成一个部门中
		name:       "调度者",      //调度者的名字
		maxWorkers: maxWorkers, //这个调度者有好多个工人
	}
}

func (d *Dispatcher) Run() {
	// 开始运行
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, fmt.Sprintf("work-%s", strconv.Itoa(i)))
		//开始工作
		worker.Start()
	}
	//监控
	go d.dispatch()

}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("调度者,接收到一个工作任务")
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			// 调度者接收到一个工作任务
			go func(job Job) {
				//从现有的对象池中拿出一个
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		default:

			//fmt.Println("ok!!")
		}

	}
}

func initialize() {
	maxWorkers := 2
	maxQueue := 4
	fmt.Println(os.Getenv("MAX_QUEUE"))
	fmt.Println(os.Getenv("MAX_QUEUE"))
	//初始化一个调试者,并指定它可以操作的 工人个数
	dispatch := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue) //指定任务的队列长度
	//并让它一直接运行
	dispatch.Run()
}

func main() {
	//初始化对象池
	initialize()
	for i := 0; i < 10; i++ {
		p := Payload{
			fmt.Sprintf("玩家-[%s]", strconv.Itoa(i)),
		}
		JobQueue <- Job{
			Payload: p,
		}
		time.Sleep(time.Second)
	}
	//close(JobQueue)
}
