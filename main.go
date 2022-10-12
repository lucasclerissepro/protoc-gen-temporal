package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	temporalPackage = protogen.GoImportPath("go.temporal.io/sdk/temporal")

	commentWidth               = 80
	generatedFilenameExtension = ".temporal.go"
	generatedPackageSuffix     = "temporal"
)

func main() {
	protogen.Options{}.Run(func(p *protogen.Plugin) error {
		p.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range p.Files {
			if f.Generate {
				generate(p, f)
			}
		}
		return nil
	})
}

func generate(p *protogen.Plugin, f *protogen.File) error {
	if len(f.Services) == 0 {
		return nil
	}

	gf := p.NewGeneratedFile(
		f.GeneratedFilenamePrefix+generatedFilenameExtension,
		protogen.GoImportPath(path.Join(
			string(f.GoImportPath),
			string(f.GoPackageName),
		)),
	)

	generatePreamble(f, gf)
	generateVersionConstraints(f, gf)

	return nil
}

func generatePreamble(f *protogen.File, gf *protogen.GeneratedFile) error {
	gf.P(fmt.Sprintf("// Code generated by %s. DO NOT EDIT MANUALLY.", filepath.Base(os.Args[0])))
	gf.P("//")

	if f.Proto.GetOptions().GetDeprecated() {
		wrapComments(gf, fmt.Sprintf("%s is a deprecated file.", f.Desc.Path()))
	} else {
		wrapComments(gf, "Source: ", f.Desc.Path())
	}

	gf.P()
	gf.P("package ", f.GoPackageName)

	return nil
}

func generateVersionConstraints(f *protogen.File, gf *protogen.GeneratedFile) error {
	// temporal.SDKVersion

	gf.P("const (")
	gf.P("temporalVersion = ", temporalPackage.Ident("SDKVersion"))
	gf.P(")")

	return nil
}

func wrapComments(gf *protogen.GeneratedFile, elems ...any) {
	text := &bytes.Buffer{}
	for _, el := range elems {
		switch el := el.(type) {
		case protogen.GoIdent:
			fmt.Fprint(text, gf.QualifiedGoIdent(el))
		default:
			fmt.Fprint(text, el)
		}
	}
	words := strings.Fields(text.String())
	text.Reset()
	var pos int
	for _, word := range words {
		numRunes := utf8.RuneCountInString(word)
		if pos > 0 && pos+numRunes+1 > commentWidth {
			gf.P("// ", text.String())
			text.Reset()
			pos = 0
		}
		if pos > 0 {
			text.WriteRune(' ')
			pos++
		}
		text.WriteString(word)
		pos += numRunes
	}
	if text.Len() > 0 {
		gf.P("// ", text.String())
	}
}
