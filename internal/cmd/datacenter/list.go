package datacenter

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
	ResourceNamePlural: "Datacenters",
	JSONKeyGetByName:   "datacenters",
	DefaultColumns:     []string{"id", "name", "description", "location"},

	Fetch: func(ctx context.Context, client hcapi2.Client, _ *pflag.FlagSet, listOpts hcloud.ListOpts, sorts []string) ([]interface{}, error) {
		opts := hcloud.DatacenterListOpts{ListOpts: listOpts}
		if len(sorts) > 0 {
			opts.Sort = sorts
		}
		datacenters, err := client.Datacenter().AllWithOpts(ctx, opts)
		var resources []interface{}
		for _, n := range datacenters {
			resources = append(resources, n)
		}
		return resources, err
	},

	OutputTable: func(_ hcapi2.Client) *output.Table {
		return output.NewTable().
			AddAllowedFields(hcloud.Datacenter{}).
			AddFieldFn("location", output.FieldFn(func(obj interface{}) string {
				datacenter := obj.(*hcloud.Datacenter)
				return datacenter.Location.Name
			}))
	},

	Schema: func(resources []interface{}) interface{} {
		dcSchemas := make([]schema.Datacenter, 0, len(resources))
		for _, resource := range resources {
			dc := resource.(*hcloud.Datacenter)
			dcSchemas = append(dcSchemas, hcloud.SchemaFromDatacenter(dc))
		}

		return dcSchemas
	},
}
