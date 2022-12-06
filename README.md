# Build entire project with docker-compose:
```
docker compose up
```
# Build backend with Docker:
```
cd back  
docker build . -f DockerFile --tag back  
docker run -p 0.0.0.0:5000:5000 --name back back
```
# Build frontend with Docker:
```
docker build . -f DockerFile --tag front
docker run -p 0.0.0.0:3000:3000 --name front front
```
