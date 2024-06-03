package usecase

import (
	"Media/models"
	"sort"

	"github.com/sirupsen/logrus"
)

type Case struct {
	Logger *logrus.Logger
	Cash   models.CashInfo
}

func NewCase(logger *logrus.Logger) *Case {
	return &Case{
		Logger: logger,
		Cash:   models.CashInfo{},
	}
}

func (c *Case) Culc() *models.Response {
	var results [][]int

	var findCombinations func(amount int, banknotes []int, combination []int)
	findCombinations = func(amount int, banknotes []int, combination []int) {
		if amount == 0 {
			// Копируем комбинацию, чтобы избежать изменения при дальнейшем использовании
			combCopy := make([]int, len(combination))
			copy(combCopy, combination)
			results = append(results, combCopy)
			return
		}
		for i, note := range banknotes {
			if note <= amount {
				// Добавляем банкноту в текущую комбинацию
				newCombination := append(combination, note)
				// Рекурсивно ищем остальные комбинации с уменьшенной суммой
				findCombinations(amount-note, banknotes[i:], newCombination)
			}
		}
	}

	// Сортируем банкноты по убыванию
	sort.Sort(sort.Reverse(sort.IntSlice(c.Cash.Banknotes)))
	// Ищем все возможные размены
	findCombinations(c.Cash.Amount, c.Cash.Banknotes, []int{})

	return &models.Response{Exchanges: results}
}
