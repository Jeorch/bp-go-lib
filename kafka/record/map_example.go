// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     Es.avsc
 *     ExampleRequest.avsc
 *     ExampleResponse.avsc
 *     HiveTask.avsc
 *     HiveTracebackTask.avsc
 *     MapExample.avsc
 *     OssTask.avsc
 */

package record

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type MapExample struct {
	TestMap *MapString
}

func NewMapExampleWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := &MapExample{}
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

func DeserializeMapExample(r io.Reader) (*MapExample, error) {
	t := NewMapExample()

	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	return t, err
}

func NewMapExample() *MapExample {
	return &MapExample{}
}

func (r *MapExample) Schema() string {
	return "{\"fields\":[{\"name\":\"testMap\",\"type\":{\"type\":\"map\",\"values\":\"string\"}}],\"name\":\"MapExample\",\"namespace\":\"com.pharbers.kafka.schema\",\"type\":\"record\"}"
}

func (r *MapExample) SchemaName() string {
	return "com.pharbers.kafka.schema.MapExample"
}

func (r *MapExample) Serialize(w io.Writer) error {
	return writeMapExample(r, w)
}

func (_ *MapExample) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *MapExample) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *MapExample) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *MapExample) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *MapExample) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *MapExample) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *MapExample) SetString(v string)   { panic("Unsupported operation") }
func (_ *MapExample) SetUnionElem(v int64) { panic("Unsupported operation") }
func (r *MapExample) Get(i int) types.Field {
	switch i {
	case 0:
		r.TestMap = NewMapString()
		return r.TestMap

	}
	panic("Unknown field index")
}
func (r *MapExample) SetDefault(i int) {
	switch i {

	}
	panic("Unknown field index")
}
func (_ *MapExample) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MapExample) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *MapExample) Finalize()                        {}

type MapExampleReader struct {
	r io.Reader
	p *vm.Program
}

func NewMapExampleReader(r io.Reader) (*MapExampleReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewMapExample()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &MapExampleReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r *MapExampleReader) Read() (*MapExample, error) {
	t := NewMapExample()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
