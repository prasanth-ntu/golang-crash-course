module 12-packages

go 1.20

replace greetings => ../greetings

replace reverse => ../reverse

require (
	greetings v0.0.0-00010101000000-000000000000
	reverse v0.0.0-00010101000000-000000000000
)
