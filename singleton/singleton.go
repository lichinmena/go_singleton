package singleton

import (
	"sync"
)

var (
	p    *person
	once sync.Once //Permite bloqeuar de manera segura que se ejecute una unica vez un proceso
)

//Puntero

func GetInstance() *person {
	once.Do(func() {
		p = &person{}
	}) //Se asegura que se ejecuta una unica vez
	return p
}

/*
NO permite ejecutar el segundo, El primero que  sea llamado,
No permite hablar al segundo,por que se hace de manera atomica
para todo el paquete
*/
/*
func ShowPerson() {
	once.Do(func() {
		fmt.Println("Me estan llamando nuevamente!!!")
	})
}
*/

//Se crea de manera privada
type person struct {
	name string
	age  int
	mux  sync.Mutex //Permite Bloquear Momentaneamente para asegurarme ejecutar correctamentelo que necesito y luego desbloquea
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock() //Asegura al final de la ejecucion libera el recurso
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock() //Asegura al final de la ejecucion libera el recurso
	return p.age
}
