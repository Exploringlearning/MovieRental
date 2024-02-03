
set -a

source .env

TAG="movie-rental"

docker rm $TAG

docker build -t "$TAG" .

docker network create bridge_for_movie_rental

docker run \
  --name="$TAG" \
  --network=bridge_for_movie_rental \
  -e DB_HOST="$DB_HOST" \
  -e DB_PORT="$DB_PORT" \
  -e DB_USER="$DB_USER" \
  -e DB_PASSWORD="$DB_PASSWORD" \
  -e DB_NAME="$DB_NAME" \
  -p 8080:8080 \
  -t "$TAG"
