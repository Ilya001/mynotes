# MyNotes

## Небольшое описание проекта
 Технологии которые я использовал:
- **Front-end**: Html, Scss, Vue.Js
- **Back-end**: Go
- **DB**: Tarantool
- **Docker**

Я создал небольшое приложение для создания заметок(подобно Google Keep). Весь Back-end работает на Golang, он отдает всего один index.html, который контролирует Vue.js.

Vue.js общается с Golang при помощи json.
- /api/notes/ - основаная ссылка на которую Vue.js кидает запросы

В свою очередь golang сохраняет или берет данные из tarantool.

## Тур по папочкам
Приложение с тарантулом. Там мы создаем структуру БД на lua и хостим все это дело на :3301
```
tarantoolapp/
    Dockerfile
    notes.lua - Основной файл для создание структуры БД.
```
Golang(Apiserver) + FrontEnd
```
webapp/
    cmd/apiserver/
        main - основной файл запуска
    configs/
        apiserver.toml - конфиг файл для запуска api сервера 
        notes.toml - конфиг файл для взаимодействия tarantool и go
    interanl/
        apiserver/
            apiserver.go - основной файл где запускается api сервер
            config.go - структура для конфиг файла .toml
        note/
            note.go
            config.go - структура для конфиг файла .toml
    public/
        static/
            css/ - скомпилированный css
            js/
                main.js - файл где находится Vue.js
            scss/ - папка со всеми стилями
    Dockerfile
    go.mod
    go.sum
    Makefile - файл через который идется сборка go.

```

## Инструкция по запуску
Клонируем
```bash
git clone https://github.com/Ilya001/mynotes
```
Переходим в папку
```bash
cd mynotes
```
Запускаем
```bash
docker-compose build
```
Поднимаем
```bash
docker-compose up
```
Переходим по ссылке http://localhost:8080/

