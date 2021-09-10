package clientone

import "example/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
