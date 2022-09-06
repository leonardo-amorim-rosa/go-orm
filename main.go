package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome, Area, Professor string
}

func main() {
	// conectando com o banco de dados sqlite
	db, err := gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao acessar banco de dados")
	} else {
		fmt.Println("Banco de dados conectado com sucesso!")
	}

	// migrando a tabela e Cursos
	db.AutoMigrate(&Curso{})

	// criar um curso novo
	db.Create(&Curso{Nome: "Programar em Go", Area: "Tecnologia", Professor: "Leonardo"})

	// buscando todos os registros
	var cursos []Curso
	db.Find(&cursos)
	fmt.Println(cursos)

	// buscar um curso
	var curso Curso
	db.First(&curso, 1)
	fmt.Println(curso.Nome)

	// atualizando o curso
	db.Model(&curso).Update("Nome", "Curso atualizado de go")

	// removendo um curso
	db.First(&curso, 1)
	db.Delete(&curso, 1)
}
