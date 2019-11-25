package rpc
type Xtype struct {
	Name string `json:"name"`
}

func (x Xtype)Hello()  string{
	return x.Name+"hello"
}