.PHONY: sqlc
sqlc:
	@sqlc generate internals/db/sqlc.yaml