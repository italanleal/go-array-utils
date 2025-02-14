# go-array-utils

Uma coleção de funções utilitárias para manipulação de slices em Go.

## Instalação

```sh
go get github.com/seu-usuario/utils
```

## Uso

Importe o pacote no seu código:

```go
package main

import (
	"fmt"
	"github.com/seu-usuario/utils"
)

func main() {
	// Exemplo de uso da função Map
	numbers := []int{1, 2, 3, 4, 5}
	squared := utils.Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squared numbers:", squared)

	// Exemplo de uso da função Reduce
	sum := utils.Reduce(numbers, func(acc, cur int) int {
		return acc + cur
	}, 0)
	fmt.Println("Sum of numbers:", sum)

	// Exemplo de uso da função ForEach
	fmt.Println("Printing numbers:")
	utils.ForEach(numbers, func(n int) {
		fmt.Println(n)
	})

	// Exemplo de uso da função Reverse
	reversed := utils.Reverse(numbers)
	fmt.Println("Reversed numbers:", reversed)

	// Exemplo de uso da função Filter
	evenNumbers := utils.Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers)
}
```

## Funções Disponíveis

### Map

Aplica uma função a cada elemento de um slice e retorna um novo slice com os resultados.

```go
func Map[T any, R any](slice []T, mapper func(T) R) []R
```

### Reduce

Acumula os valores de um slice usando uma função combinadora.

```go
func Reduce[T any](slice []T, reducer func(acc, cur T) T, initial T) T
```

### ForEach

Executa uma função para cada elemento de um slice.

```go
func ForEach[T any](slice []T, action func(T))
```

### Reverse

Inverte a ordem dos elementos de um slice.

```go
func Reverse[T any](slice []T) []T
```

### Filter

Retorna um novo slice contendo apenas os elementos que satisfazem a condição fornecida.

```go
func Filter[T any](slice []T, predicate func(T) bool) []T
```

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

