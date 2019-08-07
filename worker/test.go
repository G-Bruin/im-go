package worker

import (
	"fmt"
	"github.com/benmanns/goworker"
)

// ./worker -queues="myqueue,queues"
// RPUSH resque:queue:myqueue '{"class":"MyClass","args":["hi","there"]}'
func myFunc(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v\n", queue, args)
	return nil
}

func init() {
	settings := goworker.WorkerSettings{
		//URI:            "redis://localhost:6379/",
		//redis://ignored:{$pass}@{$host}:{$port}
		URI:            "redis://root:123456@localhost:6379/2",
		Connections:    100,
		Queues:         []string{"myqueue", "queues"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "resque:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("MyClass", myFunc)
}

func main2() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}
