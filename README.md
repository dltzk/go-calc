# Go-Web-Calculator
Веб-сервис, способный получать на вход заданное выражение по HTTP и получить в ответ его результат.

Сервис имеет 1 endpoint: api/v1/calculate

Калькулятор поддерживает следующие операции:
- Сложение
- Вычитание
- Умножение
- Деление
- Открытые и закрытые скобки (приоритизация)

# Установка и запуск

Для запуска этого проекта необходимо иметь:
- Установленный Go (версия не ниже 1.23)
- Доступ к командной строке

А для его отладки желательно использовать Postman или curl. Лично мной веб-сервис тестировался в Postman'е. Чтобы работать с нашим локальным API, также нужно установить Postman Agent (https://www.postman.com/downloads/postman-agent/).

1. Склонируйте репозиторий в удобную папку:
```bash
git clone https://github.com/dltzk/go-calc
```
2. Перейдите в каталог проекта
```bash
cd <путь_к_проекту>
```
3. Запустите файл main.go в папке cmd
```bash
go run cmd/main.go
```
4. Перейдите в Postman Desktop или зайдите на сайт, создайте Workspace, нажмите на + и готовьтесь отправлять запросы.
![Вспомогающий скриншот](https://sun9-53.userapi.com/impg/wUdOux05nfYLqvW1Ugl3e-JghpKl_oj4Bbrm6A/l5qkAYQlzTg.jpg?size=1915x919&quality=95&sign=f11719393418633767f1dd9212ecf5f9&type=album)
![2](https://sun9-79.userapi.com/impg/2Mq6jvFuPGNrABxHizZ6whzHbbLIR_GH_DwcHA/AILJth4W6qA.jpg?size=1916x912&quality=95&sign=f71ed83088704c42148470eceec8f3e9&type=album)

# Примеры запроса к API:
Нужно помнить, что сервис принимает только POST запросы, на другие запросы он будет ссылаться на неизвестную ошибку и выдавать code 500, даже если выражение верное.

### Пример запроса с 500 кодом:
1. Запрос с неправильным методом:
![3](https://sun9-44.userapi.com/impg/X6u-2yi6NbI06rgQJRm9SRgpnYx-LmFxigGoKA/KSqsj3_kOq8.jpg?size=1917x912&quality=95&sign=78c9a343be52b8b6120088576c9ba21d&type=album)
2. Запрос с поломанным body:
![4](https://sun9-27.userapi.com/impg/F1gaJYDedClGiyMFc6QhCHTV8u62mqE2UbxcXA/qskwKI2LHuY.jpg?size=1918x913&quality=95&sign=113ca6c40cd4a6e801e87f2db872eabe&type=album)

### Пример запроса с 422 кодом:
1. Expression is not valid - выражение построено неправильно
![5](https://sun9-16.userapi.com/impg/W-cWZbrVgr4TNfrLfnpgkabQGeIu2HfTk_-8Mg/7xf-sicrI40.jpg?size=1918x913&quality=95&sign=79c222a6432dbd5de7fffa4cae5c1c10&type=album)
2. Division by zero - деление на ноль)
![6](https://sun9-43.userapi.com/impg/U0IaeW6jl58oRzldTUesP9Dq0gk6HAlmI4UJGg/miM_4EvPg7o.jpg?size=1914x911&quality=95&sign=3cc95b66955349a4f6187c098290d180&type=album)

### Пример корректного запроса (Status code - 200, ответ верный):
![7](https://sun9-65.userapi.com/impg/cU9hoRxLvuJds3vaX7p_1ipotXqMn0RxW6BI5Q/_lrIOr2U5Rk.jpg?size=1919x910&quality=95&sign=7a0033f9d562217ac9fc718571968890&type=album)
