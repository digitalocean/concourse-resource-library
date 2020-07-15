package artifactory

import (
	"github.com/jfrog/jfrog-cli/artifactory/commands/docker"
)

// PullImage pulls a docker image from Artifactory
func (c *Client) PullImage(repo, imageTag string) error {
	cmd := docker.NewDockerPullCommand()
	cmd.SetImageTag(imageTag)
	cmd.SetRepo(repo)

	dets := cliDetails(c)
	cmd.SetRtDetails(&dets)

	return cmd.Run()
}
