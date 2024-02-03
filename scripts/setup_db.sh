
set -a

source .env

CONTAINER_NAME="movie-rental-db"

docker pull postgres

docker run --name "$CONTAINER_NAME" \
    -e POSTGRES_DB="$DB_NAME" \
    -e POSTGRES_USER="$DB_USER" \
    -e POSTGRES_PASSWORD="$DB_PASSWORD" \
    -p "$DB_PORT":"$DB_PORT" \
    --network=bridge_for_movie_rental \
    -d postgres