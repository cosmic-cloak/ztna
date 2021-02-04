package tests

import "testing"

func Test_CurrentIdentity(t *testing.T) {
	ctx := NewTestContext(t)
	defer ctx.Teardown()
	ctx.StartServer()
	ctx.RequireAdminLogin()

	t.Run("edge routers endpoint", func(t *testing.T) {
		ctx.testContextChanged(t)

		er1 := ctx.createAndEnrollEdgeRouter("test1")

		_ = ctx.AdminSession.requireNewEdgeRouter("test1")

		_, identityAuth := ctx.AdminSession.requireCreateIdentityOttEnrollment("testErAccess", false, "test1")

		identitySession, err := identityAuth.Authenticate(ctx)
		ctx.Req.NoError(err)

		t.Run("returns empty list with no edge router policies", func(t *testing.T) {
			ctx.testContextChanged(t)

			erContainer := identitySession.requireQuery("/current-identity/edge-routers")

			ctx.Req.True(erContainer.ExistsP("data"), "has a data attribute")

			erArray, err := erContainer.Path("data").Children()
			ctx.Req.NoError(err)
			ctx.Req.Len(erArray, 0, "expect empty edge router list")
		})

		t.Run("returns a list of one with an er policy", func(t *testing.T) {
			ctx.testContextChanged(t)
			_ = ctx.AdminSession.requireNewEdgeRouterPolicy([]string{"#test1"}, []string{"#test1"})

			erContainer := identitySession.requireQuery("/current-identity/edge-routers")

			ctx.Req.True(erContainer.ExistsP("data"), "has a data attribute")

			erArray, err := erContainer.Path("data").Children()
			ctx.Req.NoError(err)
			ctx.Req.Len(erArray, 1, "expected edge router list to have one edge router")

			erArray[0].ExistsP("id")
		})

		t.Run("returns empty list with if edge routter is deleted", func(t *testing.T) {
			ctx.testContextChanged(t)

			ctx.AdminSession.requireDeleteEntity(er1)

			erContainer := identitySession.requireQuery("/current-identity/edge-routers")

			ctx.Req.True(erContainer.ExistsP("data"), "has a data attribute")

			erArray, err := erContainer.Path("data").Children()
			ctx.Req.NoError(err)
			ctx.Req.Len(erArray, 0, "expect empty edge router list")
		})
	})
}
