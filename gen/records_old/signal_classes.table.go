// Code generated by dbml-gen-go-model. DO NOT EDIT.
// Supported by duythinht@2020
package records

// SignalClass is generated type for table 'signal_classes'
type SignalClass struct {
	ID int64 `db:"id" json:"id" mapstructure:"id"`
	// analog, discrete, virtual
	Name string `db:"name" json:"name" mapstructure:"name"`
}

// table 'signal_classes' columns list struct
type __tbl_signal_classes_columns struct {
	ID   string
	Name string
}

// table 'signal_classes' metadata struct
type __tbl_signal_classes struct {
	Name    string
	Columns __tbl_signal_classes_columns
}

// table 'signal_classes' metadata info
var _tbl_signal_classes = __tbl_signal_classes{
	Columns: __tbl_signal_classes_columns{
		ID:   "id",
		Name: "name",
	},
	Name: "signal_classes",
}

// GetColumns return list columns name for table 'signal_classes'
func (*__tbl_signal_classes) GetColumns() []string {
	return []string{"id", "name"}
}

// T return metadata info for table 'signal_classes'
func (*SignalClass) T() *__tbl_signal_classes {
	return &_tbl_signal_classes
}
