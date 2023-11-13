package gobackendgis5

import (
	"context"
	"fmt"
	pasproj "github.com/e-dumas-sukasari/webpasetobackend"
	"github.com/whatsauth/watoken"
	"testing"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	//privateKey, publicKey := watoken.GenerateKey()
	//fmt.Println(privateKey)
	//fmt.Println(publicKey)
	hasil, err := watoken.Encode("edumas1", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "Ini Polygon",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGO", "GIS", context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "Ini Polygon",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGO", "GIS", context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	conn := GetConnectionMongo("MONGO", "GIS")
	pass, _ := pasproj.HashPass("4321234")
	data := RegisterStruct{
		Username: "edumas1",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}