package init

type InitConfigFlags struct {
	B3LB        string
	APIKey      string
	Destination string
}

func NewInitConfigFlags() *InitConfigFlags {
	return &InitConfigFlags{}
}
