# gRPC

- Framework desenvolvido pela google que tem o objetivo facilitar o processo de comunicação entre sistemas de uma forma extremamente rápida, leve, independente de linguagem.

## Em quais casos podemos utilizar?

- ideal para microserviços
- mobile, browsers e backend
- geração das bibliotecas de forma automática
- Streaming bidirecional utilizando HTTP/2

## Linguagens (Suporte oficial)

-  gRPC-GO
- gRPG-JAVA
- gRPG-C (c++, python, ruby, object c, php, c#, nodejs, dart, kotlin/jvm)

## RPC - Remote Procedure Call

Client  		-> 	Server
server.somar(a,b) -> 	func somar(int a, int b){}

## Protocol Buffers

- Protocol buffers are Google`s lagnuage-neutral, platform-neutral, extensible mechanism for serializing structured data - thing XML, but smaller...

## Protocol Buffers (P.B) vs JSON

- P.B < JSOn
- Arquivos binários != JSON
- Processo de serialização é mais leve (CPU) do que JSON
- Gasta menos recursos da rede
- Processo é mais veloz

P.B Format

```
syntax = "proto3";

message SearchRequest {
	string query = 1;
	int32 page_number = 2;
	int32 result_per_page = 3;
}
```

## HTTP/2

- nome original criado pela Google era SPDY
- lançado em 2015
- dados trafegados são binários e nao texto como no HTTP 1.1
- uitliza a mesma conexão TCP para enviar e receber dados do cliente e do servidor (Multiplex)
- Server Push
- Headers são comprimidos
- Gasta menos recurso de rede
- Processo é mais veloz

## gRPC - API "unary"

Client -> Request -> Server -> Response -> Client

## gRPC - API "Server streaming

Client -> Request -> Server -> Response, Response, Response -> Client

## gRPC - API "Client streaming"

Client -> Request, Request, Request -> Server -> Response -> Client

## gRPC - API "BI directional streaming"

CLient -> Request, Request, Request -> Server -> Response, Response, Response -> Client

## REST vs gRPC

### REST

- Text/JSON
- Unidirecional
- Alta latência
- Sem contrato (maior chance de erros)
- Sem suporte a streaming (Request/Response)
- Design pré-definido
- Bibliotecas de terceiro

### gRPG

- P.B
- Bidirecional e Assíncrono
- Baixa latência
- Contrato definido (.proto)
- Suporte a streaming
- Design é livre
- Geração de código (ñ precisa de bibliotecas de terceiro)

## Instalação das ferramentas gRPC

```
sudo apt install protobuf-compiler 
brew install protobuf #Mac, requer Homebrew instalado
```

```
go mod init github.com/<seu_user>/<repo_name>
```

```
go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

Para finalizar, temos que adicionar a pasta "/go/bin" no PATH do Linux para que tudo que seja instalado nesta pasta esteja disponível como comandos no terminal. Adicione no final do seu ~/.bash

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

Execute o comando abaixo para atualizar seu terminal:

```
source ~/.bashrc
```