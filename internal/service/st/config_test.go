package st_test

import (
	"context"
	"gf-boilerplate/internal/service/st"
	"testing"

	"github.com/gogf/gf/v2/frame/gins"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	ctx = context.Background()
)

func TestConfig(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//nolint:errcheck
		gins.Config().GetAdapter().(*gcfg.AdapterFile).SetFileName(gtest.DataPath("config.test.yaml"))
		config := st.GetConfig()
		t.Assert(config.GetDebug(ctx), false)
		err := config.SetDebug(ctx, true)
		t.Assert(err, nil)
		t.Assert(config.GetDebug(ctx), true)
	})
}
