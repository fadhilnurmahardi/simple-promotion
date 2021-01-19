package config

import "os"

var config *Config

type Config struct {
	BonusItemEligible      string
	BonusItemSKU           string
	BonusItemMinBuy        string
	PayTwoForThreeEligible string
	PayTwoForThreeMinBuy   string
	PercentageMinBuy       string
	Percentage             string
	PercentageEligible     string
	HTTPAddress            string
}

func GetConfig() *Config {
	if config != nil {
		return config
	}
	return &Config{
		BonusItemEligible:      getEnvOrDefault("BONUS_ITEM_ELIGIBLE", "43N23P"),
		BonusItemSKU:           getEnvOrDefault("BONUS_ITEM_SKU", "234234"),
		BonusItemMinBuy:        getEnvOrDefault("BONUS_ITEM_MIN_BUY", "1"),
		PayTwoForThreeEligible: getEnvOrDefault("PAY_TWO_F_THREE_ELIGIBLE", "120P90"),
		PayTwoForThreeMinBuy:   getEnvOrDefault("PAY_TWO_F_THREE_MIN_BUY", "3"),
		PercentageMinBuy:       getEnvOrDefault("PERCENTAGE_MIN_BUY", "3"),
		Percentage:             getEnvOrDefault("PERCENTAGE", "10"),
		PercentageEligible:     getEnvOrDefault("PERCENTAGE_ELIGIBLE", "A304SD"),
		HTTPAddress:            getEnvOrDefault("HTTP_ADDRESS", ":8080"),
	}
}

func getEnvOrDefault(env string, defaultVal string) string {
	e := os.Getenv(env)
	if e == "" {
		return defaultVal
	}
	return e
}
