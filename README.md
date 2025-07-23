# Инструкция для старта фронтенд части

1. Поменяйте папку на /client
   `cd client`
2. Установите зависимости
   `npm install`
3. Запустите сервер
   `npm run dev`

# Инструкция для старта бекенд части

1. Поменяйте папку на /devschool-backend

   `cd devschool-backend`

2. Установите зависимости
   `go mod tidy`
3. Запустите сервер (Также должен быть запущена база данных)
   `go run main.go`

# Используемые инструменты для фронтенда

    ├── Javscript + Html + Css
    ├── Vue
    ├── Vite
    └── Postcss

# Используемые инструменты для бекенда

    ├── Golang
    └── PostgreSQL

# Важные замечания

Для взаимодействия golang с postgresql базой данных вы должны иметь конфигурацию .env файла
Например:

_devschool-backend/.env :_

```
DB_USER=YourDatabaseUser
DB_PASSWORD=YourPassword
DB_NAME=YourDatabaseName
```

# Инструменты, которые будут добавлены в будущем

Я хочу добавить TypeScript, NuxtJs и Gin для golang после некоторого времени, когда я стану лучше в программировании в целом :D

# Свяжитесь со мной в телеграме: @alexsharav
