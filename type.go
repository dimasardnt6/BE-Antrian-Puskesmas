package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Identitas_Pasien  Pasien             `bson:"identitas_pasien,omitempty" json:"identitas_pasien,omitempty"`
	Nomor_Antrian     string             `bson:"nomor_antrian,omitempty" json:"nomor_antrian,omitempty"`
	Waktu_Pendaftaran string             `bson:"waktu_pendaftaran,omitempty" json:"waktu_pendaftaran,omitempty"`
	Status_Antrian    []string           `bson:"status_antrian,omitempty" json:"status_antrian,omitempty"`
}

type Poliklinik struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Poliklinik string             `bson:"nama_poliklinik,omitempty" json:"nama_poliklinik,omitempty"`
	Deskripsi       string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
}

type Dokter struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Dokter  string             `bson:"nama_dokter,omitempty" json:"nama_dokter,omitempty"`
	Spesialisasi string             `bson:"total_spesialisasi,omitempty" json:"spesialisasi,omitempty"`
}
type Jadwal_Dokter struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Identitas_Dokter Dokter             `bson:"identitas_dokter,omitempty" json:"identitas_dokter,omitempty"`
	Jenis_Poliklinik Poliklinik         `bson:"jenis_poliklinik,omitempty" json:"jenis_poliklinik,omitempty"`
	Hari_Kerja       []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
	Jam_Mulai        string             `bson:"jam_mulai,omitempty" json:"jam_mulai,omitempty"`
	Jam_Selesai      string             `bson:"jam_selesai,omitempty" json:"jam_selesai,omitempty"`
}
