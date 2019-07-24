package twitter

type configurator func(*Twitter)

// WithURL is a Twitter configuration function used to set the twitter client secret key
func WithURL(url string) func(*Twitter) {
	return func(t *Twitter) {
		t.url = url
	}
}

// WithAPIKey is a Twitter configuration function used to set the twitter client api key
func WithAPIKey(key string) func(*Twitter) {
	return func(t *Twitter) {
		t.apiKey = key
	}
}

// WithSecretKey is a Twitter configuration function used to set the twitter client api key
func WithSecretKey(key string) func(*Twitter) {
	return func(t *Twitter) {
		t.apiSecret = key
	}
}
