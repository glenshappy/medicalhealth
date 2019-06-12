package main

import (
	"medicalhealth/video_server/scheduler/dbops"
	"medicalhealth/video_server/scheduler/taskrunner"
	"testing"
)

func TestMain(m *testing.M)  {

	m.Run()
}

func TestGetDb(t *testing.T) {
	dbops.AddVideoDeletionRecord("sjsjsjs")
	taskrunner.TestPack1()
}
