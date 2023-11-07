package sshkey

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hetznercloud/cli/internal/testutil"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func TestList(t *testing.T) {
	fx := testutil.NewFixture(t)
	defer fx.Finish()

	time.Local = time.UTC

	cmd := ListCmd.CobraCommand(context.Background(), fx.Client, fx.TokenEnsurer)

	fx.ExpectEnsureToken()
	fx.Client.SSHKeyClient.EXPECT().
		AllWithOpts(
			gomock.Any(),
			hcloud.SSHKeyListOpts{
				ListOpts: hcloud.ListOpts{PerPage: 50},
				Sort:     []string{"id:asc"},
			},
		).
		Return([]*hcloud.SSHKey{
			{
				ID:      123,
				Name:    "test",
				Created: time.Now().Add(-1 * time.Hour),
			},
		}, nil)

	out, err := fx.Run(cmd, []string{})

	expOut := `ID    NAME   FINGERPRINT   AGE
123   test   -             1h
`

	assert.NoError(t, err)
	assert.Equal(t, expOut, out)
}
