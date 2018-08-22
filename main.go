package main

//
// Usage ./release <owner> <repo> <tag_name> <file>

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

func validateArgs(a cli.Args) bool {
	if !a.Present() {
		return false
	}

	if len(a) < 4 {
		return false
	}

	return true
}

func main() {
	app := cli.NewApp()
	app.Name = "release"
	app.Usage = "Create and Release files to GitHub Releases"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Value:  "",
			Usage:  "Your GitHub Access Token",
			EnvVar: "GITHUB_ACCESS_TOKEN",
		},
	}
	app.Action = func(c *cli.Context) error {
		if !validateArgs(c.Args()) {
			return fmt.Errorf("missing args")
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: c.String("token")},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		owner := c.Args().Get(0)
		repo := c.Args().Get(1)
		tagName := c.Args().Get(2)
		filePath := c.Args().Get(3)
		preRelease := true

		release, _, err := client.Repositories.GetReleaseByTag(ctx, owner, repo, tagName)
		if err != nil {
			fmt.Printf("Creating new release for %s/%s @ %s...\n", owner, repo, tagName)
			release = &github.RepositoryRelease{
				TagName:    &tagName,
				Name:       &tagName,
				Prerelease: &preRelease,
			}
			release, _, err = client.Repositories.CreateRelease(ctx, owner, repo, release)
			if err != nil {
				return err
			}
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		fmt.Printf("Uploading Release Asset %s as %s...\n", filePath, filepath.Base(filePath))
		_, _, err = client.Repositories.UploadReleaseAsset(
			ctx,
			owner,
			repo,
			*release.ID,
			&github.UploadOptions{
				Name: filepath.Base(filePath),
			},
			file,
		)
		if err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
