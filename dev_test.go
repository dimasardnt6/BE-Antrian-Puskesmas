package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertAntrianAntrian(t *testing.T) {
	identitas_pasien := Pasien{
		Nama_Pasien:   "Dimas Ardianto",
		Nomor_Ktp:     "3217060601020007",
		Alamat:        "Bandung Barat",
		Nomor_Telepon: "089647129890",
		Tanggal_Lahir: "06 Januari 2002",
		Jenis_Kelamin: "Laki-Laki",
	}
	nomor_antrian := "1"
	waktu_pendaftaran := "08.00"
	status_antrian := []string{"Menunggu"}
	hasil := InsertAntrian(identitas_pasien, nomor_antrian, waktu_pendaftaran, status_antrian)
	fmt.Println(hasil)
}

func TestGetAntrianFromNomorAntrian(t *testing.T) {
	nomor_antrian := "1"
	data := GetAntrianFromNomorAntrian(nomor_antrian)
	fmt.Println(data)
}
