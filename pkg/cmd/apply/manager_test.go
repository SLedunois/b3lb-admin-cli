package apply

import (
	"reflect"
	"testing"

	"github.com/SLedunois/b3lb/v2/pkg/admin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestToResource(t *testing.T) {
	var in []byte
	tests := []struct {
		name      string
		mock      func()
		validator func(*testing.T, string, interface{}, error)
	}{
		{
			name: "passing a kind InstanceList should return an InstanceList",
			mock: func() {
				instanceList := &admin.InstanceList{
					Kind:      "InstanceList",
					Instances: map[string]string{},
				}

				out, err := yaml.Marshal(instanceList)
				if err != nil {
					t.Fatal(err)
				}

				in = out
			},
			validator: func(t *testing.T, kind string, resource interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "InstanceList", kind)
				assert.Equal(t, "admin.InstanceList", reflect.ValueOf(resource).Type().String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			kind, obj, err := toResource(in)
			test.validator(t, kind, obj, err)
		})
	}
}
