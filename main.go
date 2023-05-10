package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strings"
)

func qr(s string, filename string) {

	url := "http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl="
	url += s
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(resp.Body) // декдироание полченного отеа на зпрос get запрос и запись в переменную img
	if err != nil {
		panic(err)
	}
	if filename == "0" {
		file, err := os.Create("qr_code" + s + ".png")
		if err != nil {
			panic(err)
		}
		png.Encode(file, img)
	} else {
		file, err := os.Create(filename + "qr_code" + s + ".png")
		if err != nil {
			panic(err)
		}
		png.Encode(file, img)
	}
	
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

}

func keyb() {
	fmt.Print("Введите кол-во запросов:", " ")

	var n int
	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите имя папки для сохранения изображения: ")
	scanner.Scan()
	filename := scanner.Text()

	fmt.Print("Введите текст для преобразования в qr код:")

	for i := 0; i < n; i++ {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		s := sc.Text()
		s = strings.Replace(s, " ", "+", -1)
		qr(s, filename)
	}
}

func textfile() {

}

func JsonEnCode() {

}

func main() {
	var value string
	fmt.Println("Выберите действие:")
	fmt.Println("1. Ввод с терминала")
	fmt.Println("2. Считывание текстового файла")
	fmt.Println("3. Считывание JSON файла")

	fmt.Scanln(&value)

	switch value {
	case "1":
		keyb()
	case "2":
		textfile()
	case "3":
		JsonEnCode()
	default:
		fmt.Println("читай еще раз")
	}
}
