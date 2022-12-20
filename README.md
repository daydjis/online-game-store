# Описание функциональностей на бэке:

Для создания или удаления игры, необходимо авторизоваться.  
Для авторизации в БД уже создан пользователь со следующими данными:

```
login = admin_user
password = 12345678
```

В случае удачной авторизации, бэк отправляет клиенту токен в ответе.  
Полученный токен необходимо отправлять на сервер вместе с каждым запросом внутри хедера Authorization.

## Доступные ручки:

```
Получение списка игр
GET /api/games
GET /api/games?id={gameID}
```

```
Создание новой игры
POST /api/games/new
BODY: {"title": {string},
       "description": {string},
       "price": {int},
       "genres": {Array[string, string, ...]},
       "image": {string},
       "video": {string},
       "imageDescription": {string}}

- Присутствует валидация запроса
- Необходима авторизация
```

```
Удаление игры
DELETE /api/games/delete
BODY: {"id": {string}}

- Необходима авторизация
```

```
Регистрация нового пользователя
POST /api/register
BODY: {"login": {string},
       "password": {string}}

- Присутствует валидация запроса
```

```
Авторизация пользователя
POST /api/login
BODY: {"login": {string},
       "password": {string}}

- Присутствует валидация запроса
```

# Поднимаем фронт и бэк через docker-compose:

### Поднимается все в одну команду:

```
docker compose up
```

### Если же до этого у вас уже был поднят compose, а теперь вы хотите пересобрать контейнеры, чтобы подтянуть в них изменения:

```
docker compose build
docker compose up
```

# Поднимаем бэк через Docker:

```
cd back
docker build . -f DockerFile --tag back
docker run -p 0.0.0.0:5000:5000 --name back back
```

# Поднимаем фронт через Docker:

```
docker build . -f DockerFile --tag front
docker run -p 0.0.0.0:3000:3000 --name front front
```
