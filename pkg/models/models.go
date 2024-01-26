package models
 
import (
	"errors"
)
 
var ErrNoRecord = errors.New("models: подходящей записи не найдено")
 
type Policy struct {
	Id_pol			int
	Id_browser		int
	Idnk_title		int
	Idnk_description	int
	Idnk_category		int
	Key			string
	Subkey			string
	Verstart		int
	Verend			int
}
