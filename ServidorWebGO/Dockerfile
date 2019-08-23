# Seleciona a imagem base
FROM golang:1.12.6-alpine3.10 as build

# Seta a pasta do go
ENV GOPATH /go

# seta a variável de ambiente
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Cria as pastas do go
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
&& chmod -R 777 "$GOPATH" 

# Copia o arquivo
COPY ./ $GOPATH/src

# Seta o diretório de trabalho
WORKDIR $GOPATH

# Expõe a porta
EXPOSE 8000

# Executa o programa
CMD go run ./src/main.go