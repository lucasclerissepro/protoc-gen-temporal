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
	stringsPackage  = protogen.GoImportPath("strings")
	strconvPackage  = protogen.GoImportPath("strconv")
	contextPackage  = protogen.GoImportPath("context")
	fmtPackage      = protogen.GoImportPath("fmt")

	commentWidth               = 80
	generatedFilenameExtension = ".temporal.go"
	generatedPackageSuffix     = "temporal"

	minimumSupportedMajor = 1
	minimumSupportedMinor = 17
	minimumSupportedPatch = 0
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

	for _, s := range f.Services {
		generateClientInterface(f, gf, s)
		generateClientImplementation(f, gf, s)
	}

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
	gf.P("temporalSDKVersion = ", temporalPackage.Ident("SDKVersion"))
	gf.P(")")

	wrapComments(gf, `Runtime validation of temporal SDK version. 
  The overhead is minimal and only run one when the package is loaded. This is 
  the only way we've found so far to properly ensure this generated code is compatible 
  with the expected Temporal SDK version.

  We're currently trying to open a PR on the Go temporal SDK to add some simple constant
  simplifying code generation and hopefully we'll be able to switch to comptime assertion.
  `)

	gf.P("func init() {")
	gf.P("temporalVersionComponents := ", stringsPackage.Ident("Split(temporalSDKVersion, \".\")"))
	gf.P("major, err := ", strconvPackage.Ident("Atoi(temporalVersionComponents[0])"))
	gf.P("if err != nil {")
	gf.P("panic(\"failed to convert major SDK version to integer\")")
	gf.P("}")

	gf.P("minor, err := ", strconvPackage.Ident("Atoi(temporalVersionComponents[1])"))
	gf.P("if err != nil {")
	gf.P("panic(\"failed to convert minor SDK version to integer\")")
	gf.P("}")

	gf.P("patch, err := ", strconvPackage.Ident("Atoi(temporalVersionComponents[2])"))
	gf.P("if err != nil {")
	gf.P("panic(\"failed to convert patch SDK version to integer\")")
	gf.P("}")
	gf.P()

	gf.P("panicPrefix := \"Temporal SDK (\" + temporalSDKVersion + \")\"")
	gf.P()

	gf.P("if major < ", minimumSupportedMajor, "|| major > ", minimumSupportedMajor, " {")
	gf.P("panic(panicPrefix + \" major version is not supported expect minimum ", minimumSupportedMajor, " major version\")")
	gf.P("}")

	gf.P("if minor < ", minimumSupportedMinor, " {")
	gf.P("panic(panicPrefix + \" minor version is not supported expect minimum ", minimumSupportedMinor, " minor version\")")
	gf.P("}")

	gf.P("if patch < ", minimumSupportedPatch, " {")
	gf.P("panic(panicPrefix + \" patch version is not supported expect minimum ", minimumSupportedPatch, " patch version\")")
	gf.P("}")

	gf.P("}")

	return nil
}

func generateClientInterface(f *protogen.File, gf *protogen.GeneratedFile, s *protogen.Service) error {
	gf.P()
	gf.P("type Client interface { ")

	for _, m := range s.Methods {
		gf.P(m.GoName, "(ctx ", contextPackage.Ident("Context"), ", ", "req *", m.Input.GoIdent.GoName, ") error")
	}

	gf.P("}")
	gf.P()
	return nil
}

func generateClientImplementation(f *protogen.File, gf *protogen.GeneratedFile, s *protogen.Service) error {
	gf.P("type client struct {}")
	gf.P()

	for _, m := range s.Methods {
		gf.P("func (c *client) ", m.GoName, "(ctx ", contextPackage.Ident("Context"), ", ", "req *", m.Input.GoIdent.GoName, ") error {")
		gf.P("return nil")
		gf.P("}")
		gf.P()
	}

	gf.P("func New", s.GoName, "Client() Client {")
	gf.P("return &client{}")
	gf.P("}")
	gf.P()
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
