package main

import (
	"fmt" // пакет для форматированного ввода вывода
	"net/http" // пакет для поддержки HTTP протокола
	"strings" // пакет для работы с  UTF-8 строками
	"log" // пакет для логирования
	"github.com/shvetsovau/simplehttpserver/httpscerts"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,
	fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello User!") // отправляем данные на клиентскую сторону
}

func main() {
	// Проверяем, доступен ли cert файл.
	err := httpscerts.Check("cert.pem", "key.pem")

	// Если он недоступен, то генерируем новый.
	if err != nil {
			log.Fatal("Ошибка: https сертификат или приватный ключ не найден.")
		}
	}

	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	err := http.ListenAndServeTLS(":9000","cert.pem","key.pem", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
