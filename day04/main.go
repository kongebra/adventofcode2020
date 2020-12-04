package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	BirthYear string
	IssueYear string
	ExpirationYear string
	Height string
	HairColor string
	EyeColor string
	PassportID string
	CountryID string
}

func (pp passport) isValidOne() bool {
	if pp.BirthYear == "" {
		return false
	}

	if pp.IssueYear == "" {
		return false
	}

	if pp.ExpirationYear == "" {
		return false
	}

	if pp.Height == "" {
		return false
	}

	if pp.HairColor == "" {
		return false
	}

	if pp.EyeColor == "" {
		return false
	}

	if pp.PassportID == "" {
		return false
	}

	return true
}

func (pp passport) isValidTwo() bool {
	if !pp.isBirthYearValid() {
		return false
	}

	if !pp.isIssueYearValid() {
		return false
	}

	if !pp.isExpirationYearValid() {
		return false
	}

	if !pp.isHeightValid() {
		return false
	}

	if !pp.isHairColorValid() {
		return false
	}

	if !pp.isEyeColorValid() {
		return false
	}

	if !pp.isPassportIDValid() {
		return false
	}

	if !pp.isCountryIDValid() {
		return false
	}

	return true
}

func (pp passport) isBirthYearValid() bool {
	if pp.BirthYear == "" {
		return false
	}

	num, err := strconv.Atoi(pp.BirthYear)
	if err != nil {
		return false
	}

	if num < 1920 || num > 2002 {
		return false
	}

	return true
}

func (pp passport) isIssueYearValid() bool {
	if pp.IssueYear == "" {
		return false
	}

	num, err := strconv.Atoi(pp.IssueYear)
	if err != nil {
		return false
	}

	if num < 2010 {
		return false
	}

	if num > 2020 {
		return false
	}

	return true
}

func (pp passport) isExpirationYearValid() bool {
	if pp.ExpirationYear == "" {
		return false
	}

	num, err := strconv.Atoi(pp.ExpirationYear)
	if err != nil {
		return false
	}

	if num < 2020 || num > 2030 {
		return false
	}

	return true
}

func (pp passport) isHeightValid() bool {
	if pp.Height == "" {
		return false
	}

	// must include 'cm' or 'in'
	if !strings.Contains(pp.Height, "cm") && !strings.Contains(pp.Height, "in") {
		return false
	}

	// could not parse number
	num, err := strconv.Atoi(pp.Height[:len(pp.Height) - 2])
	if err != nil {
		return false
	}

	if strings.Contains(pp.Height, "cm") {
		if num >= 150 && num <= 193 {
			return true
		}
	}

	if strings.Contains(pp.Height, "in") {
		if num >= 59 && num <= 76 {
			return true
		}
	}

	return false
}

func (pp passport) isHairColorValid() bool {
	if pp.HairColor == "" {
		return false
	}

	// must be valid hex value: #000000 - #ffffff
	matched, err := regexp.Match(`^#(?:[0-9a-fA-F]{3}){1,2}$`, []byte(pp.HairColor))
	if err != nil {
		return false
	}

	return matched
}

func (pp passport) isEyeColorValid() bool {
	if pp.EyeColor == "" {
		return false
	}
	
	switch pp.EyeColor {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

func (pp passport) isPassportIDValid() bool {
	if pp.PassportID == "" {
		return false
	}

	if len(pp.PassportID) != 9 {
		return false
	}

	// 9 digit number, including leading zeroes
	_, err := strconv.ParseInt(pp.PassportID, 10, 64)
	if err != nil {
		return false
	}

	return true
}

func (pp passport) isCountryIDValid() bool {
	// ignored

	return true
}

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	// partOne(input)
	partTwo(input)
}

func getInput() ([]passport, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	var section []string
	var result []passport

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			pp := convertToPassport(section)
			result = append(result, pp)

			section = make([]string, 0)
		} else {
			section = append(section, line)
		}
	}

	return result, nil
}

func convertToPassport(section []string) passport {
	var pp passport

	line := strings.Join(section, " ")
	pairs := strings.Split(line, " ")

	for _, pair := range pairs {
		keyValue := strings.Split(pair, ":")
		key := keyValue[0]
		value := keyValue[1]

		switch (strings.ToLower(key)) {
		case "byr":
			pp.BirthYear = value
			break
		case "iyr":
			pp.IssueYear = value
			break
		case "eyr":
			pp.ExpirationYear = value
			break
		case "hgt":
			pp.Height = value
			break
		case "hcl":
			pp.HairColor = value
			break
		case "ecl":
			pp.EyeColor = value
			break
		case "pid":
			pp.PassportID = value
			break
		case "cid":
			pp.CountryID = value
			break
		default:
			fmt.Println("ERROR: Unknown key", key)
			break
		}
	}

	return pp
}

func partOne(input []passport) {
	validCount := 0

	for _, pp := range input {
		if pp.isValidOne() {
			validCount++
		}
	}

	fmt.Println("Valid count:", validCount)
	fmt.Println("Invalid count:", len(input) - validCount)
}

func partTwo(input []passport) {
	validCount := 0

	for _, pp := range input {
		if pp.isValidTwo() {
			validCount++
		}
	}

	fmt.Println("Valid count:", validCount)
	fmt.Println("Invalid count:", len(input) - validCount)
}