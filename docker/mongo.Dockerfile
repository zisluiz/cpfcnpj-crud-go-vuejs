FROM mongo:4.4
COPY ./docker/initial-data.js /docker-entrypoint-initdb.d/

EXPOSE 27017