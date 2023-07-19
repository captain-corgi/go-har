package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var (
		inputPath  = "./in"
		outputPath = "./out"
		resultList []Result
	)

	// Read files
	files, err := os.ReadDir(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		fileName := file.Name()
		if filepath.Ext(fileName) == ".har" {
			f, err := os.Open(filepath.Join(inputPath, file.Name()))
			if err != nil {
				fmt.Println(err)
				return
			}

			var data HAR
			decoder := json.NewDecoder(f)
			err = decoder.Decode(&data)
			if err != nil {
				fmt.Println(err)
				return
			}

			listAPI := make([]API, 0)
			for _, entry := range data.Log.Entries {
				listAPI = append(listAPI, API{
					Url:             entry.Request.Url,
					BodySize:        entry.Response.BodySize,
					StartedDateTime: entry.StartedDateTime,
					EndedDateTime:   entry.StartedDateTime.Add(time.Duration(entry.Time) * time.Millisecond),
					Time:            entry.Time,
					Timings:         entry.Timings,
				})
			}

			var (
				minStartTime = listAPI[0].StartedDateTime
				maxEndTime   = listAPI[0].EndedDateTime
			)
			for i := 1; i < len(listAPI); i++ {
				if listAPI[i].StartedDateTime.Before(minStartTime) {
					minStartTime = listAPI[i].StartedDateTime
				}
				if listAPI[i].EndedDateTime.After(maxEndTime) {
					maxEndTime = listAPI[i].EndedDateTime
				}
			}
			totalDuration := maxEndTime.Sub(minStartTime).Milliseconds()
			fmt.Printf("%s\n", fileName)
			fmt.Printf("- Total Duration: %dms\n", totalDuration)
			fmt.Println()
			for _, api := range listAPI {
				fmt.Printf("%s\n", api.Url)
				fmt.Printf("- Start: %s\n", api.StartedDateTime.Format("2006-01-02T15:04:05Z07:00.000"))
				fmt.Printf("- End: %s\n", api.EndedDateTime.Format("2006-01-02T15:04:05Z07:00.000"))
				fmt.Printf("- Duration: %.0fms\n", api.Time)
				fmt.Println()
			}
			fmt.Println("=========================================")

			result := Result{
				FileName:     fileName,
				API:          listAPI,
				TotalRunTime: float64(totalDuration),
			}
			resultList = append(resultList, result)
		}
	}

	// Write result
	outputFileName := fmt.Sprintf("%s/%s.json", outputPath, time.Now().String())
	f, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(resultList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
