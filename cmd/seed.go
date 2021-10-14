package cmd

import (
	"github.com/nurislam03/golang_redis/data/seeder"
	"github.com/spf13/cobra"
)

func newSeedCmd() *cobra.Command {
	// seedCmd to seed database
	var seedCmd = &cobra.Command{
		Use:   "seed",
		Short: "use seeder for seeding database",
		Long:  `Seed database`,
		Run: func(cmd *cobra.Command, args []string) {
			seeder.Execute(args)
		},
	}

	return seedCmd
}
