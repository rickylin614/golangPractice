package main

import (
	"fmt"
	"practice/demo/payment/payment/validation"
)

var (
	inputString string = `eyJTYyI6IjMiLCJNZXJObyI6IjY3ODkwMDAxIiwiZGF0YSI6IkpUZENKVEl5VDNKa1pYSk9ieVV5TWlVelFTVXlNa05OTWxoQ1VVOHhUVlJOTnkxU1dTMVNSMHdsTWpJbE1rTWxNakpCYlc5MWJuUWxNaklsTTBFbE1qSXlMakF3SlRJeUpUSkRKVEl5UTNWeUpUSXlKVE5CSlRJeU1pNHdNQ1V5TWlVeVF5VXlNbE4wWVhSMWN5VXlNaVV6UVNVeU1qRXdNQ1V5TWlVeVF5VXlNbEJoZVZScGJXVWxNaklsTTBFbE1qSXlNREl5TFRBNExURXhNVElsTTBFME55VXpRVE00SlRJeUpUSkRKVEl5UkdWaGJGUnBiV1VsTWpJbE0wRWxNakl5TURJeUxUQTRMVEV4TVRJbE0wRTBOeVV6UVRNNEpUSXlKVEpESlRJeVEyOWtaVTF6WnlVeU1pVXpRU1V5TWlVMVEzVTJOVEptSlRWRGRUUmxaRGdsTlVOMU5qSXhNQ1UxUTNVMU1qbG1KVEl5SlRKREpUSXlVbUZ1Wkc5dEpUSXlKVE5CSlRJeWJGTmpiRlZOSlRJeUpUSkRKVEl5VTJsbmJpVXlNaVV6UVNVeU1tTTJNamxpTWpBd1pEQTVaak0wTWpReE56TXlNalJqWTJRd01EZGxaRGc0SlRJeUpUZEUifQ==`
)

func main() {
	log := &validation.LogPayValidate{}
	str, err := validation.ParsePayValidateResponse(inputString, log)
	if err != nil {
		fmt.Println("err!! ", err)
	} else {
		fmt.Println("str!!", str)
	}
}
