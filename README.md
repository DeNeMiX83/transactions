Изменение баланса на счету пользователя

Чистая архитектура

Запущен на сервере: https://miheev.su/transactions/

Endpoints:
- /user/balance/decrease
- /user/balance/decrease

Пример body:
```
{
    "user_id": "c91ce776-1853-4bc1-ab06-22b6a26f9bd1",
    "balanc_id": "87fbf1b4-503d-48de-b501-9cf3819822e1",
    "amount": 200
}
```
В базу данных сохраняются фейковые данные (файл: ```app/infrastructure/db/postgresql/fake_data.sql```)

Запуск

- В deploy переименуйте ```example.env``` в ```.env``` и подставьте туда свои данные.

- Запустите docker-compose
    ```
    make compose-up
    ```