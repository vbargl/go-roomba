package httputil

import (
	"fmt"
	"io"
	"net/http"
)

const bodyLimit = 1 << 20

func CheckAndLimitBody(resp *http.Response, err error) error {
	if err != nil {
		return err
	}

	resp.Body = io.NopCloser(io.LimitReader(resp.Body, bodyLimit))
	if resp.StatusCode != http.StatusOK {
		content, err := io.ReadAll(resp.Body)
		switch {
		case err == io.ErrUnexpectedEOF:
			content = []byte(fmt.Sprintf("response body was too big"))
		case err != nil:
			content = []byte(fmt.Sprintf("failed to read response body: %v", err))
		}
		return fmt.Errorf("unexpected response: %d %s", resp.StatusCode, content)
	}

	return nil
}
