package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"
)

const (
	KernelCount  = 4
	GoroutineCnt = 10
	ArraySize    = 100_000_000
)

func main() {
	runtime.GOMAXPROCS(KernelCount)
	arr := makeIntArray(ArraySize)
	fmt.Printf("Сравнение производительности\n")
	fmt.Printf("Количество ядер: %v\n", KernelCount)
	fmt.Printf("Количество горутин: %v\n", KernelCount)
	fmt.Printf("Размер массива: %v\n", ArraySize)

	start := time.Now()
	res := sumArray(&arr)
	totalTime := time.Since(start)
	fmt.Printf("Последовательный результат вычислений: %d\n", res)
	fmt.Printf("Последовательное время  вычислений: %.3v (c)\n", totalTime.Seconds())

	start = time.Now()
	resParallel := sumArrayParallel(arr, GoroutineCnt)
	totalTime = time.Since(start)
	fmt.Printf("Параллельный результат вычислений: %d\n", resParallel)
	fmt.Printf("Параллельное время  вычислений: %.3v (с)\n", totalTime.Seconds())
}

func makeIntArray(len int) []int {
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
func makeIntChannelArray(channelCount int) []chan int {
	chs := make([]chan int, channelCount)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, 1)
	}
	return chs
}

func sumArray(arr *[]int) int {
	var sum int
	for i := 0; i < len(*arr); i++ {
		for i := 0; i < 20; i++ { // имитция дополнительной "тяжелой" работы
			time.Sleep(time.Nanosecond / 2)
		}
		sum += (*arr)[i]
	}
	return sum
}
func sumArrayParallel(arr []int, goroutineCnt int) int {
	chs := makeIntChannelArray(goroutineCnt)
	var wg sync.WaitGroup
	for i := 0; i < goroutineCnt; i++ {
		wg.Add(1)
		partSize := len(arr) / goroutineCnt
		chunkArr := make([]int, partSize)
		copy(chunkArr, arr[partSize*i:partSize*(i+1)])
		go func(i int) {
			defer wg.Done()
			sumChunk(&chunkArr, chs[i])
		}(i)
	}
	wg.Wait()
	chResult := make(chan int)
	go getSumsParallel(chResult, chs)
	return <-chResult
}
func sumChunk(arr *[]int, ch chan int) {
	ch <- sumArray(arr)
}

func getSumsParallel(result chan int, chs []chan int) {
	var sum int
	cases := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}
	remaining := len(chs)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			// Если канал закрыт, убираем case из select-a
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining -= 1
			continue
		}
		m := sync.Mutex{}
		m.Lock()
		sum += int(value.Int())
		m.Unlock()
		close(chs[chosen])
	}
	result <- sum
}
