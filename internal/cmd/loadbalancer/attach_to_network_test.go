package loadbalancer

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hetznercloud/cli/internal/testutil"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func TestAttachToNetwork(t *testing.T) {
	fx := testutil.NewFixture(t)
	defer fx.Finish()

	cmd := AttachToNetworkCmd.CobraCommand(
		context.Background(),
		fx.Client,
		fx.TokenEnsurer,
		fx.ActionWaiter)
	fx.ExpectEnsureToken()

	fx.Client.LoadBalancerClient.EXPECT().
		Get(gomock.Any(), "123").
		Return(&hcloud.LoadBalancer{ID: 123}, nil, nil)
	fx.Client.NetworkClient.EXPECT().
		Get(gomock.Any(), "my-network").
		Return(&hcloud.Network{ID: 321}, nil, nil)
	fx.Client.LoadBalancerClient.EXPECT().
		AttachToNetwork(gomock.Any(), &hcloud.LoadBalancer{ID: 123}, hcloud.LoadBalancerAttachToNetworkOpts{
			Network: &hcloud.Network{ID: 321},
		}).
		Return(&hcloud.Action{ID: 123}, nil, nil)
	fx.ActionWaiter.EXPECT().
		ActionProgress(gomock.Any(), &hcloud.Action{ID: 123}).
		Return(nil)

	out, _, err := fx.Run(cmd, []string{"123", "--network", "my-network"})

	expOut := "Load Balancer 123 attached to network 321\n"

	assert.NoError(t, err)
	assert.Equal(t, expOut, out)
}
