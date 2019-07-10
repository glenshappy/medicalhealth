package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

const (
	 width= 600
	 height=320
	 cells=100
	xyrange = 30.0
	angle=math.Pi/9
	xyscale=width/2/xyrange
	zscale=height*0.4
)
var sin30,cos30=math.Sin(angle),math.Cos(angle)
func f(x,y float64) float64 {
	r:=math.Hypot(x,y) //距离(0,0)的距离
	return math.Sin(r)/r
}

func corner(i,j int) (float64,float64) {
	x:=xyrange*(float64(i)/cells-0.5)
	y:=xyrange*(float64(j)/cells-0.5)
	z:=f(x,y)
	sx:= width/2+(x-y)*cos30*xyscale
	sy:= height/2+(x+y)*sin30*xyscale-z*zscale
	return sx,sy
}

func GetImg(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	var repStr string
	repStr=`<?xml version="1.0" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" 
"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"> `
	repStr += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey;fill:white;stroke-with:0.7' "+
		"width='%d' height='%d'"+">",width,height)
	for i:=0;i<cells; i++ {
		for j:=0;j<cells ;j++  {
			ax,ay:= corner(i+1,j)
			bx,by:= corner(i,j)
			cx,cy:= corner(i,j+1)
			dx,dy:= corner(i+1,j+1)
			repStr+=fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",ax,ay,bx,by,cx,cy,dx,dy)
		}
	}
	repStr+="</svg>"
	io.WriteString(w,repStr)
}

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.GET("/get_img",GetImg)
	return router
}

func main()  {
	var a  bool
	a=false
	fmt.Println(a)
	router:=RegisterHandlers()
	http.ListenAndServe("127.0.0.1:9292", router)
}
