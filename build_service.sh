IMAGE_NAME=social-todo-service
CACHED_BUILD=$1

if [[ -n "$CACHED_BUILD" ]]; then
  echo "Docker building cached image..."
  docker rmi ${IMAGE_NAME}-cached ${IMAGE_NAME}
  docker build -t ${IMAGE_NAME}-cached -f Dockerfile-cache .
fi

echo "Docker building main image..."
docker build -t social-todo-service:latest -f Dockerfile-multi-stage .

echo "Done!!"