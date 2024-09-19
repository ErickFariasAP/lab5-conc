# Laboratório 05 - Programação Concorrente

## Grupo

| Nome                            | Matrícula |
|---------------------------------|-----------|
| Julio Hsu                       | 120110370 |
| Lucas Emanuel Silva Melo        | 121110016 |
| Douglas Alves de Sousa          | 121111728 |
| Jackson Alves da Silva Souza    | 121110759 |
| Erick Farias de Almeida Pereira | 121110373 |
 
## 1. Como usar (servidor)

```bash
go run *.go
```

## 2. Como usar (client)

Os comandos a seguir devem ser executados dentro da pasta client.

### 2. 1. Tracking de arquivos

Para se conectar ao servidor e fazer ele ser capaz de reconhecer os arquivos do client que estão em uma pasta específica, use o comando:

```bash
go run *.go ../dataset
```

### 2. 2. Search

Para receber a informação de quais ips conectados ao servidor tem um arquivo especifico use o comando:

```bash
go run *.go search <HASH_DO_ARQUIVO>
```

## 3. Conexão entre máquinas diferentes

Caso seja desejável conectar múltiplas máquinas é necessário mudar o arquivo /config/config.go para conter as informações do servidor.