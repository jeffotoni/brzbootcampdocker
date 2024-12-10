# Usando a imagem oficial do Python 3.9
FROM python:3.9-slim

# Definindo o diretório de trabalho dentro do container
WORKDIR /app

# Copiando os arquivos necessários para o diretório de trabalho
COPY requirements.txt requirements.txt
#COPY . .
COPY . /app

# Instalando as dependências
RUN pip install --no-cache-dir -r requirements.txt

# Expondo a porta 8080
EXPOSE 8080

# Comando para iniciar a aplicação
# CMD ["python", "main.py"]
CMD ["gunicorn", "-w", "4", "-b", "0.0.0.0:8080", "main:app"]