package script

import "os"

func GetServerConfigFilePath() string {
	return "./config.json"
}

func GetServerWhiteListFilePath() string {
	return "./white_list.json"
}

func GetServerLedgerConfigFilePath() string {
	return GetServerConfig().DataPath + "/ledger_config.json"
}

func GetTemplateLedgerConfigDirPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currentPath + "/template"
}

func GetLedgerTransactionsTemplateFilePath(dataPath string) string {
	return dataPath + "/.beancount-ns/transaction_template.json"
}

func GetLedgerAccountTypeFilePath(dataPath string) string {
	return dataPath + "/.beancount-ns/account_type.json"
}

func GetLedgerPriceFilePath(dataPath string) string {
	return dataPath + "/price/prices.bean"
}

func GetLedgerMonthsFilePath(dataPath string) string {
	return dataPath + "/month/months.bean"
}
