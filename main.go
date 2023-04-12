package main

import (
	"kstyle-be-techtest/database"
	"kstyle-be-techtest/models/migration"
	"kstyle-be-techtest/routes"
)

func main() {
	database.DbConnection()
	migration.RunMigration()

	e := routes.MemberRoutes()

	e.Logger.Fatal(e.Start(":8080"))
}
