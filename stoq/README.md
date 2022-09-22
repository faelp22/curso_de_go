# A Stoq API

Trabalho de casa, sistema, curso (tcs_curso)

> Requisitos do projeto:

- Go Lang >= 1.17
- NodeJS 16

As demais dependências estão no arquivo go.mod e package.json

- https://go.dev/dl/
- https://nodejs.org/en/download/
- https://classic.yarnpkg.com/lang/en/docs/install/#windows-stable
- https://quasar.dev/start/quasar-cli

> Build do Front-End Quasar:
```bash
# entre na pasta webui
$ yarn install

# Instale o Quasar CLI
$ yarn global add @quasar/cli

# OBS antes de fazer o build veja se a variável está com esse valor [apiPath: '/' // Prod] dentro do arquivo webui\src\config\index.js

# Faça o build do Front-End
$ quasar build

```

> Build do Back-End Go:
```bash
# Baixando as dependências
$ go mod tidy

# Compilar servidor HTTP
$ go build -o stoq cmd/stoq/main.go

# Ou compilar para outra plataforma ex: windows
$ GOOS=windows GOARCH=amd64 go build -o stoq64.exe main.go

# build modo production
$ go build -ldflags "-s -w" .
# ou
$ go build -ldflags "-s -w" cmd/stoq/main.go
$ go build -ldflags "-s -w" -o stoq cmd/stoq/main.go
```
## Opções de execução
- SRV_PORT (Porta padrão 8080)
- SRV_MODE (developer ou production)
- SRV_WEB_UI (se true ativa a WEB UI)

> Exemplo de Uso:
```bash
$ ./main.exe
$ SRV_PORT=8080 SRV_MODE=developer SRV_WEB_UI=true ./main.exe
$ SRV_PORT=9090 SRV_MODE=production SRV_WEB_UI=false ./main.exe
```

> Acesse:
- http://localhost:8080/webui/
- http://localhost:8080/api/v1/products

OBS é necessário usar o Postman para consumir a API importe o arquivo tcs_curso/doc/Curso de GO API.postman_collection.json

- https://www.postman.com/downloads/
