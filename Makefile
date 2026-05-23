db-push:
	@set -a; source .env; set +a; \
	go run github.com/steebchen/prisma-client-go db push

debug:
	@$(ENV_LOAD) echo "DATABASE_URL=$$DATABASE_URL"