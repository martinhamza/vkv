package printer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintJSON(t *testing.T) {
	testCases := []struct {
		name     string
		s        map[string]interface{}
		rootPath string
		opts     []Option
		output   string
		err      bool
	}{
		{
			name:     "test: normal map to json",
			rootPath: "root",
			s: map[string]interface{}{
				"secret": map[string]interface{}{
					"key":  "value",
					"user": "password",
				},
			},
			opts: []Option{
				ToFormat(JSON),
				ShowValues(true),
			},
			output: `{
  "root/": {
    "secret": {
      "key": "value",
      "user": "password"
    }
  }
}
`,
		},
		{
			name:     "test: normal map to json only keys",
			rootPath: "root",
			s: map[string]interface{}{
				"secret": map[string]interface{}{
					"key":  "value",
					"user": "password",
				},
			},
			opts: []Option{
				ToFormat(JSON),
				OnlyKeys(true),
			},
			output: `{
  "root/": {
    "secret": {
      "key": "",
      "user": ""
    }
  }
}
`,
		},
	}

	for _, tc := range testCases {
		var b bytes.Buffer
		tc.opts = append(tc.opts, WithWriter(&b))

		p := NewPrinter(tc.opts...)

		m := map[string]interface{}{}

		m[tc.rootPath+"/"] = tc.s
		assert.NoError(t, p.Out(m))
		assert.Equal(t, tc.output, b.String(), tc.name)
	}
}
