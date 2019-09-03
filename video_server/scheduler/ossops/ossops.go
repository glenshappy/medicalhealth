package ossops

import (
	"avenssi/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	_ "medicalhealth/video_server/scheduler/dbops"
)

var EP string
var AK string
var SK string

func init() {
	AK = ""
	SK = ""
	EP = config.GetOssAddr()
	fmt.Println("init in ossops")
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("Uploading object error: %s", err)
		return false
	}

	return true
}

func DeleteObject(filename string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bueckt error: %s", err)
		return false
	}

	err = bucket.DeleteObject(filename)
	if err != nil {
		log.Printf("Deleting object error: %s", err)
		return false
	}

	return true
}
