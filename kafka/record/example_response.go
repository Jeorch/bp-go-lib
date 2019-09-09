// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     ExampleRequest.avsc
 *     ExampleResponse.avsc
 */

package record

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type ExampleResponse struct {
	JobId    string
	Progress int64
	Error    string
}

func NewExampleResponseWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := &ExampleResponse{}
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

func DeserializeExampleResponse(r io.Reader) (*ExampleResponse, error) {
	t := NewExampleResponse()

	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	return t, err
}

func NewExampleResponse() *ExampleResponse {
	return &ExampleResponse{}
}

func (r *ExampleResponse) Schema() string {
	return "{\"fields\":[{\"name\":\"JobId\",\"type\":\"string\"},{\"name\":\"Progress\",\"type\":\"long\"},{\"name\":\"Error\",\"type\":\"string\"}],\"name\":\"ExampleResponse\",\"namespace\":\"com.pharbers.kafka.schema\",\"type\":\"record\"}"
}

func (r *ExampleResponse) SchemaName() string {
	return "com.pharbers.kafka.schema.ExampleResponse"
}

func (r *ExampleResponse) Serialize(w io.Writer) error {
	return writeExampleResponse(r, w)
}

func (_ *ExampleResponse) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *ExampleResponse) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *ExampleResponse) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *ExampleResponse) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *ExampleResponse) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *ExampleResponse) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *ExampleResponse) SetString(v string)   { panic("Unsupported operation") }
func (_ *ExampleResponse) SetUnionElem(v int64) { panic("Unsupported operation") }
func (r *ExampleResponse) Get(i int) types.Field {
	switch i {
	case 0:
		return (*types.String)(&r.JobId)
	case 1:
		return (*types.Long)(&r.Progress)
	case 2:
		return (*types.String)(&r.Error)

	}
	panic("Unknown field index")
}
func (r *ExampleResponse) SetDefault(i int) {
	switch i {

	}
	panic("Unknown field index")
}
func (_ *ExampleResponse) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ExampleResponse) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *ExampleResponse) Finalize()                        {}

type ExampleResponseReader struct {
	r io.Reader
	p *vm.Program
}

func NewExampleResponseReader(r io.Reader) (*ExampleResponseReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewExampleResponse()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ExampleResponseReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r *ExampleResponseReader) Read() (*ExampleResponse, error) {
	t := NewExampleResponse()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
