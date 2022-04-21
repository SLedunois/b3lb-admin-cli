package apply

import (
	"errors"

	"github.com/SLedunois/b3lb/v2/pkg/admin"
	"gopkg.in/yaml.v3"
)

func toResource(in []byte) (string, interface{}, error) {
	var resource *Resource
	if err := yaml.Unmarshal(in, &resource); err != nil {
		return "", nil, err
	}

	switch resource.Kind {
	case "InstanceList":
		return "InstanceList", admin.InstanceList{}, nil
	case "Tenant":
		return "Tenant", admin.Tenant{}, nil
	default:
		return "", nil, errors.New("unknown resource kind")
	}
}
