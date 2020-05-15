package main

import "fmt"

type livingBeing interface {
	itsAlive() bool
}

// Interfaces : Permiten relacionar objetos del mismo tipo
type human interface {
	breathe()
	think()
	eat()
	sex() string
	itsAlive() bool
}

type animal interface {
	breathe()
	eat()
	isCarnivorous() bool
	itsAlive() bool
}

type vegetal interface {
	vegetableClassification() string
}

// Type human
type man struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	isMan     bool
	isAlive   bool
}

// Type human
type woman struct {
	man
}

func (m *man) breathe() { m.breathing = true }
func (m *man) think()   { m.thinking = true }
func (m *man) eat()     { m.eating = true }
func (m *man) sex() string {
	if m.isMan {
		return "Man"
	}
	return "Woman"
}
func (m *man) itsAlive() bool { return m.isAlive }

func humansBreathing(hu human) {
	hu.breathe()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.sex())
}

// Type animal
type dog struct {
	breathing   bool
	eating      bool
	carnivorous bool
	isAlive     bool
}

func (d *dog) breathe()            { d.breathing = true }
func (d *dog) eat()                { d.eating = true }
func (d *dog) isCarnivorous() bool { return d.carnivorous }
func (d *dog) itsAlive() bool      { return d.isAlive }

// Una función para procesar cualquier tipo de animal
func animalsBreathing(an animal) {
	an.breathe()
	fmt.Println("Soy un animal y estoy respirando")
}

func carnivorousAnimals(an animal) int {
	if an.isCarnivorous() == true {
		return 1
	}
	return 0
}

func amAlive(lb livingBeing) bool {
	return lb.itsAlive()
}

func main() {
	man1 := new(man)
	man1.isMan = true
	humansBreathing(man1)

	woman1 := new(woman)
	humansBreathing(woman1)

	totalCarnivores := 0
	dog1 := new(dog)
	dog1.carnivorous = true
	dog1.isAlive = true
	animalsBreathing(dog1)
	totalCarnivores += carnivorousAnimals(dog1)
	fmt.Printf("Total carnivoros %d \n", totalCarnivores)
	fmt.Printf("Estoy vivo = %t", amAlive(dog1))
}
