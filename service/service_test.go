package service

import (
	"context"
	"math"
	"testing"
	"time"
)

func Test_GetLimits(t *testing.T) {

	testTable := []struct {
		Name                  string
		MaxProcessItem        uint64
		ProcessPeriod         time.Duration
		AlreadyProcessedCount uint64
	}{
		{
			Name:                  "Return limits OK",
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
			if limit == 0 {
				t.Fatalf("Error: limit is not defined")
			}

			if duration == 0 {
				t.Fatal("Error: duration is not defined")
			}
		})
	}
}

func Test_Process(t *testing.T) {
	testTable := []struct {
		Name                  string
		MaxProcessItem        uint64
		ProcessPeriod         time.Duration
		AlreadyProcessedCount uint64
		WantErr               bool
	}{
		{
			Name:                  "Success if the limit is not exceeded",
			MaxProcessItem:        100,
			ProcessPeriod:         3,
			AlreadyProcessedCount: 0, // лимит - не более 33 элемента в сек. ОК
			WantErr:               false,
		},
		{
			Name:                  "Blocked if the limit is exceeded",
			MaxProcessItem:        100,
			ProcessPeriod:         3,
			AlreadyProcessedCount: 999, // в логе уже 999 выполненных элементов. не ОК
			WantErr:               true,
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
			var ctx context.Context

			itemInSec := uint64(math.Floor(float64(limit) / float64(duration)))
			batch := CreateBatch(itemInSec)

			err := externalService.Process(ctx, batch)
			if err != nil && !testCase.WantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.WantErr {
				t.Fatal("Error was expected, but got nil")
			}
		})
	}
}
