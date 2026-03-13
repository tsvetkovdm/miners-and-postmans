package main

import (
	"context"
	"fmt"
	"miners-and-postmans/miner"
	"miners-and-postmans/postman"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64

	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("--->>> Рабочий день шахтеров окончен!")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("--->>> Рабочий день почтальонов окончен!")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 3)
	mailTransferPoint := postman.PostmanPool(postmanContext, 3)

	start := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	// интересно просто
	// wg.Add(1)
	// go func(ctx context.Context) {
	// 	defer wg.Done()

	// 	<-ctx.Done()
	// 	fmt.Println("---------------------конец----------------------------")
	// }(postmanContext)

	wg.Wait()

	fmt.Println("Суммарный добытый уголь:", coal.Load())

	mtx.Lock()
	fmt.Println("Суммарное количество писем:", len(mails))
	mtx.Unlock()

	fmt.Println("Времени потрачено:", time.Since(start))

	fmt.Println("after возвращает:", <-time.After(1*time.Second))
}
