package main

import (
    "fmt"
    "regexp"
)

// CAESAR CIPHER
const CIPHER = "BPM ZMJMT CVLWCJBMLTG LMUIVLA I KMZBIQV LMOZMM WN NZMMLWU NWZ PQUAMTN; " +
"JCB QV VW KIAM, QN PM QA KWVAQABMVB, LWMA PM LMUIVL BPM ZQOPB BW LMABZWG BPM MFQABMVKM " +
"IVL BPM NZMMLWU WN WBPMZA. PM PCUQTQIBMA VW WVM. BPM NZMMLWU PM KTIQUA, PM KTIQUA NWZ ITT; " +
"BPM NZMMLWU PM ZMNCAMA, PM NWZJQLA MDMZGWVM BW MVRWG. PM QA VWB WVTG BPM ATIDM IOIQVAB BPM " +
"UIABMZ, JCB ITAW UIV IOIQVAB BPM EWZTL WN UIABMZ IVL ATIDM. BPMZMNWZM, BPIVSA BW ZMJMTTQWV, " +
"BPMZM QA AWUMBPQVO UWZM QV PQABWZG BPIV BPM ZMTIBQWV JMBEMMV UIABMZG IVL AMZDQBCLM. CVTQUQBML " +
"XWEMZ QA VWB BPM WVTG TIE. QB QA QV BPM VIUM WN IVWBPMZ DITCM BPIB BPM ZMJMT INNQZUA BPM " +
"QUXWAAQJQTQBG WN BWBIT NZMMLWU EPQTM PM KTIQUA NWZ PQUAMTN BPM ZMTIBQDM NZMMLWU VMKMAAIZG BW " +
"ZMKWOVQHM BPQA QUXWAAQJQTQBG. BPM ZMJMT ITJMZB KIUCA"

func main() {
    // First, find out how many letters the text is shifted
    // This is done shifting the first word 1-26 times
    // to see what numbers are worth testing.

    findShiftValue()

    // The shift of 18 gives "BPM" -> "THE" which
    // seems reasonable

    shift := 18

    for _, val := range CIPHER {
        matched, _ := regexp.MatchString("[A-Z]", string(val))
        if (matched) {
            fmt.Printf("%c", shiftLetter(val, shift))
        } else {
            fmt.Printf("%c", val)
        }
    }

}

func findShiftValue() {
    word := "BPM"
    for i := 1; i < 26; i++ {
        fmt.Printf("%d: ", i)
        for _, val := range word {
            fmt.Printf("%c", shiftLetter(val, i))
        }
        fmt.Println()
    }
}

func shiftLetter(letter rune, shift int) rune {
    // The letter is decoded to 0-25 (Capital A=65)
    // The shift is added and the modolu operator makes sure we
    // start over at A after Z. At the end, 65 is added to get
    // back to the decimal value of the letter
    c := (((int(letter) - 65) + shift) % 26) + 65
    return rune(c)
}