# Criando nosso client

Configurando os servers

```
consul agent -server -bootstrap-expect=3 -node={fullcycle_consulserver01} -bind={172.19.0.4} -data-dir=/var/lib/consul -config-dir=/etc/consul.d
```

Configurando o client

```
consul agent -bind={172.19.0.4} -data-dir=/var/lib/consul -config-dir=/etc/consul.d
```

Juntando os servers/clients

```
consul join {172.19.0.3}
```