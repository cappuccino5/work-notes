
func TestEarthDistance(t *testing.T) {
	lat1 := 22.601986
	lng1 := 114.119934

	lat2 :=22.609756
	lng2 := 114.029397
	result := earthDistance(lat1, lng1, lat2, lng2)
	fmt.Println(result)
}
func earthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000) // 赤道半径 6378137
	rad := math.Pi/180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1) * math.Sin(lat2) + math.Cos(lat1) * math.Cos(lat2) * math.Cos(theta))

	return dist * radius
}
