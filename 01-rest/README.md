# REST

- Muitos desenvolvedores "já sabem trabalhar com REST"
- Representational state of transfer
- Surgiu em 2000 por Roy Fielding em uma dissertação de doutorado
- Simplicidade
- Stateless
- Cacheável

## Nível de Maturidade (RIchardson Maturity Model)

### Nível 0: The Swamp of POX
### Nível 1: Utilização de resources

Verbo -> URI -> Operaçãos

GET		->	/products/1	->	Buscar
POST		->	/products	->	Inserir
PUT		->	/products/1	->	Alterar
DELETE	->	/products/1	->	Remover

### Nível 2: Verbos HTTP

GET		-> Recuperar Informação
POST		-> Inserir
PUT		-> Alterar
DELETE	-> Remover

### Nível 3: HATEOAS (Hypermedia as the Engine of Application State)

HTTP/1.1 200 OK
Content-Type: application/vnd.acme.account+json
Content-Lenght: ...
{
	"account": {
		"account_number": 1234,
		"links": {....}
	}
}

## Uma boa API REST

- Utilizar URIs únicas para serviços e itens que expostos para esses serviços
- Utiliza todos os verbos HTTP para realizar as operações em seus recursos, incluindo caching
- Provê links relacionais para os recurso s exemplificando o que pode ser feito

### HAL, Collection+JSON e3 Siren

- JSON não provê um padrão de hipermídia para realizar a linkagem
- HAL: Hypermedia Application Language
- Siren

### HTTP Method Negotiation

- HTTP possui um outro método: OPTIONS. Esse método nos permite informar quais métodos são permitidos ou não em determinado recurso.

OPTIONS /api/product HTTP/1.1
Host: domain.com.br

Resposta pode ser:

HTTP/1.1 200 OK
Allow: GET, POST

Caso envie a requisição em outro formato

HTTP/1.1 405 Not Allowed
Allow: GET, POST

### Content Negotiation

O processo de content negotiation é baseado na requisição que o client está fazendo para o server. Nesse caso ele solicita o que e como ele quer a resposta.
O server então retornará ou não a informação no formato desejado.

#### Accept Negotiation

- Client solicita a informação e o tipo de retorno pelo server baseado no media type informado por ordem de prioridade

GET /product
Accept: application/json

Resposta pode ser o retorno dos dados ou:
HTTP/1.1 406 Not Acceptable

#### Content-Type Negotiation

- Client solicita a informação e o tipo de retorno pelo server baseado no media type informado por ordem de prioridade

POST /product HTTP/1.1
Accept: application/json
Content-Type: application/json
{
	"name": "Product 1"
}

Caso o servidor não aceite o cotnent type, ele poderá retornar

HTTP/1.1 415 Unsupported Media Type