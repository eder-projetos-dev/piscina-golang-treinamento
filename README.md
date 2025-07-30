# 💧 Projeto Piscina

Este é um exercício introdutório em **Go (Golang)** para treinar estruturas de repetição, funções e testes unitários.  
O sistema simula o cálculo de **tratamentos químicos** para piscinas com base no volume de água informado e na opção selecionada pelo usuário.

---

## 📦 Módulo

Inicialização do módulo Go:

```bash
go mod init piscina
```

---

## 📁 Estrutura do Projeto

```
piscina/
├── go.mod              # Gerenciador de dependências do Go
├── piscina.go          # Lógica principal com cálculos de tratamento
└── piscina_test.go     # Casos de teste automatizados
```

---

## ▶️ Como Executar

### Executar o programa (modo interativo):

```bash
go run piscina.go
```

### Rodar os testes:

```bash
go test

# ou com detalhes
go test -v
```

---

## ✅ Funcionalidades Implementadas

- [x] Cálculo de dosagens diárias e semanais
- [x] Decantação química
- [x] Correção de pH
- [x] Superdosagem para tratamento intensivo
- [x] Validação de opções inválidas
- [x] Testes unitários com `testing.T`

---

## 🧪 Exemplo de Cenário de Teste

**Cenário**: Cálculo de dosagem semanal para piscina média  
**Dado**: `M = 50 m³` e opção 1 selecionada  
**Então**:
- 700 ml de algistático inicial
- 300 ml semanal de algistático
- 1250 ml de cloro líquido

---

## 📚 Requisitos

- Go 1.21 ou superior
- Terminal ou IDE com suporte a execução Go

---

## 🛠️ Em desenvolvimento

Este repositório será expandido com novas opções de tratamento, testes e melhorias estruturais ao longo do curso.

---

## 📄 Licença

Este projeto está sob a licença MIT.
