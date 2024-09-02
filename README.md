# MicroserviseGo

В проекте есть 3 микросервиса + 1 в качестве api-gateway. Взаимодействие между сервисами происходит по gRPC
Микросервис Api-Gateay - служит шлюзом и реализует все endpoint для взаимодействия с внешним миром, реализован свагер
Микросервис Auth - служит для работы с авторизацией и аутентификацией, реализована с помощью JWT RS256
Микросервис Posts - дает возможнсть создавать посты и оставлять комментарии к постам
Микросервис Likes - дает возмоэность проставлять и удалять реакции на посты
Проект релизован с помощью Gin, Gorm, Postgres, Goose, Zerolog, Redis, Docker, Docker-compose, Opentelemetry, Jaeger, Prometheus, Grafana, Loki

## Чтобы запустить проект у себя локально, вам необохимо
Склонировать репозиторий

`https://github.com/fazletdinov/microserviceGo.git`

Скопируйте содержимое .env.example в .env во всех микросервисах

И запустите команду в каждом микросервисе

```
make start
```
Если не установлена утилита make, то необходимо запустить следующей командой

```
docker compose -f docker-compose.{Название сервиса}.yaml up --build
```
Вышеуказанные команды запустит приложение
Далее можете посмотреть Api спецификацию (свагер) по адресу:
`http://localhost:8000/docs/index.html`

### Автор
[Idel Fazletdinov - fazletdinov](https://github.com/fazletdinov)