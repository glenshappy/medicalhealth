package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"os"
	"time"
)

const VIDEO_DIR = "./videos"

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("./templates/test.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func uploadVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func getVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	videoDir := VIDEO_DIR + "/" + vid + ".mp4"
	file, err := os.Open(videoDir)
	if err != nil {
		sendErroMsg(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), file)
}
