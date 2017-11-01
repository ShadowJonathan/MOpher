package main

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
	"unicode"
)

const begin = "case *%s:\n"
const end = "return &%s.%s{%s}, nil\n"

const beginBack = "case *%s.%s:\n"
const endBack = "return &%s{%s}, nil\n"

var tmpAt int

func extensive(p string, back bool) string {
	pts := protocolStructs[p].Type.(*ast.StructType)
	vts := structs[p].Type.(*ast.StructType)

	var before string
	var middle string
	var insert []string

	//	fields := make(map[string]bool)
	for _, F := range pts.Fields.List {
		for _, N := range F.Names {
			for _, f := range vts.Fields.List {
				for _, n := range f.Names {
					if n.Name == N.Name {
						if t, ok := f.Type.(*ast.ArrayType); ok {
							//fmt.Println(" ", n.Name, t.Elt, protocolStructs[makeExprReadable(t.Elt)])
							if _, ok := protocolStructs[makeExprReadable(t.Elt)]; ok {
								At := tmpAt
								tmpAt++
								if back {
									before += fmt.Sprintf("var tmp%d []%s\n", At, t.Elt)
								} else {
									before += fmt.Sprintf("var tmp%d []%s.%s\n", At, packageName, t.Elt)
								}
								insert = append(insert, fmt.Sprintf(`%s: tmp%d`, n.Name, At))
								lv, lb, li := makeLinker(
									structs[makeExprReadable(t.Elt)].Type.(*ast.StructType),
									protocolStructs[makeExprReadable(t.Elt)].Type.(*ast.StructType),
									back,
								)
								if back {
									middle += fmt.Sprintf(`for _, v := range i.%s {
								%s
								%s
								tmp%d = append(tmp%d, %s{%s})
								}
								`, n.Name, lv, lb, At, At, t.Elt, li)
								} else {
									middle += fmt.Sprintf(`for _, v := range i.%s {
								%s
								%s
								tmp%d = append(tmp%d, %s.%s{%s})
								}
								`, n.Name, lv, lb, At, At, packageName, t.Elt, li)
								}
							} else {
								insert = append(insert, fmt.Sprintf(`%s: i.%s`, n.Name, n.Name))
							}
						} else {
							if makeExprReadable(f.Type) != makeExprReadable(F.Type) {
								fmt.Printf("%s IS NOT SAME TYPE: %s != %s\n", n.Name, makeExprReadable(f.Type), makeExprReadable(F.Type))
							}
							insert = append(insert, fmt.Sprintf(`%s: i.%s`, n.Name, n.Name))
						}
					}
				}
			}
		}
	}
	if back {
		return fmt.Sprintf(beginBack, packageName, p) + before + middle + fmt.Sprintf(endBack, p, strings.Join(insert, ", "))
	} else {
		return fmt.Sprintf(begin, p) + before + middle + fmt.Sprintf(end, packageName, p, strings.Join(insert, ", "))
	}
}

func makeLinker(p, v *ast.StructType, back bool) (lv string, lb string, li string) {

	var values string
	var before string
	var insert []string

	fields := make(map[string]bool)
	for _, F := range p.Fields.List {
		for _, N := range F.Names {
			for _, f := range v.Fields.List {
				for _, n := range f.Names {
					if N.Name == n.Name && reflect.TypeOf(f.Type) == reflect.TypeOf(F.Type) {
						fields[n.Name] = true

						if T, ok := f.Type.(*ast.ArrayType); ok {
							if unicode.IsUpper(rune(makeExprReadable(T.Elt)[0])) {
								At := tmpAt
								tmpAt++

								if back {
									before += fmt.Sprintf("var tmp%d []%s\n", At, T.Elt)
								} else {
									before += fmt.Sprintf("var tmp%d []%s.%s\n", At, packageName, T.Elt)
								}
								insert = append(insert, fmt.Sprintf(`%s: tmp%d`, n.Name, At))

								//fmt.Println(makeExprReadable(T.Elt), structs[makeExprReadable(T.Elt)].Type.(*ast.StructType), At)
								lv, lb, li := makeLinker(
									structs[makeExprReadable(T.Elt)].Type.(*ast.StructType),
									protocolStructs[makeExprReadable(T.Elt)].Type.(*ast.StructType),
									back,
								)
								if back {
									before += fmt.Sprintf(`for _, v := range v.%s {
										%s
										%s
										tmp%d = append(tmp%d, %s{%s})
									}
								`, n.Name, lv, lb, At, At, T.Elt, li)
								} else {
									before += fmt.Sprintf(`for _, v := range v.%s {
										%s
										%s
										tmp%d = append(tmp%d, %s.%s{%s})
									}
								`, n.Name, lv, lb, At, At, packageName, T.Elt, li)
								}
							} else {
								insert = append(insert, fmt.Sprintf(`%s: v.%s`, n.Name, n.Name))
							}
						} else {
							if unicode.IsUpper(rune(makeExprReadable(f.Type)[0])) {
								lv, lb, li := makePacketLink(
									structs[makeExprReadable(f.Type)].Type.(*ast.StructType),
									protocolStructs[makeExprReadable(f.Type)].Type.(*ast.StructType),
									makeExprReadable(f.Type),
									"v",
									n.Name,
									back,
								)
								values += lv
								before += lb
								if back {
									insert = append(insert, fmt.Sprintf(`%s: %s`, n.Name, li))
								} else {
									insert = append(insert, fmt.Sprintf(`%s: %s.%s`, n.Name, packageName, li))
								}
							} else {
								if makeExprReadable(f.Type) != makeExprReadable(F.Type) {
									fmt.Printf("%s IS NOT SAME TYPE: %s != %s\n", n.Name, makeExprReadable(f.Type), makeExprReadable(F.Type))
								}
								insert = append(insert, fmt.Sprintf(`%s: v.%s`, n.Name, n.Name))
							}
						}
					}
				}
			}
		}
	}

	lv = values
	lb = before
	li = strings.Join(insert, ", ")
	//fmt.Println(lv, "\n"+strings.TrimSpace(lb)+"\n", li)
	return
}

func makePacketLink(p, v *ast.StructType, t, prefix, name string, back bool) (V, B, I string) {

	var pre = name
	if prefix != "" {
		pre = prefix + "." + pre
	}

	I += t + "{"

	for _, F := range p.Fields.List {
		for _, N := range F.Names {
			for _, f := range v.Fields.List {
				for _, n := range f.Names {
					if n.Name == N.Name {
						if T, ok := f.Type.(*ast.ArrayType); ok {
							if unicode.IsUpper(rune(makeExprReadable(T.Elt)[0])) {
								At := tmpAt
								tmpAt++
								if back {
									V += fmt.Sprintf("var tmp%d []%s\n", At, T.Elt)
								} else {
									V += fmt.Sprintf("var tmp%d []%s.%s\n", At, packageName, T.Elt)
								}
								I += n.Name + fmt.Sprintf(": tmp%d", At) + ", "

								//fmt.Println(makeExprReadable(T.Elt), structs[makeExprReadable(T.Elt)].Type.(*ast.StructType), At)
								lv, lb, li := makeLinker(
									structs[makeExprReadable(T.Elt)].Type.(*ast.StructType),
									protocolStructs[makeExprReadable(T.Elt)].Type.(*ast.StructType),
									back,
								)
								if back {
									B += fmt.Sprintf(`for _, v := range %s.%s {
								%s
								%s
								tmp%d = append(tmp%d, %s{%s})
								}
								`, pre, n.Name, lv, lb, At, At, T.Elt, li)
								} else {
									B += fmt.Sprintf(`for _, v := range %s.%s {
								%s
								%s
								tmp%d = append(tmp%d, %s.%s{%s})
								}
								`, pre, n.Name, lv, lb, At, At, packageName, T.Elt, li)
								}
							} else {
								I += n.Name + ": " + pre + "." + n.Name + ", "
							}
						} else {
							if unicode.IsUpper(rune(makeExprReadable(f.Type)[0])) {
								lv, lb, li := makePacketLink(
									structs[makeExprReadable(f.Type)].Type.(*ast.StructType),
									protocolStructs[makeExprReadable(f.Type)].Type.(*ast.StructType),
									makeExprReadable(f.Type),
									pre,
									n.Name,
									back,
								)
								V += lv
								B += lb
								if back {
									I += fmt.Sprintf(`%s: %s`, n.Name, li) + ", "
								} else {
									I += fmt.Sprintf(`%s: %s.%s`, n.Name, packageName, li) + ", "
								}
							} else {

								if makeExprReadable(f.Type) != makeExprReadable(F.Type) {
									fmt.Printf("%s IS NOT SAME TYPE: %s != %s\n", n.Name, makeExprReadable(f.Type), makeExprReadable(F.Type))
								}

								I += n.Name + ": " + pre + "." + n.Name + ", "
							}
						}
					}
				}
			}
		}
	}

	I += "}"

	return
}
