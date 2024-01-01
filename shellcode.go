	//Gerador de Shellcoder Polimorfico
	package main
	import (
		"fmt"
		"math/rand"
		"time"
	)
	var shell []string
	func Crossover(shellcode_parent1 string,shellcode_parent2 string)(string){
	rand.Seed(time.Now().UnixNano())
	fuction_point_croosver:=rand.Intn(len(shellcode_parent1))
	Child:=shellcode_parent1[:fuction_point_croosver]+shellcode_parent2[:fuction_point_croosver]
	return Child
	
	
	
	
	}
	
	func Mutation( xild string)(string){
	rand.Seed(time.Now().UnixNano())
	Ponto_de_Mutaçao:= rand.Intn(len(xild))
	child_mutaton:= xild[:Ponto_de_Mutaçao] + string(byte((int(xild[Ponto_de_Mutaçao])+1)%256)) + xild[Ponto_de_Mutaçao+1:]
	return child_mutaton
	
	}
	
	
	func main(){
	var    shellcode string
	//var  number int
	
	fmt.Println("######::'##::::'##:'########:'##:::::::'##:::::: ::'######::'##::::'##::::'###:::'##::::'##:'##### ###:'##:::::::'########::'#######::'##::: ##:")
	fmt.Println("'##... ##: ##:::: ##: ##.....:: ##::::::: ##:::::::'##.. . ##: ##:::: ##:::'## ##::: ###::'###: ##.....:: ##:::::: : ##.....::'##.... ##: ###:: ##:")
	fmt.Println("##:::..:: ##:::: ##: ##::::::: ##::::::: ##::::::: ##::: ..:: ##:::: ##::'##:. ##:: ####'####: ##::::::: ##::::::: ##::::::: ##:::: ##: ####: ##:")
	fmt.Println(". ######:: #########: ######::: ##::::::: ##::::::: ##::: :::: #########:'##:::. ##: ## ### ##: ######::: ##::::::: ######::: ##:::: ##: ## # ###:")
	fmt.Println(":..... ##: ##.... ##: ##...:::: ##::::::: ##::::::: ##::: :::: ##.... ##: #########: ##. #: ##: ##...:::: ##::::::: ##...:::: ##:::: ##: ##. ####:")
	fmt.Println("'##::: ##: ##:::: ##: ##::::::: ##::::::: ##::::::: ##::: ##: ##:::: ##: ##.... ##: ##:.:: ##: ##::::::: ##::::::: ## ::::::: ##:::: ##: ##:. ###:")
	fmt.Println(". ######:: ##:::: ##: ########: ########: ########:. ######:: ##:::: ##: ##:::: ##: ##:::: ##: ########: ###### ##: ########:. #######:: ##::. ##:")
	fmt.Println(":......:::..:::::..::........::........::........: ::......:::..:::::..::..:::::..::..:::::..::...... ..::........::........:::.....:::..::::..::")
	
	fmt.Println("Digite  o Shellcode :")
	fmt.Scanf("%s",&shellcode)
	//fmt.Println("Digte o Numero de Variantes ")
	//fmt.Scanf("%d",&number)
		var variants []string
	for i := 0; i < 33; i++ {
		parent1 := shellcode
		parent2 := shellcode
		child := Crossover(parent1, parent2)
		mutatedChild := Mutation(child)
		variants = append(variants, mutatedChild)
	}
	for index, value := range variants {
        fmt.Printf("Indexado: %d, Valor: %s\n", index, value)
    }
	}
