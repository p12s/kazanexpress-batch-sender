package service

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Service может обрабатывать элементы пачками
type Service interface {
	GetLimits() (n uint64, t time.Duration)         // лимит - кол-во элементов в период времени
	Process(ctx context.Context, batch Batch) error // обработать пачку
}

// ExternalService - имитация ответа сервиса
// хранение кол-ва обработанных элементов делать не будем,
// предполагаем что самое простое - при получении пачки элементов проверяем в БД сколько уже обработали и решаем
type ExternalService struct {
	MaxProcessItem uint64
	ProcessPeriod  time.Duration

	AlreadyProcessedCount uint64 // кол-во уже обработанных
}

// GetLimits - получение лимита сервиса
func (e *ExternalService) GetLimits() (n uint64, t time.Duration) {
	return e.MaxProcessItem, e.ProcessPeriod
}

// Process - Обработка
func (e *ExternalService) Process(ctx context.Context, batch Batch) error {
	sumProcessedCount := uint64(len(batch)) + e.AlreadyProcessedCount
	if sumProcessedCount > e.MaxProcessItem {
		return errors.New("blocked")
	}
	fmt.Printf("Success processed: %d items\n", len(batch))
	return nil
}
