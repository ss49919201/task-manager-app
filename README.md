# Operate DB

`docker exec -it mysql bin/bash`

`mysql -u user -p`

`show databases;`

`show tables;`

# Request [cURL]

## Create user

```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"太郎"}' localhost:12345/users
```