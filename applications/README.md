### Docker Compose

- O docker-compose é uma ferramenta oficial do Docker que permite definir, configurar e gerenciar aplicações multicontainers (compostas por vários containers) de maneira simples e eficiente. Com o docker-compose, você pode descrever os serviços da sua aplicação, as redes que eles utilizam e os volumes de armazenamento persistente em um único arquivo de configuração YAML.

#### Principais Funcionalidades do Docker Compose
	    1.	Definir serviços
	    •	Descreve todos os containers necessários para sua aplicação
	   
        2.	Gerenciar redes
	    •	Cria e gerencia redes personalizadas para comunicação entre os containers.
	   
        3.	Configurar volumes
	    •	Permite persistir dados fora dos containers
	   
        4.	Orquestração
	    •	Possibilita executar múltiplos containers e gerenciando dependências entre os serviços.

##### Definição básica docker-compose.yaml

```bash
version: "3.9"
services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
  redis:
    image: redis:latest
```

#### Benefícios do Docker Compose
    • Automatização
    • Reprodutibilidade
    • Facilidade de uso
    • Escalabilidade

#### Lista de Comandos
- docker-compose ps
- docker-compose up
- docker-compose up -d
- docker-compose down
- docker-compose down --volumes
- docker-compose build
- docker-compose build --no-cache
- docker-compose stop
- docker-compose restart
- docker-compose logs
- docker-compose logs -f
- docker-compose up -d <service_name>
- docker-compose stop <service_name>
- docker-compose rm <service_name>
- docker-compose config
- docker-compose exec redis bash