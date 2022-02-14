package command

import (
	"log"
	"time"

	"github.com/smhdhsn/food/internal/config"
	"github.com/smhdhsn/food/internal/db"
	"github.com/smhdhsn/food/internal/repository/mysql"
	"github.com/smhdhsn/food/internal/service"
	"github.com/spf13/cobra"
)

var (
	amount     = uint(3)                     // the amount of items being added to inventory with every buy.
	bestBefore = time.Now().AddDate(0, 2, 0) // item's best usage time.
	expiresAt  = time.Now().AddDate(0, 5, 0) // item's expiration time.
)

// buyCMD is the subcommands responsible for creating food components.
var buyCMD = &cobra.Command{
	Use:   "buy",
	Short: "Stores new food components inside database if their components' stock are finished or expired.",
	Run:   buyRun,
}

// init function will be executed when this package is called.
func init() {
	rootCMD.AddCommand(buyCMD)

	buyCMD.Flags().UintP("amount", "a", amount, "The amount of stocks added to inventory after each buy.")
}

// buyRun is responsible for running the script.
func buyRun(cmd *cobra.Command, args []string) {
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.Connect(conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	iRepo := mysql.NewInventoryRepo(dbConn)
	cRepo := mysql.NewComponentRepo(dbConn)

	iService := service.NewInventoryService(iRepo, cRepo)

	a, err := cmd.Flags().GetUint("amount")
	if err != nil {
		log.Fatal(err)
	}

	req := service.BuyComponentsReq{
		StockAmount: a,
		BestBefore:  bestBefore,
		ExpiresAt:   expiresAt,
	}

	err = iService.BuyComponents(&req)
	if err != nil {
		log.Fatal(err)
	}
}