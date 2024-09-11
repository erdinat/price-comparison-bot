package cmd

import (
	"database/sql"
	"fmt"
	"github.com/erdinat/internProjectGolang/internal/scraper"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch product data from websites and update the database",
	Long:  `fetches product data from all sites in the Sites table and updates the products,product_price_diff,product_price_diff_log`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		dbUsername := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
		db, err := sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatal("Error connecting to database:", err)
		}
		defer db.Close()
		
		scraper.ScrapeSite("Amazon", "https://www.amazon.com.tr", db)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
