package main
import (
"net/http"
gin "github.com/gin-gonic/gin"
colors "gopkg.in/go-playground/colors.v1"
"log"
"os"
)

func main() {
r := gin.Default()
f, err := os.OpenFile("color_log_file", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
if err != nil {
    log.Fatalf("error opening file: %v", err)
}
defer f.Close()
log.SetOutput(f)

r.GET("/convert/:hex", func(c *gin.Context) {
userHex := c.Params.ByName("hex")
    hex, err := colors.ParseHEX("#"+ userHex)
    if err != nil {
      log.Println("Hex color queried " + userHex + " Got error " + err.Error())
      c.JSON(http.StatusNotFound,gin.H{"message":"Unable to parse"})
    } else {
    	log.Println("Hex color queried " + userHex + " RGB is " + hex.ToRGB().String())
      c.JSON(http.StatusOK, gin.H{"color":hex.ToRGB().String()})
    }
})
r.Run() // listen and serve on 0.0.0.0:8080
}
