package nomads

type Spaceship [5][5]int

var Glider = Spaceship{
	{0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 0, 1, 1, 0},
	{0, 0, 0, 0, 0},
}

func RotateGlider(glider Spaceship) Spaceship {
	n := len(glider[0])

	for i := 0; i < len(glider[0]); i++ {
		for j := i; j < n-i-1; j++ {
			temp := glider[i][j]

			glider[i][j] = glider[n-1-j][i]
			glider[n-1-j][i] = glider[n-1-i][n-1-j]
			glider[n-1-i][n-1-j] = glider[j][n-1-i]
			glider[j][n-1-i] = temp
		}
	}

	return glider
}
