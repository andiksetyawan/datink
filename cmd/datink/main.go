package main

import (
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"

	"datink/cmd/datink/migrations"
	"datink/config"
	application "datink/internal/app"
	"datink/internal/resource/external"
)

//	@title			DATINK SERVICES
//	@description	This is a doc for Invoice.

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	app := &cli.App{
		Name: "datink",
		Commands: []*cli.Command{
			apiCommand,
			newDBCommand(migrations.Migrations),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var apiCommand = &cli.Command{
	Name:  "server",
	Usage: "start API server",
	Action: func(c *cli.Context) error {
		application.RunServer()
		return nil
	},
}

func newDB() *bun.DB {
	cfg, _ := config.NewConfig()
	logger, _ := external.NewLogger(cfg)
	return external.NewDatabase(cfg, logger).GetMaster()
}

func newDBCommand(migrations *migrate.Migrations) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "init migration tables",
				Action: func(c *cli.Context) error {
					migrator := migrate.NewMigrator(newDB(), migrations)
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					migrator := migrate.NewMigrator(newDB(), migrations)

					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context) //nolint:errcheck

					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to run (database is up to date)\n")
						return nil
					}
					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
		},
	}

}
