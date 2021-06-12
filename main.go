package main

import (
	"context"
	"fmt"
	"github.com/p12s/kazanexpress-batch-sender/service"
	"log"
	"math"
	"time"
)

/*
Принцип работы программы:
- запрашиваем макс. кол-во элементов, которое может обработать сервис. В примере - 10 элементов каждые 3 сек.
- в цикле, раз в секунду, отправляем сервису 10/3 (с округлением до целого в меньшую сторону) = 3 элемента
 */

func main() {
	// заглушка-ответ от внешнего сервиса
	externalService := service.ExternalService{
		MaxProcessItem:        10,
		ProcessPeriod:         3, // пусть период обработки будет в секундах
		AlreadyProcessedCount: 0, // для простоты примем, что уже обработано столько (сервис брал бы это кол-во из своей бд, сколько он уже обработал за последний период времени)
	}
	// имитируем запрос лимита и периода у внешнего сервиса
	limit, duration := externalService.GetLimits()


	var ctx context.Context // пустой контекст для передачи в метод

	// отправляем пачку элементов в кол-ве = n/(t*60 сек) в период времени - для равномерности
	itemInSec := uint64(math.Floor(float64(limit) / float64(duration)))
	batch := service.CreateBatch(itemInSec)

	for range time.Tick(time.Second) {
		err := externalService.Process(ctx, batch)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("Batch send period 1 sec")
	}
}
