# Criando nosso cluster

Definindo o server

```
consul agent -server -bootstrap-expect=3 -node={fullcycle_consulserver01} -bind={172.19.0.4} -data-dir=/var/lib/consul -config-dir=/etc/consul.d
```

Juntando os servers

```
consul join {172.19.0.3}
```