.PHONY: php artisan composer composerexec up down build restart logs

php:
	docker compose exec -w /var/www/html -u developer php php $(filter-out $@,$(MAKECMDGOALS))

artisan:
	docker compose exec -w /var/www/html -u developer php php artisan $(filter-out $@,$(MAKECMDGOALS))

composer:
	docker compose exec -w /var/www/html -u developer php composer $(filter-out $@,$(MAKECMDGOALS))

composerexec:
	docker compose exec -w /var/www/html -u developer php composer exec $(filter-out $@,$(MAKECMDGOALS))

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

restart:
	docker compose restart

logs:
	docker compose logs -f

# 👇 این خط کل مشکل رو حل می‌کنه
%:
	@:
