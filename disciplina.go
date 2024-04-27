package main
import (
	"fmt"
)
func main() {
	//para armazenar os valores da oito provas

	//para armazenar o valor das notas das listas


	//preciso armazenar a matriculo do aluno: _, err:= fmt.Scan(&matricula)
	//eu coloco o _ e virgula para ignorar o erro, err:= fmt.Scan(&matricula) é a mesma coisa que eu colocar var matricula int e depois matricula, err:= fmt.Scan(&matricula)
	//preciso verificar se a matricula é 0 ou -1, se for, encerrar o programa

	// preciso ler as notas das provas e armazenar em uma variavel provas
	// para ler os valores das provas eu preciso de um loop for que ira percorrer 8 vezes
	for i := 0; i < 8; i++ {
		var provas float64
		fmt.Scanf("%f", &provas)
		// eu ireia ler as notas das provas e armazenar em variavel notaProva que ira somar todas as notas das provas
		// preciso ler as notas 
		var notaProva float64
		notaProva += provas
	}
	for i := 0; i < 5; i++ {
		var lista float64
		// eu irei ler as notas das listas e armazenar em variavel notaLista que ira somar todas as notas das listas
		fmt.Scanf("%f", &lista)
		var notaLista float64
		notaLista += lista
	}

	// preciso ler a nota do trabalho final e a presenca
	var trabalhoFinal float64
	fmt.Scanf("%f", &trabalhoFinal)
	var presenca float64
	fmt.Scanf("%f", &presenca)

	// preciso calcular a media das notaas das provas que sao as notass da provas / quantidade de provas

	// preciso calcular a medias das notas das listas que as notas da listas / quantidade de lista

	// para calcular a nota final eu pego a media da prova e multiplico por 0,7 e somo a media da lista * 0.15 + nota do trabalho final * 0.15

	//eu preciso verificar a situalcao final do aluno que se 
	// if a notafinal >= 6 && presenca >= 75 entao sera aprovado
	// se a nota final >= 6 e presenca < 75 entao sera reprovado por falta
	// se a nota final < 6 e presenca >= 75 entao sera reprovado por nota
	//senao sera reprovado por falta e por nota
}