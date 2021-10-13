package dbconn

import (
	"github.com/kamva/mgm/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect returns a postgres connection pool.
func Connect() error {
	return mgm.SetDefaultConfig(
		nil,
		viper.GetString("database.name"),
		options.Client().ApplyURI(viper.GetString("database.uri")),
	)
}
