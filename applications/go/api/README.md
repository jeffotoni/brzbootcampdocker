### CRUD api rEST em Go

Este é uma api feita em Go, de forma minimalistica para aprendizado.
A porta do serviço é 8080.

##### Instalar Go ?
A reposta é não.

Isto mesmo não vamos fazer isto em sua máquina local.

Para este projeto, não vamos precisar instalar e nem configurar a linguagem Go na máquina. Vamos usar Docker para isto.

Desta forma iremos *entender de verdade o que é o Docker.*

##### Endpoints Disponíveis
---
| METHOD   |             ENDPOINT    
|----------|:----------------------:|
| POST     |  /v1/user              |
| GET      |  /v1/user              |
| GET      |  /v1/user/{id}         |
| PUT      |  /v1/user/{id}         |
| DELETE   |  /v1/user/{id}         |

##### PAYLOAD
```bash
{
  "id": 1,  
  "user": "jeff@gmail.com"
}
```

##### POST
```bash
$ curl -i -XPOST -H "Content-type:application/json" 
localhost:8080/v1/user -d '{"id":1, "name":"Lucas"}'
```

**OUTPUT:**
```bash
HTTP/1.1 201 Created
Date: Mon, 09 Dec 2024 22:59:34 GMT
Content-Length: 42
Content-Type: text/plain; charset=utf-8

{"message":"Usuário criado com sucesso"}
```

##### GET - Listar todos usuarios
```bash
$ curl -i -XGET -H "Content-type:application/json" 
localhost:8080/v1/user
```

**OUTPUT:**
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 09 Dec 2024 23:05:17 GMT
Content-Length: 62

[{"id":1,"name":"jeffotoni 1"},{"id":2,"name":"jeffotoni 2"}]
```

##### GET - Listar usuários específico
```bash
$ curl -i -XGET -H "Content-type:application/json" 
localhost:8080/v1/user/1
```

**OUTPUT:**
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 09 Dec 2024 23:05:17 GMT
Content-Length: 62

{"id":1,"name":"jeffotoni 1"}
```

##### PUT
```bash
$ curl -i -XPUT -H "Content-type:application/json" 
localhost:8080/v1/user/1
```

**OUTPUT:**
```bash
HTTP/1.1 200 OK
Date: Mon, 09 Dec 2024 23:08:41 GMT
Content-Length: 46
Content-Type: text/plain; charset=utf-8

{"message":"Usuário atualizado com sucesso"}
```

##### DELETE
```bash
$ curl -i -XDELETE -H "Content-type:application/json" 
localhost:8080/v1/user/1
```

**OUTPUT:**
```bash
HTTP/1.1 200 OK
Date: Mon, 09 Dec 2024 23:09:18 GMT
Content-Length: 44
Content-Type: text/plain; charset=utf-8

{"message":"Usuário deletado com sucesso"}
```

##### BUILD DOCKER V0
```bash
$  docker build --no-cache -f DockerfileV0 -t jeffotoni/brzapibootcampdockerv0 .
```

##### BUILD DOCKER V1
```bash
$  docker build --no-cache -f DockerfileV1 -t jeffotoni/brzapibootcampdockerv1 .
```

##### BUILD DOCKER V2
```bash
$  docker build --no-cache -f DockerfileV2 -t jeffotoni/brzapibootcampdockerv2 .
```

#####  DOCKER V3
```bash
$  docker build --no-cache -f DockerfileV3 -t jeffotoni/brzapibootcampdockerv3 .
```

#####  DETALHAR OS COMANDOS DO DOCKER BUILD
 - **docker build**
Este comando inicia o processo de construção de uma imagem Docker com base no Dockerfile especificado.

 - **--no-cache**
Esta opção garante que o Docker não usará o cache de etapas anteriores para construir a imagem. 

 - **-f DockerfileV3**
Especifica o nome do arquivo Dockerfile a ser usado durante o build.

 - **-t jeffotoni/brzapibootcampdockerv3**
Define a tag da imagem Docker.
    - *jeffotoni: É o namespace ou repositório.*
    - *brzapibootcampdockerv3: É o nome da imagem.*
    - *Por que usar uma tag: Facilita o gerenciamento e versionamento das imagens criadas.*

 - **. (ponto)**
Este parâmetro indica o contexto de build, que é o diretório onde estão os arquivos que o Docker utilizará durante a construção da imagem.

#####  LISTAR IMAGENS
```bash
$  docker images
➜  api git:(main) ✗ docker images
REPOSITORY                         TAG       IMAGE ID       CREATED             SIZE
jeffotoni/brzapibootcampdockerv3   latest    a18b8697f4bb   23 minutes ago      8.28MB
jeffotoni/brzapibootcampdockerv2   latest    8c5504c4416f   About an hour ago   11.3MB
jeffotoni/brzapibootcampdockerv0   latest    24fbd7f54e1f   2 hours ago         7.45MB
jeffotoni/brzapibootcampdockerv1   latest    ecae63e51488   2 hours ago         12MB
```

#####  SUBIR CONTAINER
```bash
$  docker run -d --rm -p 8080:8080 --name brzapibootcampdockerv3 jeffotoni/brzapibootcampdockerv3
81c270391cd56633dfb0c2ddcd8b6d531995e351eda0acdcc5320290f11aaa57
```

#####  DETALHAR OS COMANDOS DO DOCKER RUN

 - **docker run**
Este comando cria e inicia um contêiner baseado em uma imagem Docker especificada.

 - **-d**
Executa o contêiner em modo desacoplado (detached mode).

 - **--rm**
Remove o contêiner automaticamente após sua parada.

 - **-p 8080:8080**
Mapeia a porta do host para a porta do contêiner
    **- 8080:8080:**
	•	O número antes do : é a porta no host (máquina local).
	•	O número após o : é a porta no contêiner.

 - **--name brzapibootcampdockerv3**
Define o nome do contêiner como brzapibootcampdockerv3.

 - **jeffotoni/brzapibootcampdockerv3**
Especifica a **imagem** que será usada para criar o contêiner.

#####  DOCKER LISTAR CONTAINERS
```bash
$  docker ps -a
CONTAINER ID   IMAGE                              COMMAND                CREATED         STATUS         PORTS                    NAMES
81c270391cd5   jeffotoni/brzapibootcampdockerv3   "/usr/local/bin/api"   4 minutes ago   Up 4 minutes   0.0.0.0:8080->8080/tcp   brzapibootcampdockerv3
```

#####  DOCKER PARAR CONTAINER
```bash
$  docker stop brzapibootcampdockerv3
brzapibootcampdockerv3
```

```bash
$  docker stop 81c270391cd5
81c270391cd5
```

#####  DOCKER PARAR CONTAINER
```bash
$  docker stop brzapibootcampdockerv3
brzapibootcampdockerv3
```

#####  DOCKER REMOVER CONTAINER
```bash
$  docker run -d -p 8080:8080 --name brzapibootcampdockerv3 jeffotoni/brzapibootcampdockerv3
81c270391cd56633dfb0c2ddcd8b6d531995e351eda0acdcc5320290f11aaa57
```
```bash
$  docker stop brzapibootcampdockerv3
brzapibootcampdockerv3
```

##### DOCKER STOP EM TODOS CONTAINERS EM EXECUÇÃO
```bash
$ docker stop $(docker ps -q)
```

##### DOCKER START
```bash
$ docker start brzapibootcampdockerv3
brzapibootcampdockerv3
```

##### DOCKER RESTART
```bash
$ docker restart brzapibootcampdockerv3
brzapibootcampdockerv3
```

##### DOCKER INSPECT
```bash
$ docker inspect brzapibootcampdockerv3
brzapibootcampdockerv3
```

##### DOCKER RM EM CONTAINERS
```bash
$  docker rm brzapibootcampdockerv3
brzapibootcampdockerv3
```

##### DOCKER RM EM TODOS CONTAINERS EM EXECUÇÃO E PARADOS
```bash
$ docker rm $(docker ps -aq)
```

##### DOCKER VISUALIZANDO LOGS
```bash
$ docker logs brzapibootcampdockerv3
{"time":"2024-12-09T23:53:06.577971792-03:00","level":"ERROR","msg":"Usuário já existe","error":"usuario já existe"}
{"time":"2024-12-09T23:53:06.577994-03:00","level":"ERROR","msg":"Usuário já existe","error":"usuario já existe"}
{"time":"2024-12-09T23:53:06.578020292-03:00","level":"ERROR","msg":"Usuário já existe","error":"usuario já existe"}
{"time":"2024-12-09T23:53:06.578030708-03:00","level":"ERROR","msg":"Usuário já existe","error":"usuario já existe"}
```

##### DOCKER VISUALIZANDO LOGS
```bash
$ docker logs -f brzapibootcampdockerv3
{"time":"2024-12-09T23:53:06.577971792-03:00","level":"ERROR","msg":"Usuário já existe","error":"usuario já existe"}
{"time":"2024-12-09T23:53:06.577994-03:00","level":"ERROR","msg":"Usuário já existe","error":"usuario já existe"}
```

##### DOCKER SYSTEM DF
```bash
$ docker system df
TYPE            TOTAL     ACTIVE    SIZE      RECLAIMABLE
Images          4         1         5.427GB   5.415GB (99%)
Containers      1         1         12.29kB   0B (0%)
Local Volumes   0         0         0B        0B
Build Cache     200       0         6.299GB   6.299GB
```

##### DOCKER SYSTEM PRUNE
```bash
$ docker system prune -a
WARNING! This will remove:
  - all stopped containers
  - all networks not used by at least one container
  - all dangling images
  - unused build cache

Are you sure you want to continue? [y/N]
q15368obxuh92isx5uh85gh3a
su8rcwk5rayov5ozd5yomya05
jz0wgh427k2gwd6oamcsnup0f
9vil617h2csj74esth71ic5up
xniesfmhekk1pf7sfq18zly0h
sbp8o41xorkcvq0xdt9s2sojj
r4fx7xa255ki9dnd6ju0752lh
wlkd4xc2t2gf5dd7xk9cvmi44
urq6zcfpmkx7bvrc8cc6px0oj
xf4axi2u2qwlevakw4wsgqpfq

Total reclaimed space: 6.299GB
```

##### APAGAR IMAGENS
```bash
$ docker rmi jeffotoni/brzapibootcampdockerv3
Untagged: jeffotoni/brzapibootcampdockerv3:latest
Deleted: sha256:a18b8697f4bbb2742dd09c400b85457fd2e2fa0918034b45b6ae75244899f5e7
```

##### ENVIAR IMAGEM PARA REGISTRY

Vamos usar Docker Hub [Entrar por aqui](https://hub.docker.com)

O Docker Hub é um repositório de imagens de contêiner baseado em nuvem, mantido pela Docker Inc. Ele é a plataforma padrão onde desenvolvedores e equipes armazenam, compartilham e distribuem imagens de contêineres Docker.

#####  Logando no Docker Hub
```bash
$ docker login
Authenticating with existing credentials...
Login Succeeded
```

#####  Enviando Imagem
```bash
$ docker push jeffotoni/brzapibootcampdockerv3:1.0
```

#####  Baixando Imagem
```bash
$ docker pull jeffotoni/brzapibootcampdockerv3:1.0
```

#####  Alterando Tag
```bash
$ docker tag jeffotoni/brzapibootcampdockerv3:1.0 jeffotoni/brzapibootcampdockerv3:2.0
```