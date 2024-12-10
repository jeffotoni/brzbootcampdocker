# Usando a imagem oficial do Redis como base
FROM redis:latest

# Definindo a porta padrão do Redis
EXPOSE 6379

# Copiando o arquivo de configuração personalizado para dentro do container (opcional)
# COPY redis.conf /usr/local/etc/redis/redis.conf

# Iniciando o Redis com o arquivo de configuração, se necessário
# CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
CMD ["redis-server"]