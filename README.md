# Penca Backend

### Opcion 1 - Levantar ambiente con docker-compose

####  1.a - Requisitos
Configurar el archivo docker.env en la raiz del proyecto

```bash
SERVER_ADDRESS=0.0.0.0
DB_URI=usr_db:pass@tcp(dbhost:port)/db_name?tls=true&tls=skip-verify&parseTime=true
JWT_KEY=*************************
AMQP_URI="amqp://user:pass@host:port/"
```

#### 1.b - Levantar ambiente
Parados en la raiz del proyecto
```bash
docker-compose -f docker-compose-bkd.yml up     
```

### Opcion 2 - Levantar ambiente con go

####  2.a - Requisitos
Configurar el archivo app.env en la raiz del proyecto

```bash
SERVER_ADDRESS=0.0.0.0
DB_URI=usr_db:pass@tcp(dbhost:port)/db_name?tls=true&tls=skip-verify&parseTime=true
JWT_KEY=*************************
AMQP_URI="amqp://user:pass@host:port/"
```

#### 2.b - Levantar ambiente
Parados en la raiz del proyecto ejecutar el siguiente comando
```bash
go run src/main.go
```
