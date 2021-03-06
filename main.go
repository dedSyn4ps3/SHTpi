package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"time"
	"encoding/json"

	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-sht3x"
)

func sendTemp(t string) {

	client := &http.Client{}

	body := map[string]string{"key" : "SHTpi-Temp", "value" : t}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "https://groker.init.st/api/events", bytes.NewBuffer(jsonBody))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-IS-AccessKey", "YOUR_ACCESS_KEY")
	req.Header.Add("X-IS-BucketKey", "YOUR_BUCKET_KEY")
	req.Header.Add("Accept-Version", "~0")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

}

func sendHum(h string) {

        client := &http.Client{}

        body := map[string]string{"key" : "SHTpi-Humidity", "value" : h}
        jsonBody, _ := json.Marshal(body)

        req, _ := http.NewRequest("POST", "https://groker.init.st/api/events", bytes.NewBuffer(jsonBody))

        req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-IS-AccessKey", "YOUR_ACCESS_KEY")
	req.Header.Add("X-IS-BucketKey", "YOUR_BUCKET_KEY")
        req.Header.Add("Accept-Version", "~0")

        resp, err := client.Do(req)

        if err != nil {
                fmt.Println("Errored when sending request to the server")
                return
        }

        defer resp.Body.Close()
        resp_body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(resp.Status)
        fmt.Println(string(resp_body))

}

func main() {

	i2c, err := i2c.NewI2C(0x44, 1) //Cleaner declaration for setting up SHT sensor instead of using flags.
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer i2c.Close()

	sensor := sht3x.NewSHT3X()



	for {

		temp, rh, err := sensor.ReadTemperatureAndRelativeHumidity(i2c, sht3x.RepeatabilityLow)

		if err != nil {
			log.Fatal(err)
		}

		t := fmt.Sprintf("%.2f", float32(temp) * 1.8 + 19.8) //This seems to be a better value offset for accurate Fahrenheit readings
    		h := fmt.Sprintf("%.1f", float32(rh))

		go sendTemp(t)      //Using go routine prevents code interupts for the next function call.

		sendHum(h)

		time.Sleep(20 * time.Second)

	}

}
