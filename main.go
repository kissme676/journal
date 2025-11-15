package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Marks  []int
	Medium float64
}

func (s *Student) calc() {
	if len(s.Marks) == 0 {
		s.Medium = 0
		return
	}
	total := 0
	for _, m := range s.Marks {
		total += m
	}
	s.Medium = float64(total) / float64(len(s.Marks))
}

func createStudent(base map[string]Student, r *bufio.Reader) {
	fmt.Print("Введите ФИО: ")
	n, _ := r.ReadString('\n')
	n = strings.TrimSpace(n)

	if _, ok := base[n]; ok {
		fmt.Println("Такой студент уже есть")
		return
	}

	var ms []int
	fmt.Println("Введите оценки через пробел. Пустая строка завершает ввод")
	for {
		fmt.Print("> ")
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			if len(ms) == 0 {
				fmt.Println("Сначала введите хотя бы одну оценку")
				continue
			}
			break
		}

		parts := strings.Fields(line)
		for _, p := range parts {
			v, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("Ошибка ввода:", p)
				continue
			}
			if v < 1 || v > 5 {
				fmt.Println("Оценка вне диапазона:", v)
				continue
			}
			ms = append(ms, v)
		}
	}

	st := Student{Name: n, Marks: ms}
	st.calc()
	base[n] = st

	fmt.Println("Добавлен:", n)
}

func selectByAvg(base map[string]Student, max float64) []Student {
	var out []Student
	for _, st := range base {
		if st.Medium < max {
			out = append(out, st)
		}
	}
	return out
}

func showOne(s Student) {
	fmt.Printf("  Ф/И: %s | Оценки: %v | Средний: %.2f\n", s.Name, s.Marks, s.Medium)
}

func showAll(base map[string]Student) {
	fmt.Println("Все студенты:")
	for _, s := range base {
		showOne(s)
	}
}

func main() {
	db := make(map[string]Student)
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Журнал запущен")

	for {
		fmt.Print("\nКоманда (help - список): ")
		c, _ := r.ReadString('\n')
		c = strings.TrimSpace(c)

		switch c {
		case "add":
			createStudent(db, r)
		case "list":
			if len(db) == 0 {
				fmt.Println("Записей нет")
			} else {
				showAll(db)
			}
		case "filter":
			fmt.Print("Максимальный средний балл: ")
			raw, _ := r.ReadString('\n')
			raw = strings.TrimSpace(raw)
			val, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				fmt.Println("Введите число")
				continue
			}

			res := selectByAvg(db, val)
			if len(res) == 0 {
				fmt.Printf("Нет студентов со средним ниже %.2f.\n", val)
			} else {
				fmt.Printf("Студенты со средним ниже %.2f:\n", val)
				for _, st := range res {
					showOne(st)
				}
			}
		case "help":
			fmt.Println("Команды:")
			fmt.Println("  add - добавить студента")
			fmt.Println("  list - показать всех")
			fmt.Println("  filter - фильтр по среднему")
			fmt.Println("  help - помощь")
			fmt.Println("  exit - выход")
		case "exit":
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
