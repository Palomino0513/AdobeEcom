module arcanine

go 1.22.2

require (
	github.com/palomino513/arcanine/Tools v1.0.0
	github.com/palomino513/arcanine/Maria v1.0.0
)

replace (
	github.com/palomino513/arcanine/Tools => ./Tools
	github.com/palomino513/arcanine/Maria => ./Maria
)