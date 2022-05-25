# Iniciando um agente consul

Start consul agent

```
docker exec -it fullcycle_consul01 sh
consul agent -dev
```

Lista consuls

```
docker exec -it fullcycle_consul01 sh
consul members
```

Listar as informações dos nodes

```
curl localhost:8500/v1/catalog/nodes
```

Consulta a DNS

```
dig @localhost -p 8600
```

Consulta o consult (agent)

```
dig @localhost -p 8600 consul01.node.consul
```