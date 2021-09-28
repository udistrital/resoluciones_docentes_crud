package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CuadroResponsabilidad_20210924_145046 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CuadroResponsabilidad_20210924_145046{}
	m.Created = "20210924_145046"

	migration.Register("20210924_145046_CuadroResponsabilidad", m)
}

// Run the migrations
func (m *CuadroResponsabilidad_20210924_145046) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210924_145046_cuadro_responsabilidad_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CuadroResponsabilidad_20210924_145046) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210924_145046_cuadro_responsabilidad_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
