package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func SumPart(nums []int, resultCh chan int) {
	res := 0
	for _, val := range nums {
		res += val
	}
	resultCh <- res
}

func main() {
	const parts = 10
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}
	resultCh := make(chan int, parts)

	p1 := nums[:1000/10]
	p2 := nums[1000/10 : 2*1000/10]
	p3 := nums[2*1000/10 : 3*1000/10]
	p4 := nums[3*1000/10 : 4*1000/10]
	p5 := nums[4*1000/10 : 5*1000/10]
	p6 := nums[5*1000/10 : 6*1000/10]
	p7 := nums[6*1000/10 : 7*1000/10]
	p8 := nums[7*1000/10 : 8*1000/10]
	p9 := nums[8*1000/10 : 9*1000/10]
	p10 := nums[9*1000/10 : 1000]

	allparts := [][]int{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			SumPart(allparts[i], resultCh)
			defer wg.Done()
		}()
	}
	wg.Wait()
	close(resultCh)
	sumCh := 0

	for val := range resultCh {
		sumCh += val
	}
	fmt.Println("Результат: ", sumCh)
}

//много затупов, сначала затуп на разделении слайса на части, не знал
//что при "срезе" можно использовать интервал. потом затуп с wg. нейронка
//хелпанула, потому что до этого я вызывал горутину sumpart и писал wg.Done
//в теле мейна а не горутины, поэтому не работало, через анонимную горутину
//работает вроде
