package utils

/*
Map aplica uma função fn a cada elemento do slice arr e retorna um novo slice contendo os resultados.

Parâmetros:

- arr: Slice de elementos do tipo T.

- fn: Função que recebe um elemento de tipo T e retorna um valor de tipo R.

Retorno:

- Um slice contendo os valores resultantes da aplicação de fn a cada elemento de arr.

Exemplo de uso:

	numeros := []int{1, 2, 3, 4}
	dobrados := Map(numeros, func(n int) int {
		return n * 2
	})
	// dobrados agora contém [2, 4, 6, 8]
*/
func Map[T any, R any](arr []T, fn func(T) R) []R {
	var result []R
	for _, value := range arr {
		result = append(result, fn(value))
	}
	return result
}

/*
Reduce aplica uma função de redução fn a cada elemento do slice arr, acumulando os resultados em um valor único.

Parâmetros:

- arr: Slice de elementos do tipo T.

- fn: Função que recebe um acumulador (acc) e o elemento atual (cur), retornando o novo valor acumulado.

- start: Valor inicial do acumulador.

Retorno:

- O valor final acumulado após a aplicação da função fn a todos os elementos de arr.

Exemplo de uso:

	numeros := []int{1, 2, 3, 4}
	soma := Reduce(numeros, func(acc, cur int) int {
		return acc + cur
	}, 0)
	// soma agora contém 10 (1 + 2 + 3 + 4)
*/
func Reduce[T any](arr []T, fn func(acc T, cur T) T, start T) T {
	var result T = start
	for _, value := range arr {
		result = fn(result, value)
	}
	return result
}

/*
ForEach percorre todos os elementos do slice arr e aplica a função fn a cada um deles.

Parâmetros:

- arr: Slice de elementos do tipo T.

- fn: Função que recebe um elemento de tipo T e executa uma operação sobre ele, sem retorno.

Retorno:

- Nenhum retorno (void).

Exemplo de uso:

	numeros := []int{1, 2, 3, 4}
	ForEach(numeros, func(n int) {
		fmt.Println(n)
	})
	// Saída:
	// 1
	// 2
	// 3
	// 4
*/
func ForEach[T any](arr []T, fn func(T)) {
	for _, value := range arr {
		fn(value)
	}
}

/*
Reverse inverte a ordem dos elementos em um slice e retorna um novo slice com os elementos na ordem reversa.

Parâmetros:

- arr: Slice de elementos do tipo T a ser invertido.

Retorno:

- Um novo slice contendo os elementos de arr na ordem inversa.

Exemplo de uso:

	numeros := []int{1, 2, 3, 4, 5}
	invertido := Reverse(numeros)
	// invertido agora contém [5, 4, 3, 2, 1]
*/
func Reverse[T any](arr []T) []T {
	var result []T

	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}

	return result
}

/*
Filter filtra os elementos de um slice com base em uma função condicional e retorna um novo slice contendo apenas os elementos que atendem ao critério especificado.

Parâmetros:

- arr: Slice de elementos do tipo T a ser filtrado.

- fn: Função que recebe um elemento de T e retorna um booleano indicando se o elemento deve ser incluído no resultado.

Retorno:

- Um novo slice contendo apenas os elementos de arr que satisfazem a condição definida em fn.

Exemplo de uso:

	numeros := []int{1, 2, 3, 4, 5}
	pares := Filter(numeros, func(n int) bool {
		return n%2 == 0
	})
	// pares agora contém [2, 4]
*/
func Filter[T any](arr []T, fn func(T) bool) []T {
	var result []T

	for _, value := range arr {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}
