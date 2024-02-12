#!/bin/bash
source .env
goose -dir "${MIG_DIR}" postgres "${PG_DSN}" version
sleep 2 && goose -dir "${MIG_DIR}" postgres "${PG_DSN}" status
sleep 2 && goose -dir "${MIG_DIR}" postgres "${PG_DSN}" up

