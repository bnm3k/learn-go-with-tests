package c8

import "testing"

func TestSearch(t *testing.T) {
	key, entry := "foo", "bar"
	dict := Dictionary{key: entry}

	t.Run("Known word", func(t *testing.T) {
		assertDefinition(t, dict, key, entry)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	key, entry := "foo", "bar"
	dict := Dictionary{}

	t.Run("new word", func(t *testing.T) {
		err := dict.Add(key, entry)
		assertNoError(t, err)
		assertDefinition(t, dict, key, entry)
	})

	t.Run("existing word", func(t *testing.T) {
		err := dict.Add(key, entry)
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	key, entry, newEntry := "foo", "bar", "quz"
	dict := Dictionary{key: entry}

	t.Run("existing word", func(t *testing.T) {
		err := dict.Update(key, newEntry)
		assertNoError(t, err)
		assertDefinition(t, dict, key, newEntry)
	})

	t.Run("non-existing word", func(t *testing.T) {
		err := dict.Update("joe", "mama")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key, entry := "foo", "bar"
	dict := Dictionary{key: entry}

	t.Run("existing word", func(t *testing.T) {
		errOnDelete := dict.Delete(key)
		assertNoError(t, errOnDelete)
		_, errOnSearchDeletedEntry := dict.Search(key)
		if errOnSearchDeletedEntry != ErrNotFound {
			t.Errorf("expected %q to be deleted", key)
		}
	})

	t.Run("non-existing word", func(t *testing.T) {
		err := dict.Delete("goo")
		assertError(t, err, ErrDeleteWordDoesNotExist)
	})
}

func assertDefinition(t *testing.T, dict Dictionary, key, want string) {
	t.Helper()
	got, err := dict.Search(key)
	assertNoError(t, err)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("got an error in a correct scenario")
	}
}
