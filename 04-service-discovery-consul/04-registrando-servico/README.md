# Registrando serviço

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

Consultar os serviços

```
apt -U add bind-tools
dig @localhost -p 8600 SRV
dig @localhost -p 8600 nginx.service.consul
```

Vincular automaticamente a um IP Server

```
consul agent -bind={172.19.0.4} -data-dir=/var/lib/consul -config-dir=/etc/consul.d -retry-join={ip server consul}  -retry-join={ip server consul}
```

## Validando check

```
consul reload
apk add nginx
nginx # validar se rodou
curl localhost # retorno do nginx
```

Caso retorno 404 é necessário configurar o arquivo de configuração

```
apk add vim
mkdir /usr/share/nginx/html -p
vim /usr/share/nginx/html/index.html # add Hello
vim /etc/nginx/config.d/default.conf # add 'root /usr/share/nginx/html'
nginx -s reload
```

## Sincronizando servers via arquivo

Criar a pasta servers e seu respectivo arquivo

```
docker exec -it consulserver01 sh
consul agent -config-dir=/etc/consul.d
```

## Trabalhando com criptografia

```
docker exec -it consulserver01 sh
consul keygen
```