package action

import (
	"log"
	"os/exec"

	"github.com/digitalocean/concourse-resource-library/bootstrap"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/spf13/cobra"
)

var (
	repo, mod, image, tag string
)

func init() {
	newAction.Flags().StringVarP(&repo, "repo", "r", "", "remote repository")
	newAction.Flags().StringVarP(&mod, "mod", "m", "", "go module name")

	// base image flags
	newAction.Flags().StringVarP(&image, "image", "i", "alpine", "base image for Dockerfile")
	newAction.Flags().StringVarP(&tag, "tag", "t", "latest", "image tag for Dockerfile")

	root.AddCommand(newAction)
}

var newAction = &cobra.Command{
	Use:   "new",
	Short: "new bootstraps a Concourse Resource",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := git.PlainInit(args[0], false)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("repository created @ ", args[0])

		if repo != "" {
			r.CreateRemote(&config.RemoteConfig{
				Name: "origin",
				URLs: []string{repo},
			})
			log.Println("remote repository added ", repo)
		}

		p := bootstrap.Project{
			Image:    image,
			ImageTag: tag,
			Module:   mod,
		}
		err = p.Execute(args[0])
		if err != nil {
			log.Println(err)
		}

		if mod != "" {
			cmd := exec.Command("go", "mod", "init", mod)
			cmd.Dir = args[0]
			_, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}
		}
	},
}
