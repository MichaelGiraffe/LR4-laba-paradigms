// go docs
package main

//импорт был за ии, но fmt у меня был и до этого, а os добавил бы из-за ввода пользователем параметров
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const API_KEY = "https://v6.exchangerate-api.com/v6/ce571ea97c53f23430eb5791/latest/"

// ExchangeResponse — структура ответа API
//type ExchangeResponse struct {
//	Result             string             `json:"result"`
//	BaseCode           string             `json:"base_code"`
//	ConversionRates    map[string]float64 `json:"conversion_rates"`
//	ErrorType          string             `json:"error-type"`
//}

// fetchRates делает запрос к API и возвращает карту курсов
func fetchRates(value_input string, value_output string, amount float64) {
	var url = API_KEY + value_input

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("cannot do that")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("can't do either")
		return
	}

	// Используем map[string]any вместо структуры
	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	// Извлекаем вложенную карту с курсами
	// Приходится кастить к map[string]any, так как json.Unmarshal
	// кладет объекты именно в этот тип
	rates, ok := data["conversion_rates"].(map[string]any)
	if !ok {
		fmt.Println("Could not find conversion_rates in response")
		return
	}

	// Извлекаем конкретный курс и приводим его к float64
	rate, ok := rates[value_output].(float64)
	if !ok {
		fmt.Printf("Currency %s not found\n", value_output)
		return
	}

	// Считаем и выводим
	fmt.Printf("Result: %.2f %s\n", amount*rate, value_output)
}

// мою любимую скобку фигурную писать нужно после функции :(
func main() {
	fmt.Println("I want a coffee!")

	var info = "hello"
	fmt.Println(info)

	var value_input string = ""  //пользователь вводит 1 валюту
	var value_output string = "" //пользователь вводит 2 валюту для конвертации
	var amount float64 = 0

	fmt.Println("введите фигню")
	fmt.Fscan(os.Stdin, &value_input)
	fmt.Fscan(os.Stdin, &value_output)
	fmt.Fscan(os.Stdin, &amount)

	//ToUpper-ии
	value_input = strings.ToUpper(value_input)
	value_output = strings.ToUpper(value_output)
	//ToUpper-ии

	fetchRates(value_input, value_output, amount)
}
