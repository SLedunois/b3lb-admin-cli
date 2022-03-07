package instances

// ListFlags contains all List command flags
type ListFlags struct {
	CSV  bool
	JSON bool
}

// NewListFlags initialize a new ListFlags object
func NewListFlags() *ListFlags {
	return &ListFlags{
		CSV:  false,
		JSON: false,
	}
}

// ApplyFlags apply ListFlags to provided command
func (cmd *ListCmd) ApplyFlags() {
	cmd.Command.Flags().BoolVarP(&cmd.Flags.CSV, "csv", "c", cmd.Flags.CSV, "csv output")
	cmd.Command.Flags().BoolVarP(&cmd.Flags.JSON, "json", "j", cmd.Flags.JSON, "json output")
}

// AddFlags contains all Add command flags
type AddFlags struct {
	URL    string
	Secret string
}

// NewAddFlags initialize a new AddFlags object
func NewAddFlags() *AddFlags {
	return &AddFlags{}
}

// ApplyFlags apply command flags to cobra command
func (cmd *AddCmd) ApplyFlags() {
	cmd.Command.Flags().StringVarP(&cmd.Flags.URL, "url", "u", "", "BigBlueButton instance url")
	cmd.Command.MarkFlagRequired("url")
	cmd.Command.Flags().StringVarP(&cmd.Flags.Secret, "secret", "s", "", "BigBlueButton secret")
	cmd.Command.MarkFlagRequired("secret")
}
