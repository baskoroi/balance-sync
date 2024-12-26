# Balance sync demo

Synchronization between two data centers (each with their own Golang API and Postgres DB), but within the localhost environment.

## Setup

> BEFORE running the APIs, please run the following commands first!

> This setup assumes a Debian/Ubuntu OS installation. Adjust accordingly for the other distros.

### Tools

- Golang 1.23
- Postgres 14.15
- Bucardo (just install in localhost, no viable Docker image)
- Kafka

### DB initialization

```sh
$ docker run --name balance-sync-bucardo -e POSTGRES_USER=bucardo -e POSTGRES_PASSWORD=bucardo -p 54320:5432 -d postgres:14.15
$ docker run --name balance-sync-postgres-1 -e POSTGRES_USER=gorm -e POSTGRES_PASSWORD=gorm -p 54321:5432 -d postgres:14.15
$ docker run --name balance-sync-postgres-2 -e POSTGRES_USER=gorm -e POSTGRES_PASSWORD=gorm -p 54322:5432 -d postgres:14.15
```

Then run `CREATE DATABASE balance_data;` on each "data center". Do not worry about migrations, gorm will automigrate when the API runs.

## Bucardo installation

Try installing Bucardo not from your package manager, but locally via [bucardo's website](https://bucardo.org/Bucardo/).

In your Debian/Ubuntu env, simply run the following commands: 

```sh
$ sudo chmod +x /etc/bucardorc
$ sudo chown $USER /etc/bucardorc # to enable bucardo commands be run without `sudo`, hehe...
$ sudo vi /etc/bucardorc          # or whatever text editor you have...
```

Inside the `/etc/bucardorc` file, configure Bucardo to use port 54320 for their Postgres DB.

Then run `bucardo install`.
