#!/bin/bash
source .env

sleep 2 && goose -dir "${MIG_DIR}" postgres "${PG_DSN}" up -v