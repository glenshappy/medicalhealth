package taskrunner

import (
	"errors"
	"fmt"
	"log"
	"medicalhealth/video_server/scheduler/dbops"
	"sync"
)

func init() {
	fmt.Println("taskrunner 包中的第二个init")
}

func TestPack1() {
	fmt.Println("测试包引入")
}

func deleteVideo(vid string) error {
	fmt.Printf("delete success:%s\n", vid)
	return nil
	//ossfn := "videos/" + vid
	//bn := "avenssi-videos2"
	//ok := ossops.DeleteObject(ossfn, bn)
	//
	//if !ok {
	//	log.Printf("Deleting video error, oss operation failed")
	//	return errors.New("Deleting video error")
	//}
	//
	//return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
