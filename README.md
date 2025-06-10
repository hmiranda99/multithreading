# 🚀 Desafio de Multithreading em Go

Este projeto implementa uma solução em Go que utiliza multithreading para buscar informações de CEP em duas APIs diferentes simultaneamente, retornando o resultado da API mais rápida.

## 🌐 APIs Utilizadas

1. BrasilAPI: https://brasilapi.com.br/api/cep/v1/{cep}
2. ViaCEP: http://viacep.com.br/ws/{cep}/json/

## 📋 Requisitos do Desafio

- ⚡ Realizar requisições simultâneas para ambas as APIs
- 🏃 Aceitar a resposta da API mais rápida e descartar a mais lenta
- 📝 Exibir os dados do endereço e qual API retornou a resposta
- ⏱️ Limitar o tempo de resposta em 1 segundo, exibindo erro de timeout caso necessário

## 🛠️ Tecnologias Utilizadas

- 🦫 Go (Golang)
- 🔄 Goroutines para concorrência
- 📨 Channels para comunicação entre goroutines
- ⏰ Context para controle de timeout
- 🌍 HTTP client para requisições

## 🚀 Como Executar

1. ✅ Certifique-se de ter o Go instalado em sua máquina (versão 1.21 ou superior)

2. 📥 Clone este repositório:
```bash
git clone [URL_DO_REPOSITÓRIO]
cd [NOME_DO_DIRETÓRIO]
```

3. ▶️ Execute o programa passando um CEP como argumento:
```bash
go run main.go 01153000
```

O CEP deve conter 8 dígitos numéricos.

## 💡 Exemplo de Uso

```bash
# Executando com um CEP válido
go run main.go 01153000

# Saída esperada (exemplo):
Resposta mais rápida da API ViaCEP:
CEP: 01153-000
Logradouro: Rua Vitorino Carmilo
Complemento:
Bairro: Barra Funda
Cidade: São Paulo
UF: SP
```

## 📁 Estrutura do Projeto

- 📄 `main.go`: Contém a implementação principal do programa
- 📦 `go.mod`: Arquivo de gerenciamento de dependências do Go

## ⚙️ Funcionamento

1. 📥 O programa recebe um CEP como argumento da linha de comando
2. 🔄 Inicia duas goroutines simultaneamente para fazer requisições às APIs
3. 📨 Utiliza channels para receber as respostas
4. ⏰ Implementa um timeout de 1 segundo usando context
5. 📝 Exibe a resposta da API mais rápida ou uma mensagem de erro em caso de timeout

## ⚠️ Tratamento de Erros

O programa trata os seguintes casos:
- ❌ CEP não fornecido
- ❌ CEP com formato inválido
- ⏰ Timeout das requisições
- 🔌 Erros de conexão
- 📝 Erros de parsing do JSON