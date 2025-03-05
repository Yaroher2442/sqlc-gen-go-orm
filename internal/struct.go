package golang

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Yaroher2442/sqlc-gen-go-orm/internal/opts"
	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

type Struct struct {
	Table   *plugin.Identifier
	Name    string
	Fields  []Field
	Comment string
}

func (s *Struct) HasPk() bool {
	for _, field := range s.Fields {
		if field.Name == "ID" {
			return true
		}
	}
	return false
}

func (s *Struct) q() string {
	sb := strings.Builder{}
	fields := []string{}
	numbers := []string{}
	for idx, field := range s.Fields {
		fields = append(fields, strcase.ToSnake(field.Name))
		numbers = append(numbers, "$"+strconv.Itoa(idx+1))
	}
	sb.WriteString(fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		s.Table.GetName(),
		strings.Join(fields, ", "),
		strings.Join(numbers, ", "),
	))
	return sb.String()
}

func (s *Struct) upsertArgs() string {
	updateStrs := []string{}
	for _, field := range s.Fields {
		if field.Name == "ID" {
			continue
		}
		updateStrs = append(
			updateStrs,
			fmt.Sprintf(
				"%s = EXCLUDED.%s",
				strcase.ToSnake(field.Name),
				strcase.ToSnake(field.Name),
			),
		)
	}
	return strings.Join(updateStrs, ", ")
}

func (s *Struct) UpsertQuery() string {
	return s.q() + fmt.Sprintf(" ON CONFLICT DO UPDATE SET %s;", s.upsertArgs())
}

func (s *Struct) UpsertReturning() string {
	return s.q() + fmt.Sprintf(" ON CONFLICT DO UPDATE SET %s RETURNING *;", s.upsertArgs())
}

func (s *Struct) InsertQuery() string {
	return s.q() + ";"
}

func (s *Struct) InsertReturning() string {
	return s.q() + " RETURNING *;"
}

func StructName(name string, options *opts.Options) string {
	if rename := options.Rename[name]; rename != "" {
		return rename
	}
	out := ""
	name = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		if unicode.IsDigit(r) {
			return r
		}
		return rune('_')
	}, name)

	for _, p := range strings.Split(name, "_") {
		if _, found := options.InitialismsMap[p]; found {
			out += strings.ToUpper(p)
		} else {
			out += strings.Title(p)
		}
	}

	// If a name has a digit as its first char, prepand an underscore to make it a valid Go name.
	r, _ := utf8.DecodeRuneInString(out)
	if unicode.IsDigit(r) {
		return "_" + out
	} else {
		return out
	}
}
