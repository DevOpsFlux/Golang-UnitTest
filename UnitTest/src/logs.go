package main
 
import (
    "log"
    "os"
)
 
var myLogger *log.Logger
 
func main() {

    // 로그파일 오픈
    fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
	}

	defer fpLog.Close()
	
	// 표준로거를 파일로그로 변경
	log.SetOutput(fpLog)

    //defer fpLog.Close()

    //myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
 
    //....
    run()
 
    myLogger.Println("End of Program")
}
 
func run() {
    myLogger.Print("Test")
}