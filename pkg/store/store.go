package store

import (
	"context"
	"embed"

	"github.com/cristalhq/dbump"
	"github.com/cristalhq/dbump/dbump_pgx"
	"github.com/jackc/pgx/v4"
)

//go:embed migrations/*
var migrations embed.FS

func RunMigrations(ctx context.Context, conn *pgx.Conn) error {
	return dbump.Run(ctx, dbump.Config{
		Migrator: dbump_pgx.NewMigrator(conn, dbump_pgx.Config{}),
		Loader:   dbump.NewFileSysLoader(migrations, "migrations"),
		Mode:     dbump.ModeApplyAll,
	})
}
