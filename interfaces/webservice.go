package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-rti-testing/domain"
	"go-rti-testing/infrastructure"
)

func Webservice() {
	healthCheck()
	handler()
	webserviceStart()
}

func webserviceStart() {
	infrastructure.Logger("Веб-сервис запущен")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		infrastructure.Logger("Ошибка запуска веб-сервиса", err)
	}
}

func healthCheck() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "ok"}`)
	})
}

func handler() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		condition := []domain.Condition{}
		err := json.NewDecoder(r.Body).Decode(&condition)
		if err != nil {
			infrastructure.Logger("Ошибка: не удалось распарсить json", err)
		}
		product := domain.GetProduct()
		offer, err := domain.Calculate(product, condition)
		if err != nil {
			infrastructure.Logger("Error", err)
		}
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		err = json.NewEncoder(w).Encode(offer)
		if err != nil {
			infrastructure.Logger("Ошибка: неудалось собрать json", err)
		}
	})
}
