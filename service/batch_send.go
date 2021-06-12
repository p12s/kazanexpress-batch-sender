package service

// Batch - пачка элементов
type Batch []Item

// Item - элемент, абстрактный тип
type Item struct{}

// CreateBatch инициирует типизированный массив (с заранее определенным кол-вом элементов)
// с иниц. элементов массива
func CreateBatch(n uint64) Batch {
	index := int(n)
	var batch Batch = make([]Item, index)
	for i := 0; i < index; i++ {
		batch[i] = Item{}
	}
	return batch
}
