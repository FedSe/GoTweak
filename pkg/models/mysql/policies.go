package mysql
 
import (
	"database/sql"
	"github.com/FedSe/GoTweak/pkg/models"
)
 
// PolicyModel - Определяем тип который обертывает пул подключения sql.DB
type PolicyModel struct {
	DB *sql.DB
}
 
// Insert - Метод для создания новой заметки в базе дынных. 
//newkey varchar(40), IN newbrowser INT, IN newsub varchar(300), IN newtitle varchar(300)
//IN newdescr varchar(300), IN newcat varchar(300), IN newvers INT, IN newvere
func (m *PolicyModel) Insert(newkey string, newbrowser int) (int, error) {
	stmt := `CALL add_policy(?, ?)`
 
	result, err := m.DB.Exec(stmt, newkey, newbrowser)
	if err != nil {
		return 0, err
	}
 
	// Используем метод LastInsertId(), чтобы получить последний ID
	// созданной записи из таблицу snippets.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
 
	// Возвращаемый ID имеет тип int64, поэтому мы конвертируем его в тип int
	// перед возвратом из метода.
	return int(id), nil
}
 
// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *PolicyModel) Get(id int) (*models.Policy, error) {
	return nil, nil
}
 
// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *PolicyModel) Latest() ([]*models.Policy, error) {
	return nil, nil
}
