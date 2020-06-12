package record

import "io"

/**
{
  "type": "record",
  "name": "MapTest",
  "namespace": "com.pharbers.kafka.schema",
  "fields": [
    {
      "name": "testMap",
      "type": {
        "values": "string",
        "type": "map"
      }
    }
  ]
}
 */

type MapTest struct {
	TestMap map[string]string
}

func (r *MapTest) Schema() string {
	return "{\"fields\":[{\"name\":\"testMap\",\"type\":{\"type\":\"map\",\"values\":\"string\"}}],\"name\":\"MapTest\",\"namespace\":\"com.pharbers.kafka.schema\",\"type\":\"record\"}"
}

func (r *MapTest) Serialize(w io.Writer) error {
	return nil
}
