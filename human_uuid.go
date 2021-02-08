package humanuuid

import (
	"fmt"
	"github.com/goombaio/namegenerator"
	uuid "github.com/satori/go.uuid"
	"regexp"
	"time"
)

// Regexp is used to check if a string is a UUID we created using the New() function
var Regexp = regexp.MustCompile(`^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[A-Za-z]+-[A-Za-z]+$`)

// New creates a new human readable UUID.
// A "human readable UUID" is simply the first half of a full UUID + two random words at the end.
// This makes the IDs easier to read, track and communicate between humans while preserving enough uniqueness.
// Examples of output:
//		- 7e1e2527-62d6-divine-grass
//		- e6633073-dff0-broken-moon
//		- 2a761d30-a2c1-weathered-field
func New() string {
	name := genRandomName()
	uid := genRandomUUIDAndTakeHalf()
	return fmt.Sprintf("%s-%s", uid, name)
}

// IsValid checks if the given string is a valid human readable UUID.
func IsValid(val string) bool {
	return Regexp.MatchString(val)
}

// genRandomUUIDAndTakeHalf generates a full v4 UUID and takes the first 13 characters
func genRandomUUIDAndTakeHalf() string {
	uid := uuid.NewV4()
	return uid.String()[:13]
}

// genRandomName uses github.com/goombaio/namegenerator to generate a random name consisting of two random words
func genRandomName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return nameGenerator.Generate()
}
