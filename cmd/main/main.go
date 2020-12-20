package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	//"course-phones-review/gadgets/smartphones/web"
	"course-phones-review/internal/database"
	"course-phones-review/internal/logs"
	//reviews "github.com/tomiok/course-phones-review/reviews/web"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()

	client := database.NewSqlClient("app_phone_review:qwerty12345@tcp(localhost:3306)/phones_review")
	doMigrate(client, "phones_review")

	/*mongoClient := database.NewMongoClient("localhost")

	reviewHandler := reviews.NewReviewHandler(mongoClient)*/

	//handler := web.NewCreateSmartphoneHandler(client)
	//mux := Routes(handler, reviewHandler)
	mux := Routes()
	server := NewServer(mux)
	server.Run()
}

func doMigrate(client *database.MySqlClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)		
	if err != nil && err.Error() == "no change" {		
		logs.Log().Info("no migration needed")
	}
}

