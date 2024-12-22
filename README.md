Yandex-Calculator

Этот проект представляет собой веб-сервис для вычисления арифметических выражений. Пользователь может отправить запрос с арифметическим выражением и получить результат его вычисления.

Описание

Сервис предоставляет один endpoint для вычисления арифметических выражений:

- URL: /api/v1/calculate
- Метод: POST
- Тело запроса: JSON с полем expression, которое содержит арифметическое выражение.

Пример запроса:

`bash
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

СЦЕНАРИИ
Успешный запрос: 
{
  "result": "6"
}

Ошибка 422(недопустимое выражение):
{
  "error": "Expression is not valid"
}

Ошибка 500 (внутренняя ошибка сервера):
{
  "error": "Internal server error"
}

...........................................................
ИНСТРУКЦИЯ ПО ЗАПУСКУ

Склонируйте репозиторий:
git clone https://github.com/NiksonGo/Yandex-Calculator.git

Перейдите в каталог проекта:
cd Yandex-Calculator

Запустите проект:
go run ./cmd/calc_service

Проект будет запущен на порту 8080 по умолчанию
...........................................................

Пример использования

Успешное выполнение:
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "3+5*2"
}'

Ответ:
{
  "result": "13"
}

Ошибка 422 (недопустимое выражение):
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*ab" !!!ОБРАТИ ВНИМАНИЕ!!! Символ ab не является допустимым символом, т.к не предусмотрен в программе.
}'

Ответ:
{
  "error": "Expression is not valid"
}

Ошибка 500 (внутренняя ошибка):
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "(2+2"   !!!ОБРАТИ ВНИМАНИЕ!!! Несоответсвие синтаксиса,пропущена скобка,в большинстве калькуляторов такая ошибка приводит к основным ошибкам.
}'

Ответ:
{
  "error": "Internal server error"
}

