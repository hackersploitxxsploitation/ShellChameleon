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
	for j := 0; j < 33; i++ {
		Shellcode1 := shellcode
		Shellcode2 := shellcode
		Child := Crossover(Shellcode1, Shellcode2)
		MutatedChild := Mutation(child)
		shell_metamorfic = append(shell_metamorfic, MutatedChild)
	}
	for index, value := range variants {
        fmt.Printf("Indexado: %d, Valor: %s\n", index, value)
    }
	}
