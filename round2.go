package main

import (
    "bytes"
    "fmt"
    "regexp"
    "strings"
)

const CIPHER = "Lpk peuqyuof bu rlwm immw aapdwvrk. Oj ugkbw vuf. " +
    "Msghq I zij blsvtqd ab gxl stuzg, kchoofaiuomark wsqzuny nud tzm " +
    "xugzb samwvz. Fhw joxl oiy m fskzar, A bnunc jkoamak U hsl ta mgv" +
    "kk tg xgk fgz of. Omz xaoe akdvakk fata nmd tmkz rmvtuny auyeopkd" +
    "e tmziewv ziefbe zifm gzd lpodtq aoj dgtrmrk xkd hgcx roj nudtq m" +
    "oshl kuzswkafinm nauja. Ozcjmjubdm. Naw uwaxd ab nmphmt? Nul je f" +
    "hw boye A iywev bnus icketawt, fhwzk iak vu anw ixaufl za afacqr " +
    "ab. Ztal zuftwv gftgztqy gn sunw, Lx. Sofhu, iak ouze. Zm sgsl pg" +
    "he kmteev bxauttk. Dagcr Pucm - Lqaj itp Lgiztifo oz Lsa Bqgsa"

const KEYWORD = "sigma"

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func main() {

    keystream := buildKeystream(CIPHER, KEYWORD)
    
    plaintext := decipher(CIPHER, keystream)

    fmt.Println(plaintext)
}

func buildKeystream(cipher string, keyword string) string {
    var buffer bytes.Buffer

    sigmaKey := 0
    for _, val := range cipher {
        sigmaKey = sigmaKey % len(keyword)

        matched, _ := regexp.MatchString("[A-Za-z]", string(val))
        if matched {
            buffer.WriteString(string(keyword[sigmaKey]))
        } else {
            buffer.WriteString(string(val))
            sigmaKey--
        }

        sigmaKey++
    }

    return buffer.String()
}

func decipher(cipher string, keystream string) string {
    var buffer bytes.Buffer

    for key, val := range cipher {
        matched, _ := regexp.MatchString("[A-Za-z]", string(val))
        if matched {
            sigma := keystream[key]
            // Get the index in the alphabet of the keystream character
            // to know how far we should shift the letter
            shift := bytes.IndexRune([]byte(ALPHABET), rune(sigma))
            c := shiftLetter(val, shift)
            if c > 'Z' || c < 'A' {
                fmt.Printf("%c %c\n", val, c)
            }
            buffer.WriteString(string(c))
        } else {
            buffer.WriteString(string(val))
        }
    }

    return buffer.String()
}

func shiftLetter(letter rune, shift int) rune {
    // The letter is decoded to 0-25 (Capital A=65)
    // The shift is added and the modolu operator makes sure we
    // start over at A after Z. At the end, 65 is added to get
    // back to the decimal value of the letter
    c := rune(strings.ToUpper(string(letter))[0]) // To uppercase

    // Some sort of negative shift...
    for i := 0; i < shift; i++ {
        if c == 'A' {
            c = 'Z'
        } else {
            c--
        }
    }

    return rune(c)
}