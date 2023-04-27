# [BEXS] Desafio Bexs - API DE ROLÊS
## Objetivo do Desafio
Fazer uma API simples seguindo os princípios da arquitetura limpa. O tema da API foi sorteado.

## Uso
Para rodar a aplicação execute o comando `docker compose up`.
Agora sua API esta pronta para ser utilizada :)

### Utilizando a API

	GET    localhost:8081/role/:id 		=> PEGAR UM ROLE
	GET    localhost:8081/roles		=> VER TODOS OS ROLES
	POST   localhost:8081/role		=> MARCAR UM ROLE
	PUT    localhost:8081/role/:id		=> EDITAR UM ROLE CADASTRADO
	DELETE localhost:8081/role/:id		=> APAGAR UM ROLE

## JSON de cadastro
	{
    "nome": "NOME DO ROLE",
    "local": "LOCAL DO ROLE",
    "cidade": "CIDADE DO ROLE",
    "horario": "HORARIO DO ROLE",
    "dia": "DIA DO ROLE",
    "tipo do role": "TIPO DO ROLE"
	}
## Exemplo do JSON populado
	{
	"nome": "Aniversario da Amanda",
    "local": "Fat Cow | R. Iaiá, 173 - Itaim Bibi, São Paulo - SP, 04542-060",
    "cidade": "São Paulo",
    "horario": "18h - 22h",
    "dia": "01-12-2022",
    "tipo do role": "Aniversario"
	}

<font color="red">O campo dia deve ser preenchiido no formato `dia-mes-ano` | `25-11-2022`</font>
