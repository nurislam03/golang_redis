package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	r := newRootCmd()
	if err := r.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	loadConfig()
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "template",
		Short: "A template for the RESTful APIs",
		Long:  `A template for the RESTful APIs`,
	}

	rootCmd.AddCommand(
		newServeCmd(),
		newMigrateCmd(),
		newRoutersCmd(),
		newSeedCmd(),
		newVersionCmd(),
	)

	return rootCmd
}

// initConfig reads in config file and ENV variables if set.
func loadConfig() {
	viper.SetConfigName(".template.config")
	viper.SetConfigType("yaml")

	log.Println("env: " + os.Getenv("TEMPLATE_ENV"))

	env := "local"
	if envVar := os.Getenv("TEMPLATE_ENV"); envVar != "" {
		env = strings.ToLower(envVar)
	}
	viper.Set("server.env", env)

	if env == "local" {
		viper.AddConfigPath("./configs")
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		// Search config in home directory with name ".warden" (without extension).
		viper.AddConfigPath(home + "/configs")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found. Error: %v", err)
		} else {
			log.Fatalf("Config file was found but another error was produced. Error: %v", err)
		}
	}
	log.Println("Using config file:", viper.ConfigFileUsed())

	viper.SetConfigName(fmt.Sprintf(".template.%s.config", env))
	viper.MergeInConfig()
	viper.SetConfigName(fmt.Sprintf("local.%s.config", env))
	viper.MergeInConfig()

	viper.SetEnvPrefix("template")
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	viper.AutomaticEnv()

	setViperIfExistInEnv("server.port", "TEMPLATE_SERVER_PORT")
	setViperIfExistInEnv("database.uri", "TEMPLATE_DATABASE_URI")
	setViperIfExistInEnv("database.name", "TEMPLATE_DATABASE_NAME")
	setViperIfExistInEnv("log.debug", "TEMPLATE_LOG_LEVEL")

	log.Println(viper.AllKeys())
	log.Println(viper.AllSettings())
}

func setViperIfExistInEnv(key, envKey string) {
	if val := os.Getenv(envKey); len(val) > 0 {
		viper.Set(key, val)
	}
}
