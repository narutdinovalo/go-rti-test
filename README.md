# go-rti-testing
#### Запуск сервиса
1. Скачать репозиторий
2. Запустить сервис ```go run main.go```

#### endPoints
1. HealthCheck проверка работоспособности сервиса ```/health```
При испешном запуске сервиса получим ответ:
```json
{"status": "ok"}
``` 

2. Calculate получение продуктового предложения ```/calculate```
Пример:
##### Запрос 
```json
[
  {
    "ruleName": "technology",
    "value": "xpon"
  },
  {
    "ruleName": "internetSpeed",
    "value": "200"
  }
]
```  
##### Ответ
```json
{
  "name": "Игровой",
  "components": [
    {
      "name": "Интернет",
      "isMain": true,
      "prices": [
        {
          "cost": 765
        }
      ]
    }
  ],
  "totalCost": {
    "cost": 765
  }
}
``` 

#### docker
1. Запуск docker ```docker-compose up```

#### Автотестирование
1. Запуск тестов ```go test -v ./...```