package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	Dataq := DataQueue{Max: 10}
	for i := 0; i < 10; i++ {
		Dataq.Enqueue(Data(i))

	}
	Dataq.Dequeue()
	fmt.Println(Dataq)
    time.Sleep(10*time.Second)
}

type Data interface {
}

type DataQueue struct {
	Max   int
	DataQ []Data
}

type Queue interface {
	New(int) (d *DataQueue, err error)
	Enqueue(d ...Data) error
	Dequeue() (data *Data, err error)
	Size() int
}

func (dq *DataQueue) New(i int) (d *DataQueue, err error) {
	if i <= 0 {
		fmt.Println(errors.New("最大长度不可以小于1"))
		return nil, errors.New("最大长度不可以小于1")
	}
	dq.Max = i
	dq.DataQ = []Data{}
	return dq, nil
}

//入列
func (dq *DataQueue) Enqueue(d ...Data) error {
	for _, dd := range d {
		if dq.Size() == dq.Max {
			fmt.Println(errors.New("超过队列最大容量"))
			return errors.New("超过队列最大容量")
		}
		fmt.Println("入列：", dd)
		dq.DataQ = append(dq.DataQ, dd)
		fmt.Println("dq:",dq)
	}
	return nil
}

//出列
func (dq *DataQueue) Dequeue() (data *Data, err error) {
	if dq.Size() <= 0 {
		fmt.Println(errors.New("队列为空"))
		return nil, errors.New("队列为空")
	}
	d := dq.DataQ[0]
	fmt.Println("出列=========：", d)
	dq.DataQ = dq.DataQ[1:]
	return &d, nil
}

//获得切片大小
func (dq *DataQueue) Size() int {
	return len(dq.DataQ)
}
