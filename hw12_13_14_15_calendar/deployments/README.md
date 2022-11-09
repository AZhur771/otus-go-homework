Каталог с рабочими файлами docker-compose, а также конфигами для запуска сервисов с помощью docker'а

Запуск бд + rabbit + админке
```shell
docker compose -f docker-compose.dev.yaml up
```

Запуск рабочей версии приложения
```shell
CALENDAR_CONFIG_DIR=$(pwd)/configs/ docker compose up
```