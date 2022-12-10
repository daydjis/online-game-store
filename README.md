# Описание функциональностей на бэке:
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
```
```
Удаление игры
DELETE /api/games/delete
BODY: {"id": {string}}
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
