package main

import (
	"fmt"
	"strings"
)

// 1. Olá Mundo
func olaMundo() {
	fmt.Println("Olá, Mundo!")
}

// 2. Soma de dois números
func soma(a, b int) int {
	return a + b
}

// 3. Verificar se o número é par ou ímpar
func parOuImpar(n int) string {
	if n%2 == 0 {
		return "Par"
	}
	return "Ímpar"
}

// 4. Contar o número de caracteres em uma string
func contarCaracteres(s string) int {
	return len(s)
}

// 5. Inverter uma string
func inverterString(s string) string {
	var invertida string
	for i := len(s) - 1; i >= 0; i-- {
		invertida += string(s[i])
	}
	return invertida
}

// 6. Encontrar o maior número em uma lista
func maiorNumero(lista []int) int {
	max := lista[0]
	for _, num := range lista {
		if num > max {
			max = num
		}
	}
	return max
}

// 7. Criar uma função para somar todos os elementos de uma lista
func somarLista(lista []int) int {
	soma := 0
	for _, num := range lista {
		soma += num
	}
	return soma
}

// 8. Verificar se uma palavra é um palíndromo
func isPalindromo(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	invertida := inverterString(s)
	return s == invertida
}

// 9. Contar as vogais em uma string
func contarVogais(s string) int {
	vogais := "aeiouAEIOU"
	count := 0
	for _, c := range s {
		if strings.ContainsRune(vogais, c) {
			count++
		}
	}
	return count
}

// 10. Fatorial de um número
func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	// Desafio 1: Olá Mundo
	olaMundo()

	// Desafio 2: Soma de dois números
	fmt.Println("Soma de 3 e 5:", soma(3, 5))

	// Desafio 3: Par ou ímpar
	fmt.Println("Número 7 é:", parOuImpar(7))

	// Desafio 4: Contar caracteres em uma string
	fmt.Println("Número de caracteres em 'GoLang':", contarCaracteres("GoLang"))

	// Desafio 5: Inverter uma string
	fmt.Println("Inverter 'golang':", inverterString("golang"))

	// Desafio 6: Maior número em uma lista
	fmt.Println("Maior número na lista [1, 2, 10, 4]:", maiorNumero([]int{1, 2, 10, 4}))

	// Desafio 7: Somar todos os elementos de uma lista
	fmt.Printl
