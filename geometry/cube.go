package geometry

func CubeVolume(n int) (int,error) {
	if n!= 0 {
		return n * n * n, nil
	} else {
		0, errors.New("Zero length edge is not allowed")
	}
}
