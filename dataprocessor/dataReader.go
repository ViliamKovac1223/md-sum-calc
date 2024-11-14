package dataprocessor

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DataReader struct {
    data string
    record SingleFundData
    wasRecordSumPresent bool
    line string
    lineNumber int
}

func (reader *DataReader) ReadFromString(data string) (FundData, error) {
    // List of all records
    fundData := FundData{}
    reader.data = data
    isFirstRecord := true

    lines := strings.Split(data, "\n")
    for i, line := range lines {
        // Update reader's variables
        reader.line = line
        reader.lineNumber = i + 1

        // Skip empty line, or commented line
        if line == "" || reader.isCommentedLine() {
            continue
        }

        // If a line starts with '#' it means that a new SingleFundData record
        // starts here
        if strings.HasPrefix(line, "#") {
            // If this isn't the first record then add previous to the fundData 
            if !isFirstRecord {
                fundData.Sums = append(fundData.Sums, reader.record)
            }
            isFirstRecord = false
            // Reset current record
            reader.record = SingleFundData{}
            // Reset sum record indicator for new record
            reader.wasRecordSumPresent = false

            err := reader.processHeader()
            if (err != nil) {
                return fundData, err
            }

            continue
        }

        // Check if a line is a sum record
        if strings.HasPrefix(line, "- sum:") {
            // If there are two sums in one record then throw error
            if reader.wasRecordSumPresent {
                return fundData, errors.New(fmt.Sprintf(
                    "There can only be one sum under each header; line: %d", i))
            }

            reader.wasRecordSumPresent = true
            err := reader.processSum()
            if (err != nil) {
                return fundData, err
            }
            continue
        }

        // Check if a line is a basic record
        if strings.HasPrefix(line, "- ") { if reader.wasRecordSumPresent {
                return fundData, errors.New(fmt.Sprintf(
                    "There cannot be a new record after sum was defined; line: %d", i))
            }

            err := reader.processRecord()
            if (err != nil) {
                return fundData, err
            }
            continue
        }
    }
    // Add last processed record
    fundData.Sums = append(fundData.Sums, reader.record)

    return fundData, nil
}

func (reader *DataReader) UpdateString(originalString string, fundData FundData) (string, error) {
    numberOfSums := 0
    reader.data = originalString

    lines := strings.Split(originalString, "\n")
    for i, line := range lines {
        // Update reader's variables
        reader.line = line
        reader.lineNumber = i + 1

        if strings.HasPrefix(line, "- sum:") {
            reader.updateSum(fundData.Sums[numberOfSums].Sum)
            numberOfSums++

            // Update line with the new sum
            lines[i] = reader.line
        }
    }

    if numberOfSums != len(fundData.Sums) {
        return "", errors.New("The new string data doesn't have required amount of sums in it")
    }

    // Return new updated string
    newData := strings.Join(lines, "\n")
    return newData[:len(newData)-1], nil
}

func (reader *DataReader) updateSum(newSum float64) {
    // Get old sum from the orignal line
    sumNumber := strings.TrimPrefix(reader.line, "- sum:")
    sumNumber = reader.removeCommentFromEndOfTheLine(sumNumber)
    sumNumber = strings.TrimSpace(sumNumber)
    // Convert newSum to the string
    newSumString := strconv.FormatFloat(newSum, 'f', 2, 64)
    // Replace old sum with the new sum
    reader.line = strings.Replace(reader.line, sumNumber, newSumString, 1)
}

func (reader *DataReader) processHeader() error {
    // Set header for the record
    header := strings.TrimPrefix(reader.line, "#")
    header = strings.TrimSpace(header)
    reader.record.Header = header

    if header == "" {
        return errors.New(fmt.Sprintf("Header must be defined, on line %d", reader.lineNumber))
    }

    return nil
}

func (reader *DataReader) processSum() error {
    // Remove "- sum:" from the string
    sumNumber := strings.TrimPrefix(reader.line, "- sum:")
    // Remove any comment that might be left
    sumNumber = reader.removeCommentFromEndOfTheLine(sumNumber)
    sumNumber = strings.TrimSpace(sumNumber)
    num, err := strconv.ParseFloat(sumNumber, 64)

    if err != nil {
        return errors.New(fmt.Sprintf("Invalid number on line %d", reader.lineNumber))
    }
    reader.record.Sum = num

    return nil
}

func (reader *DataReader) removeCommentFromEndOfTheLine(line string) string {
    re := regexp.MustCompile("<!--.*-->")
    cleanedStr := re.ReplaceAllString(line, "")
    return cleanedStr
}

func (reader *DataReader) processRecord() error {
    // Get number from record
    num, err := reader.extractNumberfomStringToRecord()
    if err != nil {
        return err
    }

    // Get date from record
    dateStr := strings.Split(reader.line, ";")[0]
    dateStr = strings.TrimPrefix(dateStr, "- ")

    // Add a new record to the list
    reader.record.Records = append(reader.record.Records, FundDataRecord{dateStr, num})

    return nil
}

func (reader *DataReader) extractNumberfomStringToRecord() (float64, error) {
    subStrings := strings.Split(reader.line, ";")
    if len(subStrings) != 3 {
        return 0, errors.New(fmt.Sprintf(
            "Record must have exactly three columns; line: %d",
            reader.lineNumber))
    }

    numStr := strings.TrimSpace(subStrings[1])
    // Parse middle value and handle potential error
    num, err := strconv.ParseFloat(numStr, 64)
    if err != nil {
        return 0, errors.New(fmt.Sprintf(
            "Record must have a valid number in 2nd column; line: %d",
            reader.lineNumber))
    }

    return num, nil
}

func (reader *DataReader) isCommentedLine() bool {
    trimmedLine := strings.TrimSpace(reader.line)
    return strings.HasPrefix(trimmedLine, "<!--") && strings.HasSuffix(trimmedLine, "-->")
}
