// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"

)

var (
	ConfigFile	= "config.json"
	Address		= ""
	Port		= ""
	Dir 		= ""	
	c 		Config
)

// struct for loading configuration information
type Config struct {

    Address             string
    Port	        string
    Dir 		string
    Tmpfile		string
}

func init() {
 // open our config and parse
   confBytes, e := ioutil.ReadFile(ConfigFile)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(1)
    }

    if e := json.Unmarshal(confBytes, &c); e != nil {
        fmt.Println(e.Error())
        os.Exit(2)
    }

}

func main() {
	listen := c.Address + ":" + c.Port
	log.Printf("listening on %q...", listen)
	err := http.ListenAndServe(listen, http.FileServer(http.Dir(c.Dir)))
	log.Fatalln(err)
}
