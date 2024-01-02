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
		
		func Genrate_payload(shellcode string,number int){
		var mop []string
		for i := 0; i < number; i++ {
			shellcode_parent1 := shellcode
			shellcode_parent2 := shellcode
			Child := Crossover(shellcode_parent1, shellcode_parent2)
			MutatedChild := Mutation(Child)
			mop = append(mop, MutatedChild)
		}
		for iter, shell_code := range mop {
			fmt.Printf("indexado: %d, valor: %s\n", iter, shell_code)
		}
		
		
		}
		func main(){
		
		
		fmt.Println("######::'##::::'##:'########:'##:::::::'##:::::: ::'######::'##::::'##::::'###:::'##::::'##:'##### ###:'##:::::::'########::'#######::'##::: ##:")
		fmt.Println("'##... ##: ##:::: ##: ##.....:: ##::::::: ##:::::::'##.. . ##: ##:::: ##:::'## ##::: ###::'###: ##.....:: ##:::::: : ##.....::'##.... ##: ###:: ##:")
		fmt.Println("##:::..:: ##:::: ##: ##::::::: ##::::::: ##::::::: ##::: ..:: ##:::: ##::'##:. ##:: ####'####: ##::::::: ##::::::: ##::::::: ##:::: ##: ####: ##:")
		fmt.Println(". ######:: #########: ######::: ##::::::: ##::::::: ##::: :::: #########:'##:::. ##: ## ### ##: ######::: ##::::::: ######::: ##:::: ##: ## # ###:")
		fmt.Println(":..... ##: ##.... ##: ##...:::: ##::::::: ##::::::: ##::: :::: ##.... ##: #########: ##. #: ##: ##...:::: ##::::::: ##...:::: ##:::: ##: ##. ####:")
		fmt.Println("'##::: ##: ##:::: ##: ##::::::: ##::::::: ##::::::: ##::: ##: ##:::: ##: ##.... ##: ##:.:: ##: ##::::::: ##::::::: ## ::::::: ##:::: ##: ##:. ###:")
		fmt.Println(". ######:: ##:::: ##: ########: ########: ########:. ######:: ##:::: ##: ##:::: ##: ##:::: ##: ########: ###### ##: ########:. #######:: ##::. ##:")
		fmt.Println(":......:::..:::::..::........::........::........: ::......:::..:::::..::..:::::..::..:::::..::...... ..::........::........:::.....:::..::::..::")
		var shellcode string
    var variantes int

    fmt.Println("Digite o shellcode:")
    fmt.Scanln(&shellcode)

    fmt.Println("Digite o  numero de variantes:")
    fmt.Scanln(&variantes)
	Genrate_payload(shellcode,variantes)

    
		
		
		
			
		
		}
