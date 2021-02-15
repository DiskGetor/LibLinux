package tag

import (
	"bufio"
	"fmt"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
	"github.com/DiskGetor/liblinux/tool"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	IsPb2            = true
	TemplateFileName = `template.go`
	W                = new(struct {
		log.ObjLog
		errcheck.ObjErrCheck
		tool.ObjTool
	})
)

type (
	InterfaceTag interface {
		AddPbTag()
		AddfbTag()
	}
	ObjTag struct{}
)

func init() {
	if _, err := os.Open(TemplateFileName); err != nil {
		//W.WriteFile(TemplateFileName, `package main`)
	}
}

func (ObjTag) AddPbTag() {
	p := NewParser()
	p.UseFilter(FilterStruct)
	p.ParseFile(TemplateFileName, `protobuf`)
}

func (ObjTag) AddfbTag() {
	p := NewParser()
	p.UseFilter(FilterStruct)
	p.ParseFile(TemplateFileName, `flatbuf`)
}

func NewParser() *ParserTag {
	p := &ParserTag{nodeFilter: new(nodeFilter)}
	p.nodeFilter = new(nodeFilter)
	return p
}

func FilterStruct(f *ast.File, Key string) {
	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.StructType:
			parseStruct(t, Key)
			return false
		}
		return true
	})
}

func parseStruct(n *ast.StructType, Key string) {
	if n.Fields.NumFields() == 0 {
		return
	}
	p := new(TagCtx)
	for i, field := range n.Fields.List {
		if field.Tag != nil {
			if strings.Contains(field.Tag.Value, `pack`) {
				p.IsPacked = true
			}
		}
		ident, ok := field.Type.(*ast.StarExpr)
		if !ok {
			W.CheckErr(`get field err`)
			return
		}
		p.Kind = W.TypeStr2Kind(fmt.Sprint(ident.X))
		p.ID = i + 1
		p.FieldName = field.Names[0].String()
		p.Key = Key
		p.resetTag()
		field.Tag = p.Tag
	}
}

type TagCtx struct {
	Key       string
	FieldName string
	ID        int
	Tag       *ast.BasicLit
	Kind      reflect.Kind
	Wire      string
	Option    string
	IsPacked  bool
}

func (p *TagCtx) resetTag() {
	p.Tag = new(ast.BasicLit)
	p.Tag.Kind = token.STRING
	p.GoType()
	p.IsPacked = false
	version := ``
	if IsPb2 {
		version = `proto2`
	} else {
		version = `proto3`
	}
	p.Tag.Value =
		"`" +
			p.Key +
			`:` +
			strconv.Quote(
				p.Wire+
					`,`+
					fmt.Sprint(p.ID)+
					`,`+
					p.Option+
					`,name=`+
					p.FieldName+
					`,`+version) + "`"
}

func (p *TagCtx) GoType() {
	switch p.Kind {
	case reflect.Float32:
		p.Wire = "fixed32"
		p.Option = `req`
	case reflect.Float64:
		p.Wire = "fixed64"
		p.Option = `req`
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		p.Wire = "varint"
		p.Option = `req`
	case reflect.Struct:
		if p.IsPacked {
			p.Wire = "bytes"
		} else {
			p.Wire = "group"
		}
		p.Option = `req`
	case reflect.Array, reflect.String, reflect.Map:
		p.Wire = "bytes"
		p.Option = `req`
	case reflect.Slice:
		p.Wire = "bytes"
		p.Option = `rep`
	default:
		W.CheckErr(`unknown type for ` + fmt.Sprint(p.Kind))
		return
	}
}

func (p *ParserTag) ParseFile(fileName string, Key string) {
	FileSet := token.NewFileSet()
	f, err := parser.ParseFile(FileSet, fileName, nil, parser.ParseComments)
	if !W.CheckErr(err) {
		return
	}
	p.Run(f, Key)
	writeFile(FileSet, f, fileName)
}

func (n *nodeFilter) UseFilter(f FilterNode) {
	if n.Stack == nil {
		n.Stack = make([]FilterNode, 0)
	}
	n.Stack = append(n.Stack, f)
}

type ParserTag struct {
	*nodeFilter
}

type FilterNode func(*ast.File, string)

type nodeFilter struct {
	Stack []FilterNode
}

func (n *nodeFilter) Run(f *ast.File, Key string) {
	for _, filter := range n.Stack {
		filter(f, Key)
	}
}

func writeFile(FileSet *token.FileSet, AstFile *ast.File, fileName string) {
	f, err := os.Create(fileName)
	if !W.CheckErr(err) {
		return
	}
	defer func() { W.CheckErr(f.Close()) }()
	Writer := bufio.NewWriter(f)
	if !W.CheckErr(format.Node(Writer, FileSet, AstFile)) {
		return
	}
	W.CheckErr(Writer.Flush())
}
