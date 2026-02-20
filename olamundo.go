package main

import "fmt"

/*

1. [escalabilidade]: Pois o cliente tem que estar seguro de que o site
consiga suportar e funcionar de maneira correta mesmo com sobrecarga
de acessos, o que nao acontece nesse exemplo, onde o cliente tem um
grande prejuizo com o atraso na hora de ser efetuado a compra.

2.[confiabilidade]: Pois no momento de pico de acessos os desenvolvedores
acabaram por nao deixar maneiras de lidar com esse crescimento repentino,
gerando assim uma sobre carga de informaçoes que nao tinham para onde ir,
ocasionando esse vazamento fatal para a reputaçao da empresa do cliente.

3. [manutenibilidade]: Pois devido a falta de organisaçao de uma equipe enterior,
o problema acaba por ser muito mais demorado, trazendo estresse para a nova equipe.
E prejuizo para a empresa que facara com o site inativo ate a resoluçao do problema.


*/

func main() {

	fmt.Print("Ambiente configurado por: Prof.Marcos Andre")
	fmt.Print("partiu carnaval, mas com o midset de alta performance!")

}
