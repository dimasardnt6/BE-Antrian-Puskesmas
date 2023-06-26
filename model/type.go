package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fullname        string             `bson:"fullname,omitempty" json:"firstname,omitempty"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Confirmpassword string             `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
}

type Pasien struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Pasien   string             `bson:"nama_pasien,omitempty" json:"nama_pasien,omitempty"`
	Nomor_Ktp     string             `bson:"nomor_ktp,omitempty" json:"nomor_ktp,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Nomor_Telepon string             `bson:"nomor_telepon,omitempty" json:"nomor_telepon,omitempty"`
	Tanggal_Lahir string             `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Jenis_Kelamin string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty"`
}

type Antrian struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Poli                Poliklinik         `bson:"poli,omitempty" json:"poli,omitempty"`
	Identitas_Pasien    Pasien             `bson:"identitas_pasien,omitempty" json:"identitas_pasien,omitempty"`
	Nomor_Antrian       int                `bson:"nomor_antrian,omitempty" json:"nomor_antrian,omitempty"`
	Tanggal_Pendaftaran primitive.DateTime `bson:"tanggal_pendaftaran,omitempty" json:"tanggal_pendaftaran,omitempty"`
	Status_Antrian      string             `bson:"status_antrian,omitempty" json:"status_antrian,omitempty"`
}

type Poliklinik struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_Poliklinik  string             `bson:"kode_poliklinik,omitempty" json:"kode_poliklinik,omitempty"`
	Nama_Poliklinik  string             `bson:"nama_poliklinik,omitempty" json:"nama_poliklinik,omitempty"`
	Deskripsi        string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Identitas_Dokter Dokter             `bson:"dokter,omitempty" json:"dokter,omitempty"`
}

type Dokter struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Dokter  string             `bson:"nama_dokter,omitempty" json:"nama_dokter,omitempty"`
	Spesialisasi string             `bson:"spesialisasi,omitempty" json:"spesialisasi,omitempty"`
}
