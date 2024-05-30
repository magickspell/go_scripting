package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
ФЛОУ теста - 3 этапа
1 подготовка данных
2 выполнение кода
3 проверка что все правильно выполнилось
*/

func Test_Calc(t *testing.T) {
	t.Parallel() // обычно тесты запускают паралельно

	t.Run("sum is calculated right", func(t *testing.T) {
		t.Parallel() // паралелим тесты внутри теста

		result := 3

		res := CalcSum(1, 2)

		// на каждый асерт принимается t (тест), т.к. если проверка не пройдет все сфейлится
		assert.Equal(t, result, res)
		// а можно свой ассерт сделать
		if res != result {
			t.Errorf("Expected %d, got %d", result, res)
		}
	})

	t.Run("error for broken string", func(t *testing.T) {
		t.Parallel()

		res, err := CalcSumStr("text", "3")

		assert.NotNil(t, err)
		t.Log(err)
		assert.Equal(t, 0, res)
	})
}
