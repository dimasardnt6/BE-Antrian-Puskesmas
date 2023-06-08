package namapackage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertAntrian(identitas_pasien Pasien, nomor_antrian string, waktu_pendaftaran string, status_antrian []string) (InsertedID interface{}) {
	var antrian Antrian
	antrian.Identitas_Pasien = identitas_pasien
	antrian.Nomor_Antrian = nomor_antrian
	antrian.Waktu_Pendaftaran = waktu_pendaftaran
	antrian.Status_Antrian = status_antrian
	return InsertOneDoc("antrian_puskesmas", "data_antrian", antrian)
}

func GetAntrianFromNomorAntrian(nomor_antrian string) (data Pasien) {
	antrian := MongoConnect("antrian_puskesmas").Collection("data_antrian")
	filter := bson.M{"nomor_antrian": nomor_antrian}
	err := antrian.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getAntrianFromNomorAntrian: %v\n", err)
	}
	return data
}
