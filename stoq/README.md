# A Stoq Api

> Requisitos do projeto:

- Go Lang >= 1.17
- NodeJS 16

As demais dependências estão no arquivo go.mod

> Para preparar o ambiente e compilar:
```bash
$ go get
$ # Compilar servidor HTTP
$ go build -o stoq cmd/stoq/main.go
$ # Ou compilar para outra plataforma ex: windows
$ GOOS=windows GOARCH=amd64 go build -o stoq64.exe main.go
$ # build modo production
$ go build -ldflags "-s -w" . 
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
