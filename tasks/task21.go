package tasks

import "fmt"

type DataStorage interface {
	Save(data string)
	Load() string
}

type FlashDrive struct {
	Data string
}

func (f *FlashDrive) Save(data string) {
	f.Data = data
}

func (f *FlashDrive) Load() string {
	return f.Data
}

type MemoryCardAdapter struct {
	CardData string
	Card     MemoryCard
}

func (a *MemoryCardAdapter) Save(data string) {
	a.Card.SaveToCard(data)
}

func (a *MemoryCardAdapter) Load() string {
	return a.Card.LoadFromCard()
}

type MemoryCard interface {
	SaveToCard(data string)
	LoadFromCard() string
}

type SDMemoryCard struct {
	CardData string
}

func (c *SDMemoryCard) SaveToCard(data string) {
	c.CardData = data
}

func (c *SDMemoryCard) LoadFromCard() string {
	return c.CardData
}

func TransferData(storage DataStorage, data string) {
	storage.Save(data)
	fmt.Println("Сохраненные данные:", storage.Load())
}

func Task21() {
	flashDrive := &FlashDrive{}
	TransferData(flashDrive, "Данные на флешке")

	sdCard := &SDMemoryCard{}

	cardAdapter := &MemoryCardAdapter{
		Card: sdCard,
	}

	TransferData(cardAdapter, "Данные на карте памяти через адаптер")
}
