package location

import (
	"context"

	"github.com/spf13/pflag"

	"github.com/hetznercloud/cli/internal/cmd/base"
	"github.com/hetznercloud/cli/internal/cmd/output"
	"github.com/hetznercloud/cli/internal/hcapi2"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/hetznercloud/hcloud-go/v2/hcloud/schema"
)

var ListCmd = base.ListCmd{
	ResourceNamePlural: "locations",
	JSONKeyGetByName:   "locations",
	DefaultColumns:     []string{"id", "name", "description", "network_zone", "country", "city"},

	Fetch: func(ctx context.Context, client hcapi2.Client, _ *pflag.FlagSet, listOpts hcloud.ListOpts, sorts []string) ([]interface{}, error) {
		opts := hcloud.LocationListOpts{ListOpts: listOpts}
		if len(sorts) > 0 {
			opts.Sort = sorts
		}
		locations, err := client.Location().AllWithOpts(ctx, opts)

		var resources []interface{}
		for _, n := range locations {
			resources = append(resources, n)
		}
		return resources, err
	},

	OutputTable: func(_ hcapi2.Client) *output.Table {
		return output.NewTable().
			AddAllowedFields(hcloud.Location{})
	},

	Schema: func(resources []interface{}) interface{} {
		locationSchemas := make([]schema.Location, 0, len(resources))
		for _, resource := range resources {
			location := resource.(*hcloud.Location)
			locationSchemas = append(locationSchemas, hcloud.SchemaFromLocation(location))
		}
		return locationSchemas
	},
}
