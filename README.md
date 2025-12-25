# GolangReceiptService

1. Запустить сборку контейнеров 
```
docker compose up
```

2. Скрипт для отправки сообщений
```
python ./script_reviews.py
```

3. Покрытие reviewService тестами:
```
cd services/reviews
go test ./internal/services/reviewService -cover
```

Ожидается:
```
ok      reviews_service/internal/services/reviewService 0.007s  coverage: 100.0% of statements
```
