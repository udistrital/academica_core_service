# academica_core_service

Servicio Go/Beego para `academica_core_service`.

Esta primera fase solo inicializa la aplicación y valida la conexión Oracle usando el mismo patrón de `moodle_service`: `database.BuildOracleConnectionString()` de `github.com/udistrital/utils_oas` y un pool global `sqlx.DB`.

## Requisitos

- Go 1.25
- Acceso a la base de datos Oracle
- Docker, opcional

## Configuración

Copie `.env.example` a `.env` y complete las variables locales. El archivo `.env` está ignorado por Git para evitar subir credenciales reales.

```sh
cp .env.example .env
```

Variables usadas por `conf/app.conf`:

```txt
ACADEMICA_CORE_SERVICE_RUN_MODE=dev
PARAMETER_STORE=
ACADEMICA_CORE_SERVICE_ORUSER=
ACADEMICA_CORE_SERVICE_ORPASS=
ACADEMICA_CORE_SERVICE_ORHOST=
ACADEMICA_CORE_SERVICE_ORPORT=
ACADEMICA_CORE_SERVICE_ORSERVICE=
ACADEMICA_CORE_SERVICE_ORSCHEMA=
```

## Ejecución local

```sh
go mod download
go run .
```

Si la conexión Oracle es correcta, Beego inicia en el puerto `8080`. Si hay error de credenciales, host, puerto, servicio o esquema, el servicio registra el error y termina.

## Docker

Primero compile el binario:

```sh
go build -o main .
docker compose up --build
```

## CI/CD

El repositorio sigue los lineamientos OATI para APIs Go/Beego:

- `.drone.yml`: ejecuta SonarQube, `go vet`, validación de `go fmt`, `golangci-lint`, compilación, publicación en ECR y despliegue ECS.
- `sonar-project.properties`: registra el proyecto `academica_core_service` y usa la estructura base Go.
- `Dockerfile`: empaqueta el binario `main` y `conf/app.conf` en una imagen distroless.
- `.dockerignore`: evita enviar `.env`, Git, logs y artefactos locales al contexto de Docker.

Secretos requeridos en Drone:

```txt
SONAR_HOST
SONAR_TOKEN
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_REGION
```

Convención de despliegue:

- `release/*`: publica tags `${DRONE_COMMIT:0:7}` y `release`, despliega servicio ECS `academica_core_service_test` en cluster `test`.
- `master`: publica tags `${DRONE_COMMIT:0:7}` y `latest`, despliega servicio ECS `academica_core_service_prod` en cluster `oas`.

## Estructura

- `main.go`: carga `.env`, construye la cadena Oracle, inicializa DB y arranca Beego.
- `conf/app.conf`: configuración Beego y variables Oracle.
- `models/db.go`: pool global `sqlx.DB`.
- `routers/`, `controllers/`, `requests/`, `responses/`: estructura base para fases siguientes.
