package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaResoluciones_20210120_153032 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaResoluciones_20210120_153032{}
	m.Created = "20210120_153032"

	migration.Register("CrearEsquemaResoluciones_20210120_153032", m)
}

// Run the migrations
func (m *CrearEsquemaResoluciones_20210120_153032) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210120_153032_crear_esquema_resoluciones_up.sql")

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
func (m *CrearEsquemaResoluciones_20210120_153032) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210120_153032_crear_esquema_resoluciones_down.sql")

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
