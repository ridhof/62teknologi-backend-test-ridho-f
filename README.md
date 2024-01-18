# Bizsearch - Business Search API Server

An API server to manage and find businesses and reviews

## Requirement

[Read documents here](./documents/README.md)

## Database ERD

![Database ERD](./documents/database.svg)

## Todo

- [X] Design database schema
- [X] Design APIs
- [X] Database Migrations
- [X] Database Seeding
- [X] Database Query
- [X] APIs

## Setup

1. Run `pnpm install` to setup NodeJS depedencies
2. Run `pnpm prepare` to prepare husky for pre-commit hook validation
3. Run `cp .env.example .env` to copy env config file
4. Run `docker compose up -d` or `pnpm docker:up` to iniate PostgreSQL database instance with Docker

## Scripts

- `pnpm dbml:render` to render dbml into svg image
- `pnpm dbml:validate` to validate dbml definition
- `pnpm docker:up` to setup depedencies in docker such as postgresql database
- `pnpm docker:down` to set down dependencies in docker such as postgresql database by deleting persistent data
- `pnpm docker:psql` to access postgresql through psql cli
- `pnpm app:server:install` to install golang depedencies
- `pnpm app:server:dev` to run golang as a dev
- `pnpm app:server:build` to build golang code
- `pnpm app:server:start` to run the built code
- `pnpm app:server:test` to test the golang code
- `pnpm test` to test the whole application
- `pnpm prepare` to setup husky

## Depedencies

- [dbml-renderer](https://github.com/softwaretechnik-berlin/dbml-renderer) to render dbml file into a svg image
- [husky](https://typicode.github.io/husky/) to test and validate codebase before commit
