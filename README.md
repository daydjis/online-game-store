# Build backend with Docker:
cd back
docker build . -f DockerFile --tag back
docker run -p 0.0.0.0:5000:5000 --name back back
