package handling

import (
	"crypto/md5"
	"fmt"
)

func hash(data []byte, control string) (status int, hash string) {
	status = 0

	if data == nil {
		status = 0
		return
	}

	hash = fmt.Sprintf("%x", md5.Sum(data))
	if hash == control {
		status = 1
	}

	return
}
