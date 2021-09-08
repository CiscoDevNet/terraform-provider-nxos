package provider

import "time"

// Retry defaults
const Factor = 3
const MinDelay = 4 * time.Second
const MaxDelay = 60 * time.Second
