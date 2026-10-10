// Package twofer determines the dialogue said by the user
// as they give their extra cookie to another person in the queue.
package twofer

// ShareWith returns the dialogue that should be spoken by the user
func ShareWith(name string) string {
	if name == "" {
        return "One for you, one for me."
    } else {
        return "One for " + name + ", one for me."
    }
}
