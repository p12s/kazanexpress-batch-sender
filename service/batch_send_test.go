package service

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func Test_CreateBatch(t *testing.T) {

	testTable := []struct {
		Name                  string
		MaxProcessItem        uint64
		ProcessPeriod         time.Duration
		AlreadyProcessedCount uint64
	}{
		{
			Name:                  "OK",
			MaxProcessItem:        100,
			ProcessPeriod:         3,
			AlreadyProcessedCount: 0,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {

			externalService := ExternalService{
				MaxProcessItem:        testCase.MaxProcessItem,
				ProcessPeriod:         testCase.ProcessPeriod,
				AlreadyProcessedCount: testCase.AlreadyProcessedCount,
			}
			limit, duration := externalService.GetLimits()

			itemCount := uint64(math.Floor(float64(limit) / float64(duration)))
			batch := CreateBatch(itemCount)

			// проверка длины
			if batch == nil || uint64(len(batch)) != itemCount {
				t.Fatalf("Error: batch length is wrong")
			}

			// проверка типов элементов
			for i := 0; i < len(batch); i++ {
				got := reflect.TypeOf(batch[i])
				expected := reflect.TypeOf(Item{})
				if reflect.TypeOf(batch[i]) != reflect.TypeOf(Item{}) {
					t.Fatalf("Error: item element has wrong type: got: %s, expected: %s\n",
						got, expected)
				}
			}

		})
	}
}
