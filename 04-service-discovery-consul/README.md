# Service Discovery e Consul

## Agent, Client e Server

### Agent

Processo que fica sendo executado em todos nós do cluster. Pode estar sendo executado em Cliente Mode ou Server Mode

### Client

Registra os serviços lcoalmente, health check, encaminha as informações e configurações dos serviços para o Server

### Server

Mantém o estado do cluster, registra os serivços, Membership (quem é client e quem é server), retorno de queries (DNS ou API), troca de informações entre datacenters, etc.

### Dev Mode

- NUNCA utilize em produção
- Teste de features, exemplos
- Não escala
- Registra tudo em memória

## Código fonte

gRPC: https://github.com/codeedu/fc2-grpc
graphql: https://github.com/codeedu/fc2-graphql