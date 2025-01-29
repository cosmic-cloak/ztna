package getziti

import (
	"fmt"
	"strings"

	c "ztna-core/ztna/ztna/constants"

	"github.com/blang/semver"
)

func InstallZitiEdgeTunnel(targetVersion, targetOS, targetArch, binDir string, verbose bool) error {
	var newVersion semver.Version

	if targetVersion != "" {
		newVersion = semver.MustParse(strings.TrimPrefix(targetVersion, "v"))
	} else {
		v, err := GetLatestGitHubReleaseVersion(c.OpenZitiOrg, c.ZITI_EDGE_TUNNEL_GITHUB, verbose)
		if err != nil {
			return err
		}
		newVersion = v
	}

	fmt.Println("Attempting to install '" + c.ZITI_EDGE_TUNNEL + "' version: " + newVersion.String())
	return FindVersionAndInstallGitHubRelease(
		c.OpenZitiOrg, c.ZITI_EDGE_TUNNEL, c.ZITI_EDGE_TUNNEL_GITHUB, targetOS, targetArch, binDir, newVersion.String(), verbose)
}
