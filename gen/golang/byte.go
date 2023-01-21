/*
 *  Copyright (c) 2017, https://github.com/snow1emperor
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gengolang

import "unicode"

// IsLower check letter is lower case or not
func IsLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsUpper check letter is upper case or not
func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLetter check character is a letter or not
func IsLetter(b byte) bool {
	return IsLower(b) || IsUpper(b)
}

// IsSpaceQuote return wehter a byte is space or quote characters
func IsSpaceQuote(b byte) bool {
	return IsSpace(b) || b == '"' || b == '\''
}

// IsSpace only call unicode.IsSpace
func IsSpace(b byte) bool {
	return unicode.IsSpace(rune(b))
}

// ToLower convert a byte to lower case
func ToLower(b byte) byte {
	if IsUpper(b) {
		b = b - 'A' + 'a'
	}

	return b
}

// ToUpper convert a byte to upper case
func ToUpper(b byte) byte {
	if IsLower(b) {
		b = b - 'a' + 'A'
	}

	return b
}

// ToLowerString convert a byte to lower case string
func ToLowerString(b byte) string {
	if IsUpper(b) {
		b = b - 'A' + 'a'
	}

	return string(b)
}

// ToUpperString convert a byte to upper case string
func ToUpperString(b byte) string {
	if IsLower(b) {
		b = b - 'a' + 'A'
	}

	return string(b)
}
