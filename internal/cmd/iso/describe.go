package iso

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/hetznercloud/cli/internal/cmd/base"
	"github.com/hetznercloud/cli/internal/cmd/util"
	"github.com/hetznercloud/cli/internal/hcapi2"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

// DescribeCmd defines a command for describing a iso.
var DescribeCmd = base.DescribeCmd{
	ResourceNameSingular: "iso",
	ShortDescription:     "Describe a iso",
	JSONKeyGetByID:       "iso",
	JSONKeyGetByName:     "isos",
	NameSuggestions:      func(c hcapi2.Client) func() []string { return c.Location().Names },
	Fetch: func(ctx context.Context, client hcapi2.Client, cmd *cobra.Command, idOrName string) (interface{}, interface{}, error) {
		iso, _, err := client.ISO().Get(ctx, idOrName)
		if err != nil {
			return nil, nil, err
		}
		return iso, hcloud.SchemaFromISO(iso), nil
	},
	PrintText: func(ctx context.Context, client hcapi2.Client, cmd *cobra.Command, resource interface{}) error {
		iso := resource.(*hcloud.ISO)

		cmd.Printf("ID:\t\t%d\n", iso.ID)
		cmd.Printf("Name:\t\t%s\n", iso.Name)
		cmd.Printf("Description:\t%s\n", iso.Description)
		cmd.Printf("Type:\t\t%s\n", iso.Type)
		cmd.Printf(util.DescribeDeprecation(iso))

		architecture := "-"
		if iso.Architecture != nil {
			architecture = string(*iso.Architecture)
		}
		cmd.Printf("Architecture:\t%s\n", architecture)

		return nil
	},
}
