package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

var Interfaces map[string]reflect.Type



func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() 
	defer func() {
		
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

type lexer struct {
	scan  scanner.Scanner
	token rune 
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { 
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}






















func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		
		switch lex.text() {
		case "nil":
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		case "t":
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) 
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) 
		switch v.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v.SetInt(int64(i))
			lex.next()
			return
		case reflect.Float32, reflect.Float64:
			v.SetFloat(float64(i))
			lex.next()
			return
		}
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64) 
		v.SetFloat(float64(f))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() 
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}


func fieldByName(structure reflect.Value, name string) reflect.Value {
	field := structure.FieldByName(name)
	if field != (reflect.Value{}) {
		return field
	}
	typ := structure.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Tag.Get("sexpr") == name {
			return structure.FieldByName(field.Name)
		}
	}
	return reflect.Value{}
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: 
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: 
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: 
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, fieldByName(v, name))
			lex.consume(')')
		}

	case reflect.Map: 
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	case reflect.Interface: 
		name := strings.Trim(lex.text(), `"`)
		lex.next()
		typ, ok := Interfaces[name]
		if !ok {
			panic(fmt.Sprintf("no concrete type registered for interface %s", name))
		}
		val := reflect.New(typ)
		read(lex, reflect.Indirect(val))
		v.Set(reflect.Indirect(val))

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

func init() {
	Interfaces = make(map[string]reflect.Type)
}
