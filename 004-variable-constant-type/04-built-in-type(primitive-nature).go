/* Trim function is part of standard library from string package */
/* With the help of this function we can understand the concept of primitive nature */

/** DON'T EXECUTE, JUST for demonstration **/

// In both the return a copy of the new value is passed.
// User should take care of the return value based on requirement.
/* This is best practice and is used across system libraries,
Hence even if we can send pointers for built-in type, we will send values so that function or method can 
work on copy of the data */
func trim(s string, cutset string) string {
	if s == "" | cutset == "" {
		return s // we are operating on a different copy. We can but we don't pass string as pointer.
	}

	return TrimFunc(s, makeCutsetFunc(cutset))
}