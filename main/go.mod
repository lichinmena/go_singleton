module example/main

go 1.17

replace example/clientone => ../clientone

replace example/clienttwo => ../clienttwo

replace example/singleton => ../singleton

require (
	example/clientone v0.0.0-00010101000000-000000000000
	example/clienttwo v0.0.0-00010101000000-000000000000
	example/singleton v0.0.0-00010101000000-000000000000
)
