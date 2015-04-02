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
            //shift := strings.Index(string(rune(sigma)), ALPHABET)
            shift := bytes.IndexRune([]byte(ALPHABET), rune(sigma))
            c := shiftLetter(val, shift)
            if c > 90 || c < 65 {
                fmt.Printf("%c %c ", val, c)
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
    letter = rune(strings.ToUpper(string(letter))[0]) // To uppercase
    c := (((int(letter) - 65) + shift) % 26) + 65
    return rune(c)
}