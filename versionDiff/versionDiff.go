package versionDiff

import "fmt"
import "time"

// versionDiff 比较版本号的天数差，例如2021.11.23.1537与2021.11.24.1736的天数差为-1，2021.11.23.1537与2021.11.14.1736的天数差为9
func versionDiff(version1 string, version2 string) int32 {
	if version1 == "" || version2 == "" {
		return 0
	}
	version1 = version1[0 : len(version1)-5]
	version2 = version2[0 : len(version2)-5]
	t1, _ := time.Parse("2006.01.02", version1)
	t2, _ := time.Parse("2006.01.02", version2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(int(t1.Sub(t2).Hours() / 24))

	return 0
}

func main() {
	a := "2021.11.23.1537"
	b := "2021.10.24.1537"
	versionDiff(a, b)
}
