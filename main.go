package main

import (
	"time"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)
//Student 结构体
type Student struct{
	Name string
	Age int
}

//初始化获取表
func mongoInit()*mongo.Collection{
	//设置参数
	mongoOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), mongoOptions)//*mongo.Client 
	if err != nil {
		fmt.Println("连接错误", err)
	}
	student:= client.Database("student").Collection("user")//*mongo.Collection
	return student
}

//ConnectDb 初始化连接池
func ConnectDb(uri,name string,num uint64,time time.Duration)(*mongo.Database,error){
	o:=options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	ctx,cancel:= context.WithTimeout(context.Background(),time)
	defer cancel()
	client,err:=mongo.Connect(ctx,o)
	if err != nil{
		return nil,err
	}
	return client.Database(name),nil
}

func bsonInit(){
	var m = bson.M{
		"name":bson.M{"$age":"tom"},
		"dd":bson.M{"$eq":"tom"},
		"df":1,
		"ff":bson.A{"1","2"},
	}
	fmt.Println(m);

	var d = bson.D{
		{
			"name",
			bson.D{{
				"$in",
				bson.A{"1","2"},
			}},
		},
		{
			"ff",
			bson.D{{
				"$in",
				bson.A{"1","2"},
			}},
		}}
	fmt.Println(d);
	
	filter := bson.D{{
		"name", 
		"Speike",
		}}	
	fmt.Printf("%#v\n",filter[0]);
	update:= bson.D{
		{
			"$push",
			bson.D{{"test","sing"},},
		},
	}
	fmt.Println(update);
}


func insert(con *mongo.Database){
	S1:=&Student{
		"ssd",
		10,
	}
	students:=con.Collection("student")
	fmt.Println("表",students);
	ret,err:=students.InsertOne(context.TODO(),S1)
	if err != nil{
		fmt.Println("插入错误",err)
	}
	fmt.Println(ret);
}

func update(con *mongo.Database){
	filter:=bson.D{
		{"name","hdd"},
	}
	update:=bson.D{
		{"$set",bson.D{{"age",10}}},	
	}

	students:=con.Collection("student")
	fmt.Println("表",students);
	ret,err:=students.UpdateOne(context.TODO(),filter,update)
	if err != nil{
		fmt.Println("插入错误",err)
	}
	fmt.Println(ret);
}

func find(con *mongo.Database){
	var result Student
	students:=con.Collection("student")
	ret:=students.FindOne(context.TODO(),bson.D{{"name","fhd"}}).Decode(&result)

	fmt.Println(ret);
	fmt.Println(result);
}

func main() {
	//bsonInit()

	
	uri:="mongodb://localhost:27017"
	con,err:= ConnectDb(uri,"student",50,time.Second*3)
	if err != nil{
		fmt.Println("连接错误",err)
		return
	}
	//insert(con)
	
	update(con)

	//find(con)
}
