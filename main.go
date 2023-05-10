package main

import (
	//"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
        "bufio"
        "strings"
)

func main() {
  sc := bufio.NewScanner(os.Stdin)
  
  sc.Scan()
  
  s := sc.Text()
  
  string := strings.Replace(s, " ","+", -1)
  
  url := "http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl="
  url += string
	
  resp, err := http.Get(url)
	
  if err != nil {
		panic(err)
	}
  
	img, _, err := image.Decode(resp.Body) // декдироание полченного отеа на зпрос get запрос и запись в переменную img	
  if err != nil {
		panic(err)
	}

	file, err := os.Create("qr_code.png") // создание файла и переменной file 
	
  if err != nil {
		panic(err)
	}
  
	png.Encode(file, img) //декодирование в файл из переменной img

	if err != nil {
		panic(err)
	}
  
}

