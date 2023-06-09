package test

import (
	"fmt"
	"testing"

	"github.com/dimasardnt6/BE-Antrian-Puskesmas/model"
	"github.com/dimasardnt6/BE-Antrian-Puskesmas/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test Insert
func TestInsertPasien(t *testing.T) {
	nama_pasien := "Dimas Ardianto"
	nomor_ktp := "3217060601020007"
	alamat := "Bandung Barat"
	nomor_telepon := "089647129890"
	tanggal_lahir := "6 Januari 2002"
	jenis_kelamin := "Laki-Laki"
	insertedID, err := module.InsertPasien(module.MongoConn, "data_pasien", nama_pasien, nomor_ktp, alamat, nomor_telepon, tanggal_lahir, jenis_kelamin)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertAntrian(t *testing.T) {
	poli := model.Poliklinik{
		Kode_Poliklinik: "PLUM",
		Nama_Poliklinik: "Poli Umum",
		Deskripsi:       "memberikan pelayanan kedokteran berupa pemeriksaan kesehatan, pengobatan dan penyuluhan kepada pasien atau masyarakat",
	}
	identitas_pasien := model.Pasien{
		Nama_Pasien:   "Dimas Ardianto",
		Nomor_Ktp:     "3217060601020007",
		Alamat:        "Bandung Barat",
		Nomor_Telepon: "089647129890",
		Tanggal_Lahir: "06 Januari 2002",
		Jenis_Kelamin: "Laki-Laki",
	}
	nomor_antrian := 1
	status_antrian := "Menunggu"
	insertedID, err := module.InsertAntrian(module.MongoConn, "data_antrian", poli, identitas_pasien, nomor_antrian, status_antrian)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertPoliklinik(t *testing.T) {
	kode_poliklinik := "PLUM"
	nama_poliklinik := "Poliklinik Umum"
	deskripsi := "memberikan pelayanan kedokteran berupa pemeriksaan kesehatan, pengobatan dan penyuluhan kepada pasien atau masyarakat"
	insertedID, err := module.InsertPoliklinik(module.MongoConn, "data_poliklinik", kode_poliklinik, nama_poliklinik, deskripsi)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertDokter(t *testing.T) {
	nama_dokter := "Dr.Johnson"
	spesialisasi := "Dokter Sepesialis Anak"
	insertedID, err := module.InsertDokter(module.MongoConn, "data_dokter", nama_dokter, spesialisasi)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

// Test Get
func TestGetPasienFromID(t *testing.T) {
	id := "6482eaab8de5676bccccab77"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPasienFromID(objectID, module.MongoConn, "data_pasien")
	if err != nil {
		t.Fatalf("error calling GetPasienFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAntrianFromID(t *testing.T) {
	id := "6482d2e36957c6996b458aa8"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetAntrianFromID(objectID, module.MongoConn, "data_antrian")
	if err != nil {
		t.Fatalf("error calling GetAntrianFromID: %v", err)
	}
	fmt.Println(biodata)
}
func TestGetPoliklinikFromID(t *testing.T) {
	id := "6482d6430fd934b1a3d071ec"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	poliklinik, err := module.GetPoliklinikFromID(objectID, module.MongoConn, "data_poliklinik")
	if err != nil {
		t.Fatalf("error calling GetPoliklinikFromID: %v", err)
	}
	fmt.Println(poliklinik)
}

func TestGetDokterFromID(t *testing.T) {
	id := "6482d9ec9704737c2f981105"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	dokter, err := module.GetDokterFromID(objectID, module.MongoConn, "data_dokter")
	if err != nil {
		t.Fatalf("error calling GetDokterFromID: %v", err)
	}
	fmt.Println(dokter)
}

func TestGetAllPasien(t *testing.T) {
	data := module.GetAllPasien(module.MongoConn, "data_pasien")
	fmt.Println(data)
}

func TestGetAllAntrian(t *testing.T) {
	data := module.GetAllAntrian(module.MongoConn, "data_antrian")
	fmt.Println(data)
}

func TestGetAllPoliklinik(t *testing.T) {
	data := module.GetAllPoliklinik(module.MongoConn, "data_poliklinik")
	fmt.Println(data)
}

func TestGetAllDokter(t *testing.T) {
	data := module.GetAllDokter(module.MongoConn, "data_dokter")
	fmt.Println(data)
}

// Test Delete

func TestDeleteAntrianByID(t *testing.T) {
	id := "6482c069ef5013c1dc9f16a0" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteAntrianByID(objectID, module.MongoConn, "data_antrian")
	if err != nil {
		t.Fatalf("error calling DeleteAntrianByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetAntrianFromID
	_, err = module.GetAntrianFromID(objectID, module.MongoConn, "data_antrian")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
