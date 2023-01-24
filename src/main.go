package main

import (
	"log"

	"github.com/spf13/cobra"
)

var postgresConn string
var snowflakeConn string
var dbName string
var configFile string

func main() {
	var rootCmd = &cobra.Command{
		Use:   "connect",
		Short: "connect to databases",
		Long:  `connect to databases and save the connection strings`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("PostgreSQL connection string:", postgresConn)
			log.Println("Snowflake connection string:", snowflakeConn)
			log.Println("Database name:", dbName)
			log.Println("Configuration file:", configFile)
		},
	}

	rootCmd.Flags().StringVarP(&postgresConn, "postgres", "p", "", "PostgreSQL connection string")
	rootCmd.Flags().StringVarP(&snowflakeConn, "snowflake", "s", "", "Snowflake connection string")
	rootCmd.Flags().StringVarP(&dbName, "dbname", "d", "", "database name")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "config file")
	rootCmd.MarkFlagRequired("postgres")
	rootCmd.MarkFlagRequired("snowflake")
	rootCmd.MarkFlagRequired("dbname")
	rootCmd.MarkFlagRequired("config")

	rootCmd.Execute()
}
