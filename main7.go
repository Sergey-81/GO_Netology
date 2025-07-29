package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Обработка файла с использованием функции process
	if err := ReadProcessWrite(
		"input.txt",
		"output.txt",
		process,
	); err != nil {
		log.Fatal("Ошибка обработки файла:", err)
	}

	// Для демонстрации - чтение и вывод результата
	data, err := os.ReadFile("output.txt")
	if err != nil {
		log.Fatal("Ошибка чтения выходного файла:", err)
	}
	fmt.Print("Содержимое выходного файла:\n", string(data))
}

func ReadProcessWrite(
	inputPath string,
	outputPath string,
	process func(string) (string, error),
) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("чтение файла: %w", err)
	}

	processedData, err := process(string(data))
	if err != nil {
		return fmt.Errorf("обработка данных: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(processedData), 0644); err != nil {
		return fmt.Errorf("запись файла: %w", err)
	}

	return nil
}

func process(s string) (string, error) {
	return strings.ToUpper(s), nil
}