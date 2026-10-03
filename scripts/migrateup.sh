#!/usr/bin/env bash

cd sql/schema
goose turso $DATABASE_URL up
