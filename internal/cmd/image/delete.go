package image

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/hetznercloud/cli/internal/cmd/base"
	"github.com/hetznercloud/cli/internal/hcapi2"
	"github.com/hetznercloud/cli/internal/state"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

var DeleteCmd = base.DeleteCmd{
	ResourceNameSingular: "image",
	ShortDescription:     "Delete an image",
	NameSuggestions:      func(c hcapi2.Client) func() []string { return c.Image().Names },
	Fetch: func(ctx context.Context, client hcapi2.Client, cmd *cobra.Command, idOrName string) (interface{}, *hcloud.Response, error) {
		return client.Image().Get(ctx, idOrName)
	},
	Delete: func(ctx context.Context, client hcapi2.Client, _ state.ActionWaiter, cmd *cobra.Command, resource interface{}) error {
		image := resource.(*hcloud.Image)
		if _, err := client.Image().Delete(ctx, image); err != nil {
			return err
		}
		return nil
	},
}
