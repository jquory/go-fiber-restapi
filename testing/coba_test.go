package testing

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

type ErrBadRequest struct {
	s string
}

type ErrNotFound struct {
	s string
}

func (b *ErrBadRequest) Error() string {
	return b.s
}

func (b *ErrBadRequest) New(s string) ErrNotFound {
	return ErrNotFound{s: s}
}

func TwoSum(numberOne, numberTwo int) (int, error) {

	if numberOne == 0 {
		return 0, errors.New("Angka tidak boleh 0")
	}

	return numberOne + numberTwo, nil
}

func TestSum(t *testing.T) {
	result, err := TwoSum(2, 2)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	fmt.Println(result)
}

func TestCoba(t *testing.T) {
	fmt.Println("hello world")

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		time.Sleep(2 * time.Second)
		fmt.Println("hello from goroutine")
	}(wg)

	fmt.Println("hello from main")
	wg.Wait()

	// di service
	err := &ErrBadRequest{"bad reuquest"}

	// controller
	if errors.Is(err, &ErrBadRequest{}) {
		fmt.Println("error bad request")
	}

	if strings.Contains(err.Error(), "bad") {
		fmt.Println("bad request juga")
	}
}

func TestCobaFunction(t *testing.T) {
	data, err := SaveData(2)
	if err != nil {
		// jika ada error
		fmt.Println(err.Error())
		return
	}

	// success get data
	fmt.Println(data)
}

func SaveData(id int) (string, error) {
	if id == 0 {
		return "", errors.New("id tidak boleh 0")
	}

	// success
	data := fmt.Sprintf("success get data dengan id %v", id)
	return data, nil
}

func TestTwoSum(t *testing.T) {
	twSum := func(nums []int, target int) []int {
		temp := map[int]int{}

		for id, num := range nums {
			hasilKurang := target - num
			if res, ok := temp[hasilKurang]; ok {
				return []int{res, id}
			}

			temp[num] = id
		}

		return []int{}
	}

	fmt.Println(twSum([]int{1, 2, 3, 7}, 9))
}

func TestBestTimeToBuy(t *testing.T) {
	maxProfit := func(prices []int) int {

		if len(prices) == 1 {
			return 0
		}
		maxxprofit := 0

		left := 0
		right := 1

		for right < len(prices) {
			profit := prices[right] - prices[left]
			maxxprofit = max(maxxprofit, profit)

			if prices[right] < prices[left] {
				left = right
			}

			right++
		}

		return maxxprofit
	}

	fmt.Println(maxProfit([]int{100, 50, 300}))
}
