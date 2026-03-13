package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	for {

		// Задача 1: Шахтер СРАЗУ завершит (прервет) выполнение работы, если рабочий день завершен

		fmt.Println("Я шахтер номер", n, "начал добычу угля!")

		select {
		case <-ctx.Done():
			fmt.Println("Шахтер номер:", n, "ушёл домой!")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я шахтер ноиер:", n, "добыл", power, "угля.")
		}

		select {
		case <-ctx.Done():
			fmt.Println("Шахтер номер:", n, "ушёл домой!")
			return
		case transferPoint <- power:
			fmt.Println("Я шахтер номер:", n, "передал", power, "угля.")
		}

		// Задача 2: Шахтер доделает свою работу ДО КОНЦА, даже если рабочий день уже завершен

		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Я шахтер номер", n, "завершил смену!")
		// 	return
		// default:
		// 	fmt.Println("Я шахтер номер", n, "начал добычу угля!")
		// 	time.Sleep(time.Second)
		// 	fmt.Println("Я шахтер номер", n, "добыл", power, "угля.")

		// 	transferPoint <- power
		// 	fmt.Println("Я шахтер номер", n, "передал", power, "угля.")
		// }
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
