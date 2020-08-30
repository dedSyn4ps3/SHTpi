# SHTpi
Grab sensor values and stream them to your Initial State dashboard for logging and displaying.
================================

Written in the Go programming language, this program is meant to retrieve values from a SHT31 Temperature & Humidity sensor. The sensor is soldered onto a small factor proto-HAT that is mounted on top of a Raspberry Pi Zero-W. Any Pi version can be used, as well as any size proto board, or even none if you prefer.

After coming accross the excellent work of @d2r2, I took full advantage of incorporating his Go libraries to be able to communicate with the sht31 via I2C on the Pi!
Unlike some of the sparse code on Github related to what I wanted to do, I figured it would make more sense to write my own small program to regularly retrieve values by utilizing goroutines, and then sending that data via HTTPS and JSON objects to my Initial State dashboard.

The only adjustments necessary is to provide your own ACCESS_KEY and BUCKET_KET in the code, and then compile right there on your Pi. Make sure before doing that, you have at least go 1.14 and run the 'go get' command for the following modules:
        
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-sht3x"

After that, compile and run and you should be able to see the new feeds show up to add to your dashboard!

***Note: After compiling, it's best to run the program as follows to prevent debug info from populating the terminal. There is an additional library that can be used do deal with this (see d2r2's github), but this was the easiest and quickist solution for me:

	./main > /dev/null &

This will run the program in the background with no debug info displayed to the screen, allowing you to continue using the terminal***
