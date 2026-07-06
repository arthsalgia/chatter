package services

import (
	"fmt"
	"sort"
)

type Freqs struct {
	Data             string `json:"data" dataframe:"data"`
	NumberOfMessages int    `json:"number_of_messages" dataframe:"number_of_messages"`
}

func NthCommonString(arr []string, n int) []any {
	if len(arr) == 0 || n <= 0 {
		return []any{}
	}

	inputMap := make(map[string]int, len(arr)/10)
	for _, val := range arr {
		inputMap[val]++
	}

	freqsList := make([]Freqs, 0, len(inputMap))
	for data, count := range inputMap {
		freqsList = append(freqsList, Freqs{
			Data:             data,
			NumberOfMessages: count,
		})
	}

	sort.Slice(freqsList, func(i, j int) bool {
		return freqsList[i].NumberOfMessages > freqsList[j].NumberOfMessages
	})

	if n > len(freqsList) {
		fmt.Println("Requested rank exceeds dataset size")
		return []any{}
	}

	sort.Slice(freqsList, func(i, j int) bool {
		return freqsList[i].NumberOfMessages > freqsList[j].NumberOfMessages
	})

	var items []Freqs
	for i := range n {
		items = append(items, freqsList[i])
	}

	return []any{items}
}
