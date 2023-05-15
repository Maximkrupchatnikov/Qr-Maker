package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// функция для создания qr кода, принимает переменную S, имя файля для сохранения
func qr(s string, filename string) {
	url := "http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl="
	url = url + s
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// декдироание полченного отеа на зпрос get запрос и запись в переменную img
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		panic(err)
	}
	//создание файла с qr кодом
	if filename == "0" {
		file, err := os.Create("qr_code" + s + ".png")
		if err != nil {
			fmt.Println("Ошибка при созранении qr кода", err)
			return
		}
		png.Encode(file, img)
		defer file.Close()
	} else {
		file, err := os.Create(filename + "/" + "qr_code" + s + ".png")
		if err != nil {
			fmt.Println("Ошибка сохранения qr кода", err)
			return
		}
		png.Encode(file, img)
		defer file.Close()
	}
}

func keyB() {
	fmt.Print("Введите кол-во запросов:", " ")
	//считывание кол-во запросов
	var n int
	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите имя папки для сохранения изображения: \n")
	scanner.Scan()
	filename := scanner.Text()

	fmt.Print("Введите текст для преобразования в qr код: \n")

	for i := 0; i < n; i++ {
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		s = strings.Replace(s, " ", "+", -1)
		qr(s, filename)
	}
	fmt.Println("Программа завершила свое выполнение ")
	main()
}

func textfile() {
	fmt.Println("Введите имя папки c текстом для преобразования в qr код: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	file := scanner.Text()

	var filename string
	fmt.Println("Введите имя папки для сохоранеения туда qr кодов")
	fmt.Scan(&filename)
	start := time.Now()
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка сканирования папки")
		return
	}

	files, err := os.Open(file)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}

	scanner = bufio.NewScanner(files)
	scanner.Split(bufio.ScanWords)
	var wg sync.WaitGroup
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.TrimSpace(word)

		wg.Add(1)
		go func(word, filename string) {
			defer wg.Done()
			qr(word, filename)
		}(word, filename)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка считывания файла", err)
		return
	}
	wg.Wait()
	fmt.Println("Программа завершила свое выполнение за", time.Since(start))
	main()
}

func main() {
	var value string
	// Выводим пользвательский интерфейс
	fmt.Println("Выберите действие для создания qr кода:")
	fmt.Println("1. Ввод c терминала")
	fmt.Println("2. Считывание текстового файла")
	fmt.Println("3. Выйти из программы")

	fmt.Scanln(&value)
	// стек вызова функций
	switch value {
	case "1":
		keyB()
	case "2":
		textfile()
	case "3":
		break
	default:
		fmt.Println("читай еще раз")
	}
}
