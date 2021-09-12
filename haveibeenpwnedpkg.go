package haveibeenpwnedpkg

import(
	"io/ioutil"
	"strings"
	"net/http"
	"strconv"
	"crypto/sha1"
	"encoding/hex"
)

var url = "https://api.pwnedpasswords.com/range/"

func HaveibeenpwnPassword(password string) (int, error) {
	hash := SHA1(password)
	hashList, err := GetHashList(hash)
	if err != nil {
		return 0, err
	}
	for _, s := range hashList {
		if strings.ToUpper(hash[5:]) == strings.Split(s, ":")[0] {
			how, _ := strconv.Atoi(strings.Split(s, ":")[1])
			return how, nil
		}
	}
	return 0, nil
}

func GetHashList(hash string) ([]string, error) {
	resp, err := http.Get(url + hash[0:5])
	if err != nil {
		return []string{}, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	hashList := strings.Split(string(respBody), "\r\n")
	return hashList, nil
}

func SHA1(plainText string) string {
	hash := sha1.New()
	hash.Write([]byte(plainText))
	result := hex.EncodeToString(hash.Sum(nil))
	return result
}
