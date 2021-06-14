## FeirasLivresAPI

API para cadastro de feiras livres da cidade de São Paulo.

### 📋 Pré-requisitos

Ferramentas: 

* [Golang](https://golang.org/doc/install)
* [Nodemon](https://nodemon.io/)
* [MySql](https://www.mysql.com/)

### 🛠️ Construído com

* [echo](https://echo.labstack.com/) - Framework Web
* [go mod](https://blog.golang.org/using-go-modules) - Dependência
* [viper](https://github.com/spf13/viper) - Configuração 
* [logrus](github.com/sirupsen/logrus) - Log
* [sqlx](https://github.com/jmoiron/sqlx) - Gereciamento de conexão de bancos relacionais
* [validator](github.com/go-playground/validator/v10) - Validador de structs
* [Mockgen](https://github.com/golang/mock) - Mock para testes

### Configuração de ambiente local

### Crie a seguinte estrutura em seu diretório raiz:

```
 root
 │  ├── go
 │      ├── src
 │          ├── github.com
 │              ├── unico

```

### Download do projeto

* Clonar repositorio dentro na pasta `unico`

### Configurar banco de dados MySql

* Configurar HOST e PASSWORD no arquivo `config.json` na raiz do projeto.

```
{
  "database": {
      "host": "root",
      "password": "123"
  }
}

```

### ⚙️ Executando os testes

* `make test`: executa os testes.


### 🚗 Rodando

* `make run`: comando padrão para executar o programa.
* `make run-watch`: comando com live reload

- OBS.: O Projeto rodara na porta `:7000`


### Importar dados das feiras 

- Na raiz do projeto executar o comando abaixo: 
`mysql -u <HOST> -p<PASSWORD> feira < DEINFO_AB_FEIRASLIVRES_2014.sql`

exemplo:
`mysql -u root -p123 feira < DEINFO_AB_FEIRASLIVRES_2014.sql`

obs. executar após rodar o projeto.

### 🗂 Arquitetura

### Descrição dos diretórios e arquivos mais importantes:

- `./api/v1`: Este diretório possui a configuração e registro de todos os sub-modulos.
- `./api/v1/v1.go`: Nesse arquivo está toda parte de registros dos sub-modulos que existem nesse diretório com o path `/v1/**`.
- `./model`: Este diretório possui todos os arquivos de modelos globais do projeto
- `./app`: Aqui se encontra todo o código que é utilizado para orquestração e regras de negôcio do serviço.
- `./store`: Aqui se encontra todo o código que é utilizado para interação com o banco de dados.
- `./db`: Diretório para criação de banco e tabela.
- `./utils`: Sub-modulos necessários para manutenção do projeto em geral.


### Endpoints

* **Create**
- `POST - http://localhost:7000/v1/feiras`
- body request:

```
{
    "id_feira": "123456",
    "long": "1111111111",
    "lat": "22222222222",
    "setcens": "xxxxxx",
    "areap": "xxxxxxxxxxx",
    "coddist": "xxxxxxxxxx",
    "distrito": "xxxxxxxxxx",
    "codsubpref": "xxxxxxxxx",
    "subprefe": "xxxxxxxxxxxxx",
    "regiao5": "xxxxxx",
    "regiao8": "xxxxx",
    "nome_feira": "xxxxxxx",
    "registro": "xxxxxxx",
    "logradouro": "xxxxxxxxxx",
    "numero": "xxxxxxx",
    "bairro": "xxxxxxxx",
    "referencia": "xxxxxxxxxxxxx"
}

```
- Resposta
- 201 = Status Created


* **Update**
- `PUT - http://localhost:7000/v1/feiras/:id`
- body request com campos obrigatorios:
```
{
    "long": "1111111111",
    "lat": "22222222222",
    "setcens": "xxxxxx",
    "areap": "xxxxxxxxxxx",
    "coddist": "xxxxxxxxxx",
    "distrito": "xxxxxxxxxx",
    "codsubpref": "xxxxxxxxx",
    "subprefe": "xxxxxxxxxxxxx",
    "regiao5": "xxxxxx",
    "regiao8": "xxxxx",
    "nome_feira": "xxxxxxx",
    "registro": "xxxxxxx",
    "logradouro": "xxxxxxxxxx"
}

```
- Resposta
- 200 = Status OK


* **ReadOne**
- `GET - http://localhost:7000/v1/feiras/:id`
- Resposta
- 200 = Status OK


* **Delete**
- `DELETE - http://localhost:7000/v1/feiras/:id`
- Resposta
- 204 = Status No Content


* **Search por DISTRITO**
- `GET - http://localhost:7000/v1/feiras//search`
- body request:

```
{
    "distrito": "exemplo"
}

```
- Resposta
- 200 = Status OK

OBS.: Para mais detalhes dos endpoints importar collection `Feira.postman_collection.json` e executar via postman


## Logs de erros armazenados no arquivo `logs.txt`

