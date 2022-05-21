package request

type Device struct {
	ID                string `json:"id"`
	Mac               string `json:"mac"`
	ConnectProtocol   string `json:"connect_protocol"`
	AuthKey           string `json:"auth_key"`
	CloseStrategy     string `json:"close_strategy"`
	ConnStrategy      string `json:"conn_strategy"`
	CryptMethod       string `json:"crypt_method"`
	AuthVer           string `json:"auth_ver"`
	ManuMacPos        string `json:"manu_mac_pos"`
	SerMacPos         string `json:"ser_mac_pos"`
	BleSimpleProtocol string `json:"ble_simple_protocol"`
}

type RequestDeviceAuthorize struct {
	DeviceNum  string    `json:"device_num"`
	DeviceList []*Device `json:"device_list"`
	OpType     string    `json:"op_type"`
	ProductID  string    `json:"product_id"`
}
