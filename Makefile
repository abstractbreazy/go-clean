build:
	@docker-compose up --build
up:
	@docker-compose up -d --remove-orphans	
rebuild:
	@docker-compose up -d --remove-orphans --no-deps --force-recreate golang-app