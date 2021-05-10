// Code generated by dbml-gen-go-model. DO NOT EDIT.
// Supported by duythinht@2020
package records

// IbaServer is generated type for table 'iba_servers'
type IbaServer struct {
	ID int64 `db:"id" json:"id" mapstructure:"id"`
	// имя сервера
	Name string `db:"name" json:"name" mapstructure:"name"`
	// цех
	Zex string `db:"zex" json:"zex" mapstructure:"zex"`
	// место устанкови
	Stort string `db:"stort" json:"stort" mapstructure:"stort"`
	// заметки
	Comment string `db:"comment" json:"comment" mapstructure:"comment"`
}

// table 'iba_servers' columns list struct
type __tbl_iba_servers_columns struct {
	ID      string
	Name    string
	Zex     string
	Stort   string
	Comment string
}

// table 'iba_servers' metadata struct
type __tbl_iba_servers struct {
	Name    string
	Columns __tbl_iba_servers_columns
}

// table 'iba_servers' metadata info
var _tbl_iba_servers = __tbl_iba_servers{
	Columns: __tbl_iba_servers_columns{
		Comment: "comment",
		ID:      "id",
		Name:    "name",
		Stort:   "stort",
		Zex:     "zex",
	},
	Name: "iba_servers",
}

// GetColumns return list columns name for table 'iba_servers'
func (*__tbl_iba_servers) GetColumns() []string {
	return []string{"id", "name", "zex", "stort", "comment"}
}

// T return metadata info for table 'iba_servers'
func (*IbaServer) T() *__tbl_iba_servers {
	return &_tbl_iba_servers
}