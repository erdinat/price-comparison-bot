package main

import (
	"fmt"
	"github.com/erdinat/internProjectGolang/cmd"
	"github.com/erdinat/internProjectGolang/internal/database"
)

func main() {
	db := database.Conn()
	defer db.Close()

	fmt.Println("Veritabanı bağlantısı başarılı!")

	cmd.Execute()
}
