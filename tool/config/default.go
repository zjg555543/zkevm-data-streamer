package config

// DefaultValues is the default configuration
const DefaultValues = `
[StreamServer]
Port = 8080
Filename = "datastreamer.bin"

[StateDB]
User = "state_user"
Password = "state_password"
Name = "state_db"
Host = "localhost"
Port = "5432"
EnableLog = false	
MaxConns = 200
`
