package testing

import (
	"fmt"
	"testing"
)

func TestGetData(t *testing.T) {
	data := [][]string{
		{
			"Uzumaki Naruto", "Uzumaki Ujang",
		},
		{
			"",
		},
	}
	getData := func(nums []string) []string {

		if len(nums) > 1 {
			for _, member := range data[0] {
				for _, member2 := range data[1] {
					if member == member2 {
						return []string{member}
					}
				}
			}
		}

		return data[0]
	}

	fmt.Println(getData([]string{"uzumaki", "ujang"}))
}
