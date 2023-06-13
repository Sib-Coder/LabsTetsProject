# Заметки по работе с бд

## Голая бд
```bash
docker run --name name-postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
```
## С таблицей и данными 
Для того чтобы работать с контейнером установите docker.

Зайтите в папку с Dockerfile и пропишите
```bash
sudo docker build -t postgres_db .
```
Вы собрали образ. Теперь создайте и запустите контейнер.
```bash
sudo docker run -d   -p 5432:5432 --name tsupsql  postgres_db
```
tsupsql - имя контенейра.
Ваша Бд запушена на порту 5432.