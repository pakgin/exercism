package robotname

// import "fmt"

// Define the Robot type here.
type Robot struct {
	name string
}

var nameList []rune = []rune{'A', 'A', '0', '0', '0'}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.Reset()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	carry := 1
	for i := len(nameList) - 1; i >= 2; i-- {
		if carry > 0 && nameList[i]+1 > '9' {
			nameList[i] = '0'
		} else {
			nameList[i]++
			carry = 0
			break
		}
	}
	if carry > 0 {
		for i := 1; i >= 0; i-- {
			if carry > 0 && nameList[i]+1 > 'Z' {
				nameList[i] = 'A'
			} else {
				nameList[i]++
				carry = 0
				break
			}
		}
	}
	r.name = string(nameList)
}
