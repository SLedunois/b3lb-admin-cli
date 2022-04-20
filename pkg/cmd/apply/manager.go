package apply

import (
	"github.com/SLedunois/b3lb/v2/pkg/admin"
	"gopkg.in/yaml.v3"
)

func toResource(in []byte) (string, interface{}, error) {
	var resource *Resource
	if err := yaml.Unmarshal(in, &resource); err != nil {
		return "", nil, err
	}

	return "InstanceList", admin.InstanceList{}, nil
}
