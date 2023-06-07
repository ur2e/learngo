package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word alreadys exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil // error 아닐 경우는 nil
	}

	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	// 만약 추가하려는 단어가 존재 X라면 -> 추가하면 된다.

	_, err := d.Search(word) // 당장은 definition 필요없으니까 _처리
	if err == errNotFound {
		d[word] = def
	} else if err == nil { // 단어가 이미 존재하는 경우
		return errWordExists
	}

	return nil
}

// Update a word
func (d Dictionary) Update(word, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDef
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return errNotFound
	}
	delete(d, word)
	return nil
}
