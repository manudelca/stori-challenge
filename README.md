# stori-challenge

Para correr el challenge es necesario:
```
docker
docker-compose
make
```

Para la lógica de mail, es necesario hacer setup del environment creando un archivo
```
stori.env
```
Tomando como base el archivo `.env`

Para ello es necesario contar con `app_password` para permitir el forwarding del email en nombre de esta aplicación

Para levantar correr:
```
make docker-compose-up
```

Para ver los logs generados:
```
make docker-compose-logs
```