php() {
    docker-compose exec -w /var/www/html -u developer php php "$@"
}

composer() {
    docker-compose exec -w /var/www/html -u developer php composer "$@"
}
