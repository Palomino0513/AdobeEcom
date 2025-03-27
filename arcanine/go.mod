module arcanine

go 1.23.0

toolchain go1.24.1

require (
	github.com/palomino513/arcanine/Maria v1.0.0
	github.com/palomino513/arcanine/Tools v1.0.0
)

replace (
	github.com/palomino513/arcanine/Maria => ./Maria
	github.com/palomino513/arcanine/Tools => ./Tools
)
