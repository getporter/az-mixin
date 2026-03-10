package az

import (
	_ "embed"
	"fmt"
)

//go:embed schema/schema.json
var schema string

func (m *Mixin) PrintSchema() error {
	_, err := fmt.Fprint(m.Out, schema)
	return err
}
