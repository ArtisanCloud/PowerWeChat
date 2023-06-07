package power

import (
	"bytes"
	"encoding/json"
	"io"
)

type JsonEncoder struct {
	json.Encoder
	Data interface{}
}

func (e *JsonEncoder) Encode() (io.Reader, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(e.Data); err != nil {
		return nil, err
	}

	reader := bytes.NewReader(buf.Bytes())
	return reader, nil
}
