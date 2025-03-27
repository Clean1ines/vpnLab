package zerocore

import (
    "fmt"
    "net/http"
)

// Обработчик маршрута /hello – возвращает строку "hello"
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
}

// Обработчик маршрута /headers – выводит все HTTP-заголовки запроса
func headersHandler(w http.ResponseWriter, r *http.Request) {
    for name, values := range r.Header {
        for _, value := range values {
            fmt.Fprintf(w, "%s: %s\n", name, value)
        }
    }
}

func main() {
    // Регистрируем обработчики на маршрутах
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/headers", headersHandler)

    fmt.Println("Сервер запущен на порту 8090")
    // Запускаем сервер. Если возникнет ошибка – выводим её.
    if err := http.ListenAndServe(":8090", nil); err != nil {
        fmt.Println("Ошибка запуска сервера:", err)
    }
}
