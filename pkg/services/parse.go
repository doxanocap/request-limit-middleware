package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"test-task-rlm/pkg/rlm/models"
)

func ParseParams() {
	if len(os.Args) == 5 && os.Args[1] == "-i" {
		m, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		b, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		bti, err := strconv.Atoi(os.Args[4])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		models.DefaultParams = models.Params{MaxRate: m, BlockTime: int64(b), BTIncrement: bti}

		fmt.Printf(" # --- Default Params Were Initialized --- # \n")
	}
}
