package prisma

import "github.com/joho/godotenv"

func LoadConfig() error {
	if err := godotenv.Load("../../internal/prisma/.env"); err != nil {
		return err
	}

	return nil
}
