package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var files []string
var packageName string
var packets []string
var structs = map[string]*ast.TypeSpec{}

const (
	packetSearchString = "This is a Minecraft packet"
	searchString       = "This is a packet"
)

func main() {
	fi, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fi {
		if _, err := strconv.Atoi(strings.TrimSuffix(f.Name(), ".go")); err != nil {
			files = append(files, f.Name())
		}
	}
	for _, file := range files {
		fs := token.NewFileSet()
		parsedFile, err := parser.ParseFile(fs, file, nil, parser.ParseComments)
		if err != nil {
			log.Fatalln("COULD NOT PARSE", file, "BECAUSE", err)
		}
		packageName = parsedFile.Name.Name
		for _, decl := range parsedFile.Decls {
			switch decl := decl.(type) {
			case *ast.GenDecl:
				if decl.Tok != token.TYPE {
					continue
				}

				if len(decl.Specs) != 1 {
					return
				}

				tSpec, ok := decl.Specs[0].(*ast.TypeSpec)
				if !ok {
					continue
				}
				_, ok = tSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				structs[tSpec.Name.Name] = tSpec

				if decl.Doc == nil {
					continue
				}
				doc := decl.Doc.Text()
				pos := strings.Index(doc, packetSearchString)
				if pos == -1 {
					pos = strings.Index(doc, searchString)
					if pos == -1 {
						continue
					}
				}

				packets = append(packets, tSpec.Name.Name)
			}
		}
	}
	/*
		fmt.Println("DONE")
		fmt.Println("Packets:")
		for _, p := range packets {
			fmt.Println("Packet", p)
		}

		fmt.Println("Structs")
		for s, ts := range structs {
			fmt.Println(s)
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				fmt.Println("Is not a struct")
			} else {
				for _, f := range st.Fields.List {
					fmt.Println("  Field", f.Names, makeExprReadable(f.Type))
				}
			}
		}
	*/
	//fmt.Println("\nTranslate switch:")
	//fmt.Printf(makeSwitch())

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintln("package protocol\n"))

	buf.WriteString(fmt.Sprintf(`import (
	"github.com/ShadowJonathan/mopher/lib"
	"reflect"
	"fmt"
	"github.com/ShadowJonathan/mopher/versions/%s"
)
	`, strings.TrimPrefix(packageName, "_")))

	buf.WriteString(makeSwitch(false))
	buf.WriteString("\n\n\n")
	buf.WriteString(makeSwitch(true))

	b, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("ERR MAKING FORMATTED:", err)
		b = buf.Bytes()
	}

	o, err := os.Create("../../translate" + packageName + ".go")
	if err != nil {
		log.Fatalln(err)
	}
	defer o.Close()
	o.Write(b)
}

func makeSwitch(back bool) (t string) {
	if back {
		t += `func Back` + packageName + `(i interface{}) (lib.MetaPacket, error) {
	switch i := i.(type) {
	`

		for _, p := range packets {
			if protocolHasPacket(p) {
				rs, errs, exc := makeRepresented(p, back)
				if exc {
					t += extensive(p, back)
				} else {
					if errs != "" {
						errs = "/*" + errs + "*/"
					}

					t += fmt.Sprintf(`case *%s.%s:
			return &%s{%s}, nil %s
			`, packageName, p, p, rs, errs)
				}
			} else {
				t += fmt.Sprintln("// FIXME add", p)
			}
		}

		t += `
	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
	}`
	} else {
		t += `func Translate` + packageName + `(i interface{}) (lib.Packet, error) {
	if p, ok := i.(lib.Packet); ok {
		return p, nil
	}
	switch i := i.(type) {
	`
		for _, p := range packets {
			if protocolHasPacket(p) {
				rs, errs, exc := makeRepresented(p, back)
				if exc {
					t += extensive(p, back)
				} else {
					if errs != "" {
						errs = "/*" + errs + "*/"
					}

					t += fmt.Sprintf(`case *%s:
			return &%s.%s{%s}, nil %s
			`, p, packageName, p, rs, errs)
				}
			} else {
				t += fmt.Sprintln("// FIXME add", p)
			}
		}

		t += `
	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
	}`
	}
	return
}

func makeExprReadable(e ast.Expr) (t string) {
	switch e := e.(type) {

	case *ast.StarExpr:
		t = "*" + makeExprReadable(e.X)
	case *ast.SelectorExpr:
		t = makeExprReadable(e.X) + "." + fmt.Sprint(e.Sel)
	case *ast.Ident:
		t = fmt.Sprint(e)
	case *ast.ArrayType:
		var l string
		if e.Len != nil {
			l = makeExprReadable(e.Len)
		}
		t = "[" + l + "]" + makeExprReadable(e.Elt)

	default:
		fmt.Println(">>> Cannot switch to type", reflect.TypeOf(e))
		t = fmt.Sprint(e)
	}

	return t
}

func makeRepresented(p string, back bool) (t string, errs string, extensive bool) {
	pts := protocolStructs[p].Type.(*ast.StructType)
	vts := structs[p].Type.(*ast.StructType)

	fields := make(map[string]string)
	for _, F := range pts.Fields.List {
		for _, N := range F.Names {
			for _, f := range vts.Fields.List {
				for _, n := range f.Names {
					if N.Name == n.Name {
						fields[n.Name] = ""

						if makeExprReadable(f.Type) != makeExprReadable(F.Type) {
							fmt.Printf("%s IS NOT SAME TYPE: %s != %s\n", n.Name, makeExprReadable(f.Type), makeExprReadable(F.Type))
							if back {
								fields[n.Name] = makeExprReadable(F.Type)
							} else {
								fields[n.Name] = makeExprReadable(f.Type)
							}
						} else {
							fields[n.Name] = ""
						}
					}

					//fmt.Println(n.Name, reflect.TypeOf(F.Type))

					if T, ok := f.Type.(*ast.ArrayType); ok {
						if unicode.IsUpper(rune(makeExprReadable(T.Elt)[0])) {
							extensive = true
							return
						}
					}
				}
			}
		}
	}

	tStatements := []string{}
	for n, wrap := range fields {
		if wrap == "" {
			tStatements = append(tStatements, fmt.Sprintf("%s: i.%s", n, n))
		} else {
			tStatements = append(tStatements, fmt.Sprintf("%s: %s(i.%s)", n, wrap, n))
		}
	}

	t = strings.Join(tStatements, ", ")
	return
}
