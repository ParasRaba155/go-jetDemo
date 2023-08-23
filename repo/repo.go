package repo

import (
	"log"

	"github.com/TntraParas/jet-demo/jet/public/table"
	"github.com/go-jet/jet/v2/qrm"
)

func CreateUser(db qrm.Executable) {
	stmt2 := table.Users.INSERT(
		table.Users.Email, table.Users.PasswordHash, table.Users.FullName,
	).VALUES(
		"paras.raba@tntra.io", "raba@paras", "Paras Raba",
	)
	log.Printf("stmt2 :%+v", stmt2)

	_, err := stmt2.Exec(db)
	if err != nil {
		log.Printf("CreateUser: error in exec db: %s", err)
	}
}
