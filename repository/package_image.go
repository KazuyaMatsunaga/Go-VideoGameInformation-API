package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
)

func GetPkg(db *sqlx.DB, pkg *model.PackageImg) (*model.PackageImg, error) {
	var pkgv model.PackageImg
	if err := db.Get(&pkgv,`SELECT id, package_image_url FROM package_image WHERE package_image_url = ? LIMIT 1`, pkg.PackageImgURL); err != nil {
		return nil, err
	}
	return &pkgv, nil
}

func WritePackageImg(db *sqlx.Tx, pkg *model.PackageImg) (result sql.Result, err error) {
	stmt, err := db.Prepare(`INSERT INTO package_image (package_image_url) VALUES (?)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	return stmt.Exec(pkg.PackageImgURL)
}