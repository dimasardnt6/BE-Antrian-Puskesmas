package module

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dimasardnt6/BE-Antrian-Puskesmas/model"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "antrian_puskesmas",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// Insert Function
func InsertPasien(db *mongo.Database, col string, nama_pasien string, nomor_ktp string, alamat string, nomor_telepon string, tanggal_lahir string, jenis_kelamin string) (insertedID primitive.ObjectID, err error) {
	pasien := bson.M{
		"nama_pasien":   nama_pasien,
		"nomor_ktp":     nomor_ktp,
		"alamat":        alamat,
		"nomor_telepon": nomor_telepon,
		"tanggal_lahir": tanggal_lahir,
		"jenis_kelamin": jenis_kelamin,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), pasien)
	if err != nil {
		fmt.Printf("InsertPasien: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertAntrian(db *mongo.Database, col string, poli model.Poliklinik, identitas_pasien model.Pasien, nomor_antrian int, status_antrian string) (insertedID primitive.ObjectID, err error) {
	antrian := bson.M{
		"poli":                poli,
		"identitas_pasien":    identitas_pasien,
		"nomor_antrian":       nomor_antrian,
		"tanggal_pendaftaran": primitive.NewDateTimeFromTime(time.Now().UTC()),
		"status_antrian":      status_antrian,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), antrian)
	if err != nil {
		fmt.Printf("InsertAntrian: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertPoliklinik(db *mongo.Database, col string, kode_poliklinik string, nama_poliklinik string, deskripsi string) (insertedID primitive.ObjectID, err error) {
	poliklinik := bson.M{
		"kode_poliklinik": kode_poliklinik,
		"nama_poliklinik": nama_poliklinik,
		"deskripsi":       deskripsi,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), poliklinik)
	if err != nil {
		fmt.Printf("InsertPoliklinik: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertDokter(db *mongo.Database, col string, nama_dokter string, spesialisasi string) (insertedID primitive.ObjectID, err error) {
	dokter := bson.M{
		"nama_dokter":  nama_dokter,
		"spesialisasi": spesialisasi,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), dokter)
	if err != nil {
		fmt.Printf("InsertDokter: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// Get Function
func GetPasienFromID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.Pasien, errs error) {
	pasien := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := pasien.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data, fmt.Errorf("no data found for ID %s", _id)
		}
		return data, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return data, nil
}

func GetAntrianFromID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.Antrian, errs error) {
	antrian := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := antrian.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data, fmt.Errorf("no data found for ID %s", _id)
		}
		return data, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return data, nil
}

func GetPoliklinikFromID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.Poliklinik, errs error) {
	poliklinik := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := poliklinik.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data, fmt.Errorf("no data found for ID %s", _id)
		}
		return data, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return data, nil
}

func GetDokterFromID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.Dokter, errs error) {
	dokter := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := dokter.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data, fmt.Errorf("no data found for ID %s", _id)
		}
		return data, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return data, nil
}

// Get All Function
func GetAllPasien(db *mongo.Database, col string) (data []model.Pasien) {
	pasien := db.Collection(col)
	filter := bson.M{}
	cursor, err := pasien.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPasien :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllAntrian(db *mongo.Database, col string) (data []model.Antrian) {
	antrian := db.Collection(col)
	filter := bson.M{}
	cursor, err := antrian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllAntrian :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllPoliklinik(db *mongo.Database, col string) (data []model.Poliklinik) {
	poliklinik := db.Collection(col)
	filter := bson.M{}
	cursor, err := poliklinik.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPoliklinik :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllDokter(db *mongo.Database, col string) (data []model.Dokter) {
	dokter := db.Collection(col)
	filter := bson.M{}
	cursor, err := dokter.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllDokter :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// Update Function

func UpdatePasien(db *mongo.Database, col string, id primitive.ObjectID, nama_pasien string, nomor_ktp string, alamat string, nomor_telepon string, tanggal_lahir string, jenis_kelamin string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_pasien":   nama_pasien,
			"nomor_ktp":     nomor_ktp,
			"alamat":        alamat,
			"nomor_telepon": nomor_telepon,
			"tanggal_lahir": tanggal_lahir,
			"jenis_kelamin": jenis_kelamin,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePasien: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateAntrian(db *mongo.Database, col string, id primitive.ObjectID, poli model.Poliklinik, identitas_pasien model.Pasien, nomor_antrian int, status_antrian string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"poli":                poli,
			"identitas_pasien":    identitas_pasien,
			"nomor_antrian":       nomor_antrian,
			"tanggal_pendaftaran": primitive.NewDateTimeFromTime(time.Now().UTC()),
			"status_antrian":      status_antrian,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateAntrian: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdatePoliklinik(db *mongo.Database, col string, id primitive.ObjectID, kode_poliklinik string, nama_poliklinik string, deskripsi string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kode_poliklinik": kode_poliklinik,
			"nama_poliklinik": nama_poliklinik,
			"deskripsi":       deskripsi,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePoliklinik: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateDokter(db *mongo.Database, col string, id primitive.ObjectID, nama_dokter string, spesialisasi string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_dokter":  nama_dokter,
			"spesialisasi": spesialisasi,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateDokter: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

// Delete Function

func DeleteAntrianByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	antrian := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := antrian.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeletePoliklinikByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	poliklinik := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := poliklinik.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
