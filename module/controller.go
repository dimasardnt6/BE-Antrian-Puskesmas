package module

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/dimasardnt6/BE-Antrian-Puskesmas/model"
	"golang.org/x/crypto/argon2"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func InsertUser(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		// fmt.Printf("InsertOneDoc: %v\n", err)
		return insertedID, fmt.Errorf("kesalahan server")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// Login Function

func GetUserFromEmail(email string, db *mongo.Database, col string) (result model.User, err error) {
	collection := db.Collection(col)
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, fmt.Errorf("email tidak ditemukan")
		}
		return result, fmt.Errorf("kesalahan server")
	}
	return result, nil
}

func SignUp(db *mongo.Database, col string, insertedDoc model.User) (insertedID primitive.ObjectID, err error) {
	if insertedDoc.Fullname == "" || insertedDoc.Email == "" || insertedDoc.Password == "" || insertedDoc.Confirmpassword == "" {
		return insertedID, fmt.Errorf("Data tidak boleh kosong")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return insertedID, fmt.Errorf("email tidak valid")
	}
	if !strings.Contains(insertedDoc.Email, "@gmail.com") {
		return insertedID, fmt.Errorf("email harus menggunakan domain @gmail.com")
	}
	userExists, _ := GetUserFromEmail(insertedDoc.Email, db, col)
	if insertedDoc.Email == userExists.Email {
		return insertedID, fmt.Errorf("email sudah terdaftar")
	}
	if insertedDoc.Confirmpassword != insertedDoc.Password {
		return insertedID, fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return insertedID, fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return insertedID, fmt.Errorf("password terlalu pendek")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), nil, 1, 64*1024, 4, 32)
	insertedDoc.Password = hex.EncodeToString(hashedPassword)
	insertedDoc.Confirmpassword = ""
	return InsertUser(db, col, insertedDoc)
}

func LogIn(db *mongo.Database, col string, insertedDoc model.User) (userName string, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return userName, fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return userName, fmt.Errorf("email tidak valid")
	}
	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db, col)
	if err != nil {
		return
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), nil, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		return userName, fmt.Errorf("password salah")
	}
	return existsDoc.Fullname, nil
}

// func SignUp(db *mongo.Database, col string, insertedDoc model.User) (insertedID primitive.ObjectID, err error) {
// 	if insertedDoc.FirstName == "" || insertedDoc.LastName == "" || insertedDoc.Email == "" || insertedDoc.Password == "" {
// 		return insertedID, fmt.Errorf("Data tidak boleh kosong")
// 	}
// 	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
// 		return insertedID, fmt.Errorf("email tidak valid")
// 	}
// 	userExists, _ := GetUserFromEmail(insertedDoc.Email, db, col)
// 	if insertedDoc.Email == userExists.Email {
// 		return insertedID, fmt.Errorf("email sudah terdaftar")
// 	}
// 	if insertedDoc.Confirmpassword != insertedDoc.Password {
// 		return insertedID, fmt.Errorf("konfirmasi password salah")
// 	}
// 	if strings.Contains(insertedDoc.Password, " ") {
// 		return insertedID, fmt.Errorf("password tidak boleh mengandung spasi")
// 	}
// 	if len(insertedDoc.Password) < 8 {
// 		return insertedID, fmt.Errorf("password terlalu pendek")
// 	}
// 	salt := make([]byte, 16)
// 	_, err = rand.Read(salt)
// 	if err != nil {
// 		return insertedID, fmt.Errorf("kesalahan server")
// 	}
// 	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
// 	insertedDoc.Password = hex.EncodeToString(hashedPassword)
// 	insertedDoc.Salt = hex.EncodeToString(salt)
// 	insertedDoc.Confirmpassword = ""
// 	return InsertUser(db, col, insertedDoc)
// }

// func LogIn(db *mongo.Database, col string, insertedDoc model.User) (userName string, err error) {
// 	if insertedDoc.Email == "" || insertedDoc.Password == "" {
// 		return userName, fmt.Errorf("mohon untuk melengkapi data")
// 	}
// 	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
// 		return userName, fmt.Errorf("email tidak valid")
// 	}
// 	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db, col)
// 	if err != nil {
// 		return
// 	}
// 	salt, err := hex.DecodeString(existsDoc.Salt)
// 	if err != nil {
// 		return userName, fmt.Errorf("kesalahan server")
// 	}
// 	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
// 	if hex.EncodeToString(hash) != existsDoc.Password {
// 		return userName, fmt.Errorf("password salah")
// 	}
// 	return existsDoc.FirstName + " " + existsDoc.LastName, nil
// }

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

func InsertAntrian(db *mongo.Database, col string, poli model.Poliklinik, identitas_pasien model.Pasien, status_antrian string) (insertedID primitive.ObjectID, err error) {
	nomor_antrian, _ := GetAntrianTerakhir(db)
	antrian := bson.M{
		"poli":                poli,
		"identitas_pasien":    identitas_pasien,
		"nomor_antrian":       nomor_antrian + 1,
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

func InsertPoliklinik(db *mongo.Database, col string, kode_poliklinik string, nama_poliklinik string, deskripsi string, dokter model.Dokter) (insertedID primitive.ObjectID, err error) {
	poliklinik := bson.M{
		"kode_poliklinik": kode_poliklinik,
		"nama_poliklinik": nama_poliklinik,
		"deskripsi":       deskripsi,
		"dokter":          dokter,
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

func GetAntrianTerakhir(db *mongo.Database) (int, error) {
	var data model.Antrian
	antrian := db.Collection("data_antrian")
	filter := bson.D{{}}
	opts := options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}})
	err := antrian.FindOne(context.TODO(), filter, opts).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data.Nomor_Antrian, fmt.Errorf("no data found for ID %s", data.ID)
		}
		return data.Nomor_Antrian, fmt.Errorf("error retrieving data for ID %s: %s", data.ID, err.Error())
	}
	return data.Nomor_Antrian, nil
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

func GetUserFromID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.User, errs error) {
	user := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
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

func GetAllUser(db *mongo.Database, col string) (data []model.User) {
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

func UpdatePoliklinik(db *mongo.Database, col string, id primitive.ObjectID, kode_poliklinik string, nama_poliklinik string, deskripsi string, dokter model.Dokter) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kode_poliklinik": kode_poliklinik,
			"nama_poliklinik": nama_poliklinik,
			"deskripsi":       deskripsi,
			"dokter":          dokter,
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

func UpdateUser(db *mongo.Database, col string, id primitive.ObjectID, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		fmt.Printf("UpdateUser: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

// Delete Function

func DeletePasienByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	pasien := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := pasien.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

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

func DeleteDokterByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	dokter := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := dokter.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeleteUserByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	user := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
