package pointers

import "log"

type Person struct {
	name    string
	surname string
	age     int
}

// * значение по адресу
// & адрес в памяти
// *& значение по адресу
// если передать адрес(&) и после менять через * то изменится передаваемый аргумент без создания нового или копии
// * для работы с исходной переменной, & передавать ее, без указателей работать через создание копии и возврата нового объекта

// SetDataWithPointer Изменение исходной структуры с указателем
func (p *Person) setDataWithPointer(name string, surname string, age int) {
	p.name = name
	p.surname = surname
	p.age = age
}

// SetDataWithoutPointer Изменение исходной структуры без указателя
func (p Person) setDataWithoutPointer(name string, surname string, age int) {
	p.name = name
	p.surname = surname
	p.age = age
}

// setData Изменение исходной структуры с указателем но не через функцию структуры
func setData(human *Person, name string, surname string, age int) {
	human.name = name
	human.surname = surname
	human.age = age
}

func SetDataMain() {
	var human Person
	human.age = 201212
	setData(&human, "Murasha", "MRSHK", 21)
	log.Println("Структура после установки информации с указателем &: ", human)
	human.setDataWithPointer("Murasha Updated", "MURASHIN'KA", 21)
	log.Println("Структура после установки информации с указателем через функцию структуры: ", human)
	human.setDataWithoutPointer("Murasha BEZ POINTERA", "IDK", 8)
	log.Println("Структура после установки информации без указателя: ", human)
}
