FROM postgres:14-alpine

ENV POSTGRES_PASSWORD=postgres 

COPY ./task_tsu.sql /docker-entrypoint-initdb.d/

