package main

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message.")
	log.SetPrefix("INFO: ")
	log.Println("This is and info message.")

	//log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a long message with date, time and longfile.")

	infoLogger.Println("This is and info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is and error message")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // each mode is octal (base 8)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime)
	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug message.")
	warnLogger1.Println("This is a warning message.")
	infoLogger1.Println("This is an info mesage.")
	errorLogger1.Println("This is an error.")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

//Go's os package follows traditional Unix file permission conventions,
// so any Linux file mode resource also applies to Go.​

/* IMPORTANTIMPORTANTIMPORTANTIMPORTANTIMPORTANTIMPORTANTIMPORTANTIMPORTANTIMPORTANT
| Octal Code | rwx Symbolic | User | Group | Others | Description                     |
| ---------- | ------------ | ---- | ----- | ------ | ------------------------------- |
| 0700       | rwx------    | rwx  | ---   | ---    | User all, group/others none     |
| 0600       | rw-------    | rw   | ---   | ---    | User read/write only            |
| 0400       | r--------    | r    | ---   | ---    | User read only                  |
| 0200       | -w-------    | w    | ---   | ---    | User write only                 |
| 0100       | --x------    | x    | ---   | ---    | User execute only               |
| 0777       | rwxrwxrwx    | rwx  | rwx   | rwx    | Everybody all                   |
| 0666       | rw-rw-rw-    | rw   | rw    | rw     | Everybody read/write            |
| 0644       | rw-r--r--    | rw   | r     | r      | User read/write, group/others r |
| 0555       | r-xr-xr-x    | r-x  | r-x   | r-x    | Everybody read/execute          |
| 0444       | r--r--r--    | r    | r     | r      | Everybody read                  |
| 0333       | -wx-wx-wx    | wx   | wx    | wx     | Everybody write/execute         |
| 0222       | -w--w--w-    | w    | w     | w      | Everybody write                 |
| 0111       | --x--x--x    | x    | x     | x      | Everybody execute               |
*/
