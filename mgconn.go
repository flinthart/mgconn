package mgconn

import (
	"context"

	"github.com/magiconair/properties"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Credentials struct {
	userName string
	password string
	authdb   string
	address  string
	port     string
}

func InitMongodbWithFile(f string, prefix string) (*mongo.Client, error) {

	props := properties.MustLoadFile(f, properties.UTF8)
	return InitMongodbWithProperties(props,prefix)

}

func InitMongodbWithProperties(p *properties.Properties, prefix string) (*mongo.Client, error) {

	c := Credentials{}
	c.userName = p.GetString(prefix+"_MDB_USER", "")
	c.password = p.GetString(prefix+"_MDB_PSWD", "")
	c.authdb = p.GetString(prefix+"_MDB_DB", "")
	c.address = p.GetString(prefix+"_MDB_ADDR", "")
	c.port = p.GetString(prefix+"_MDB_PORT", "")

	return InitMongodb(c)

}

func InitMongodb(c Credentials) (*mongo.Client, error) {

	connectionString := "mongodb://" + c.userName + ":" + c.password + "@" + c.address + ":" + c.port + "/?authSource=" + c.authdb
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err

}


