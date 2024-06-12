#!/bin/sh

goose -dir migrations \
  postgres "user=postgres password=postgres host=localhost port=8432 database=go_db sslmode=disable"
  status
