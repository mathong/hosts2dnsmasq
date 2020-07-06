package tests

var (
	// Hosts is a hosts list example
	Hosts = []string{"testD0main.com", "an0thertestdomain.com"}
	// Config is the dnsmasq config file that should be generated with Hosts
	Config = "address=/testD0main.com/0.0.0.0\naddress=/an0thertestdomain.com/0.0.0.0\n"
)
