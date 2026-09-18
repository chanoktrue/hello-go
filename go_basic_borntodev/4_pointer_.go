package main

type Employee struct {
	name   string
	salary float64
}

func (e *Employee) getSalary(amount float64) {
	e.salary += amount
}

func main() {
	emp := Employee{name: "John", salary: 50000.0}
	emp.getSalary(1000.0)
	println(emp.name, emp.salary)
}
