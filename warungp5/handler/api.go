package handler

import (
	"fmt"
	"log"
	_ "mysql-master"
	"net/http"
	"warungp5/server"

	"github.com/labstack/echo"
)

type menu struct {
	Id_menu     int64  `json:"id_menu"`
	Nama_menu   string `json:"nama_menu"`
	Deskripsi   string `json:"deskripsi"`
	Url_gambar  string `json:"url_gambar"`
	Jenis       string `json:"jenis"`
	Harga       string `json:"harga"`
	Total_order int    `json:"total_order"`
}

var data []menu

func BacaData(c echo.Context) error {
	menu_makanan()

	return c.JSON(http.StatusOK, data)
}

func menu_makanan() {
	data = nil

	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_menu, nama_menu, deskripsi, url_gambar, jenis, harga FROM tbl_menu ORDER BY nama_menu ASC")
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga)
		if err != nil {
			log.Fatal(err.Error())
		}
		data = append(data, each)
		fmt.Println(data)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
}

func BacaDataPopuler(c echo.Context) error {
	menu_populer()

	return c.JSON(http.StatusOK, data)
}

func menu_populer() {
	data = nil

	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_menu, nama_menu, deskripsi, url_gambar, jenis, harga, sum(total_order) as total_order FROM view_total_order GROUP BY id_menu ORDER BY total_order DESC LIMIT 8")
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga, &each.Total_order)
		if err != nil {
			log.Fatal(err.Error())
		}
		data = append(data, each)
		fmt.Println(data)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
}

type jenis_struct struct {
	Jenis string `json:"jenis"`
}

var data_jenis []jenis_struct

func BacaDataJenis(c echo.Context) error {
	jenis_makanan()

	return c.JSON(http.StatusOK, data_jenis)
}

func jenis_makanan() {
	data_jenis = nil

	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT jenis FROM tbl_menu GROUP BY jenis ORDER BY jenis ASC")
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = jenis_struct{}
		var err = rows.Scan(&each.Jenis)
		if err != nil {
			log.Fatal(err.Error())
		}
		data_jenis = append(data_jenis, each)
		fmt.Println(data_jenis)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
}

func TambahData(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var url_gambar = c.FormValue("Url_gambar")
	var jenis = c.FormValue("Jenis")
	var harga = c.FormValue("Harga")

	_, err = db.Exec("INSERT INTO tbl_menu values (?,?,?,?,?,?);", nil, nama, deskripsi, url_gambar, jenis, harga)
	if err != nil {
		log.Print(err.Error())
		return c.JSON(http.StatusOK, "Gagal Menambahkan Menu!")
	} else {
		log.Print("Berhasil Menambahkan Menu!")
		return c.JSON(http.StatusOK, "Berhasil Menambahkan Menu!")
	}
}

func UbahData(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var id_menu = c.FormValue("Id_menu")
	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var url_gambar = c.FormValue("Url_gambar")
	var jenis = c.FormValue("Jenis")
	var harga = c.FormValue("Harga")

	_, err = db.Exec("UPDATE tbl_menu SET nama_menu = ?, deskripsi = ?, url_gambar = ?, jenis = ?, harga = ? WHERE id_menu = ?;", nama, deskripsi, url_gambar, jenis, harga, id_menu)
	if err != nil {
		log.Print(err.Error())
		return c.JSON(http.StatusOK, "Gagal Mengubah Menu!")
	} else {
		log.Print("Berhasil Mengubah Menu!")
		return c.JSON(http.StatusOK, "Berhasil Mengubah Menu!")
	}
}

func HapusData(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var id_menu = c.FormValue("Id_menu")

	_, err = db.Exec("DELETE FROM tbl_menu WHERE id_menu = ?;", id_menu)
	if err != nil {
		log.Print(err.Error())
		return c.JSON(http.StatusOK, "Gagal Menghapus Menu!")
	} else {
		log.Print("Berhasil Menghapus Menu!")
		return c.JSON(http.StatusOK, "Berhasil Menghapus Menu!")
	}
}

func BacaSatuData(c echo.Context) error {

	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var id_menu = c.FormValue("Id_menu")

	row := db.QueryRow("SELECT id_menu, nama_menu, deskripsi, url_gambar, jenis, harga FROM tbl_menu WHERE id_menu = ?", id_menu)
	var each = menu{}
	var scan = row.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga)

	if scan != nil {
		log.Print(scan.Error())
		return c.JSON(http.StatusOK, "ID Menu tidak di temukan!")
	} else {
		log.Print(each)
		return c.JSON(http.StatusOK, each)
	}
}

func InputOrder(c echo.Context) error {
	db, err := server.Koneksi()
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var id_menu = c.FormValue("id_menu")

	var nama_pemesan = c.FormValue("nama_pemesan")
	var nomor_telepon = c.FormValue("nomor_telepon")
	var jumlah = c.FormValue("jumlah")
	var alamat = c.FormValue("alamat")

	_, err = db.Exec("INSERT INTO tbl_order values (?,?,?,?,?,?);", nil, id_menu, nama_pemesan, nomor_telepon, alamat, jumlah)
	if err != nil {
		log.Print(err.Error())
		return c.HTML(http.StatusOK, "<strong>Gagal Melakukan Pemesanan!</strong>")
	} else {
		log.Print("Berhasil Melakukan Pemesanan!")
		return c.HTML(http.StatusOK, "<script>alert(`Berhasil Melakukan Pemesanan!... Silahkan Tunggu Telepon Dari Kami. Terimakasih!`); window.location = window.location.origin;</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
