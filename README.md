[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/mUq898-f)
# 2023-autumn-AB-Go-HW-9-template
GRPC

### Домашнее задание: Создание gRPC сервиса для потоковой передачи файлов
#### Задача: Разработайте gRPC сервис на Go для потоковой передачи файлов между клиентом и сервером.

### Требования к Proto-определениям:
 * Основной метод: отправка имени файла и получение потока байтов файла.
 * Метод получения списка файлов: возвращает список всех файлов на сервере.
 * Методы информации о файлах: получение информации о файлах на сервере, включая размер и другие характеристики.

### Критерии:
  * Чистая архитектура: чёткое разделение клиента и сервера.
  * Доменные модели: конвертация запросов в доменные модели.
  * Сервер как прокси: сервер обрабатывает запросы, перенаправляя их в прикладной слой.
  * Клиент с таймаутами: реализация таймаутов и других полезных функций в клиенте.

### Дополнительные критерии:
 * Обработка ошибок и логирование: корректная обработка ошибок и логирование на сервере и клиенте.
 * Тестирование: написание unit-тестов для ключевых компонентов.
 * Документация: комментарии к коду и инструкция по запуску сервиса.
 * Валидация: добавить валидацию запросов