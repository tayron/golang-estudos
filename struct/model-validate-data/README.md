# Validacao de Dados com validator

Este exemplo mostra como validar structs em Go usando a biblioteca `github.com/go-playground/validator/v10`.

O arquivo principal esta em [main.go](/mnt/hd-backup/Estudos/golang-estudos/struct/model-validate-data/main.go).

## O que o exemplo faz

- Define uma struct `RequestData` com tags `validate`.
- Cria uma instancia reutilizavel do `validator`.
- Registra traducao para portugues.
- Usa o nome do campo do `json` nas mensagens de erro.
- Demonstra um payload valido e outro invalido.

## Como funcionam as tags

Cada campo pode receber regras de validacao na tag `validate`.

Exemplos usados no projeto:

- `required`: campo obrigatorio.
- `len=14`: string deve ter tamanho exato.
- `numeric`: aceita apenas digitos.
- `oneof=J F`: valor deve ser um dos informados.
- `min=1` e `max=12`: limita intervalo numerico.
- `gt=0`: valor deve ser maior que zero.
- `datetime=02/01/2006`: string deve seguir esse formato de data.

Importante: no `validator`, o layout de data segue o padrao do Go. Para representar `dd/mm/aaaa`, use `02/01/2006`.

## Estrutura principal

### 1. Modelagem com tags

A struct define as regras direto no modelo:

```go
type RequestData struct {
    CnpjAdministradora string  `json:"cnpj_administradora" validate:"required,len=14,numeric"`
    Mes                int     `json:"mes" validate:"required,min=1,max=12"`
    Valor              float64 `json:"valor" validate:"required,gt=0"`
    Vencimento         string  `json:"vencimento" validate:"required,datetime=02/01/2006"`
}
```

### 2. Criacao do validator

O exemplo encapsula o `validator` em `RequestValidator` para evitar recriar tudo a cada validacao.

Isso melhora o exemplo porque:

- deixa o codigo mais organizado;
- reaproveita tradutor e configuracao;
- facilita usar a mesma logica em varios pontos da aplicacao.

### 3. Traducao das mensagens

O pacote `translations/pt` registra mensagens padrao em portugues.

Assim, em vez de mensagens tecnicas em ingles, a saida fica mais amigavel para quem consome a API ou para quem esta estudando.

### 4. Uso do nome do campo JSON

O exemplo registra uma funcao com `RegisterTagNameFunc` para exibir nomes como `cnpj_administradora` em vez de `CnpjAdministradora`.

Isso e especialmente util em APIs JSON, porque aproxima o erro do payload real recebido pelo backend.

## Como executar

```sh
cd struct/model-validate-data
go run .
```

## Saida esperada

O programa valida dois payloads:

- o primeiro deve passar com sucesso;
- o segundo deve listar varios erros traduzidos.

## Quando usar validator

`validator` e uma boa escolha quando voce quer:

- validar payloads de API;
- proteger regras basicas antes de persistir dados;
- centralizar regras simples de formato e obrigatoriedade;
- produzir mensagens consistentes de erro.

## Boas praticas

- Reutilize a instancia de `validator` em vez de criar uma nova toda vez.
- Use tags `json` junto com `RegisterTagNameFunc` para mensagens melhores.
- Separe validacoes simples via tag de validacoes complexas de negocio.
- Para regras mais especificas, crie validadores customizados.
