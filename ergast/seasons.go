package ergast

import "fmt"
import "net/http"

//TODO add year verification so we don't go before 1950 or into the future.
func (c *Client) GetSeason(year int) ([]byte, error) {
	url := fmt.Sprintf(c.BaseURL+"%d.json", year)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := doRequest(req)
	if err != nil {
		return nil, err
	}
	s := string(bytes[:])
	fmt.Print(s)
	return bytes, nil
}
