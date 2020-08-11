package config

type UserAPI struct {
	Matrix *Global `json:"-"`

	Listen          Address         `json:"Listen" comment:"Listen address for this component."`
	Bind            Address         `json:"Bind" comment:"Bind address for this component."`
	AccountDatabase DatabaseOptions `json:"AccountDatabase" comment:"Database configuration for the account database."`
	DeviceDatabase  DatabaseOptions `json:"DeviceDatabase" comment:"Database configuration for the device database."`
}

func (c *UserAPI) Defaults() {
	c.Listen = "localhost:7781"
	c.Bind = "localhost:7781"
	c.AccountDatabase.Defaults()
	c.DeviceDatabase.Defaults()
	c.AccountDatabase.ConnectionString = "file:userapi_accounts.db"
	c.DeviceDatabase.ConnectionString = "file:userapi_devices.db"
}

func (c *UserAPI) Verify(configErrs *ConfigErrors, isMonolith bool) {
	checkNotEmpty(configErrs, "UserAPI.Listen", string(c.Listen))
	checkNotEmpty(configErrs, "UserAPI.Bind", string(c.Bind))
	checkNotEmpty(configErrs, "UserAPI.AccountDatabase.ConnectionString", string(c.AccountDatabase.ConnectionString))
	checkNotEmpty(configErrs, "UserAPI.DeviceDatabase.ConnectionString", string(c.DeviceDatabase.ConnectionString))
}
