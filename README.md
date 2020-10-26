# Transactions Routine
Transactions Routine é uma aplicação para simular transações bancárias, como compra, saque e pagamento.

A aplicação foi desenvolvida utilizando a linguagem GoLang e conceitos de Clean Architecture descritos por Robert C. Martin em seu livro "Clean Architecture: A Craftsman's Guide to Software Structure". Tais conceitos de clean architeture permitem escrever um código organizado, encapsulando a lógica de negócio e mantendo componentes desacoplados, o que torna possível desenvolver uma aplicação testável, independente de interface, banco de dados ou qualquer agente externo.

## Dependências
- **Docker:** É uma plataforma aberta para criação, execução e publicação (deploy) de containers. Um Container é a forma de empacotar sua aplicação e suas dependências (bibliotecas) de forma padronizada. [Mais Informações](https://www.docker.com/)

- **Docker Compose:** É uma ferramenta para definir e rodar aplicações multi-container Docker. Com o docker composer é possível utilizar um arquivo YAML para configurar todos os serviços de sua aplicação e, com um único comando, criar e executar todos os serviços definidos. [Mais Informações](https://docs.docker.com/compose/)
  
- **Make:** É um utilitário que compila automaticamente programas e bibliotecas do arquivo fonte através da leitura de instruções contidas em arquivos denominados Makefiles, que especificam como obter o programa de destino. [Mais Informações](https://pt.wikipedia.org/wiki/Make)

## Configuração
Por motivos de segurança, dados de configuração da aplicação devem ficar contidos em um arquivo **.env** na raiz do projeto. Para facilitar a execução existe um arquivo **.env.example** com um exemplo de dados de configuração.

## Como Executar
1. Executar ***make configure*** para gerar um arquivo **.env** na raiz do projeto utilizando o arquivo de exemplo.
2. Executar ***make start*** para subir todos os containeres Docker necessários para executar a aplicação.

## Testes
Após subir a aplicação utilizando oa comandos mostrados na seção anterior, executar ***make test*** para rodar todos os testes existentes.

## REST API

A aplicação expõe uma API REST HTTP em um porta configurável. Por padrão 8080.

No caso de erro, a API responde com um Status Code HTTP apropriado, e em determinados casos, o payload conterá mais informações sobre o erro.
Exemplo:
```
{
    "error": "Invalid accountID"
}
```

## Endpoints
### `GET /accounts/{accoundID}`

Endpoint responsável por buscar uma conta cadastrada no sistema.

Parâmetros:
```
accountID: integer
```

Response
```json
HTTP/1.1 200 OK
Date: Mon, 26 Oct 2020 01:17:32 GMT
Content-Type: application/json
Content-Length: 51

{
    "account_id": 2,
    "document_number": "1234567890021"
}
```

### `POST /accounts`

Endpoint responsável por cadastrar uma nova conta para um cliente à partir de um número de documento.

Headers:
```
Content-Type: application/json
```

Request Body:
```json
{
    "document_number": "1234567890021"
}
```

Response:
```json
HTTP/1.1 201 Created
Date: Mon, 26 Oct 2020 01:20:43 GMT
Content-Type: application/json
Content-Length: 51

{
    "account_id": 4,
    "document_number": "1234567890022"
}
```

### `POST /transactions`

Endpoint responsável por registrar uma nova transação para uma conta cadastrada, informando também o valor e o tipo da operação (ver tabela abaixo).


| ID | Descrição        | Tipo da Operação |
|----|------------------|------------------|
|  1 | COMPRA À VISTA   | Débito           |
|  2 | COMPRA PARCELADA | Débito           |
|  3 | SAQUE            | Débito           |
|  4 | PAGAMENTO        | Crédito          |

**Atenção:** Todas as operações do tipo débito serão salvas com valor negativo.

Headers:
```
Content-Type: application/json
```

Request Body:
```json
{
    "account_id": 1,
    "operation_type_id": 1,
    "amount": 100.50
}
```

Response:
```json
HTTP/1.1 201 Created
Date: Mon, 26 Oct 2020 01:31:56 GMT
Content-Type: application/json
Content-Length: 111

{
    "account_id": 2,
    "operation_type": "COMPRA A VISTA",
    "amount": -100.5,
    "created_at": "2020-10-26T01:31:56.3622751Z"
}
```
