package main

import (
	"example/clientone"
	"example/clienttwo"
	"example/singleton"
	"fmt"
	"sync"
)

func main() {

	//Agregar concurrencia
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() { //Funciones anonimas
			defer wg.Done()
			clientone.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			clienttwo.IncrementAge()
		}()
	}

	wg.Wait()

	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
	//singleton.ShowPerson() No se ejecuta por el tema de atomicidad
}
