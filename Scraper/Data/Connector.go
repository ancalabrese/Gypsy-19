package Data

import "github.com/ancalabrese/Gypsy-19/Scraper/Data/Country"

type DbCLient interface {
	UpdateDB(Country.Lists) error
	GetLists() (Country.Lists, error)
}
