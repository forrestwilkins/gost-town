life_bin = ./bin/life
fractals_bin = ./bin/fractals

rects: 
	go build -o $(life_bin) ./internal/life/main.go

fractal:
	go build -o $(fractals_bin) ./internal/fractals/main.go