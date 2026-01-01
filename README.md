# taxi-price

This is my first Go project - simplified version of the Yandex taxi price calculator - nothing fancy here, just a little practice for me

# Running the script

Just open **main.exe** file or start the script in IDE with command

```bash
go run .
```

# Project structure

```bash
taxi-price/
├── cmd/             
│   └── app/
│       └── main.go                 # Main file to run
├── internal/                       # Internal packages
│   ├── calculator/                 # Calculator package
│   │   ├── basePrice.go
│   │   └── totalPrice.go
│   └── multipliers                 # Multipliers package
│       ├── timeMultiplier.go       
│       ├── trafficMultiplier.go
│       └── weatherMultiplier.go
├── main.exe                        # Script to start
└── go.mod                  
```
