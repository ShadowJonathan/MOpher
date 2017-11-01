package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

var protocolPackets []string
var protocolStructs = map[string]*ast.TypeSpec{}

func init() {
	fs := token.NewFileSet()
	pack, err := parser.ParseDir(fs, "../../", func(info os.FileInfo) bool {
		return strings.Contains(info.Name(), "handshaking.go") || strings.Contains(info.Name(), "serverbound.go") || strings.Contains(info.Name(), "clientbound.go")
	}, parser.ParseComments)
	if err != nil {
		log.Fatal("COULD NOT PARSE BECAUSE", err)
	}

	for _, p := range pack {
		for _, file := range p.Files {
			for _, decl := range file.Decls {
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

					protocolStructs[tSpec.Name.Name] = tSpec

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

					protocolPackets = append(protocolPackets, tSpec.Name.Name)
				}
			}
		}
	}

	fmt.Printf("Found packets: %s\n", protocolPackets)
}

func protocolHasPacket(p string) bool {
	for _, P := range protocolPackets {
		if p == P {
			return true
		}
	}
	return false
}
