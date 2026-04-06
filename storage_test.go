package main

import (
	"testing"
)

func TestStorage(t *testing.T) {
	storage := NewStorage()
	all := *storage.ReadAll()

	if len(all) > 0 {
		t.Errorf("Storage should be empty.")
	}

	_, err := storage.Update(0, "Should not update.")
	if err == nil {
		t.Errorf("Existing item should have not been updated.")
	}

	storage.Write("test item!?-.,<>()")
	all2 := *storage.ReadAll()
	if len(all2) != 1 {
		t.Errorf("Storage should have one item")
	}

	written, err2 := storage.Read(0)
	if err2 != nil {
		t.Errorf("Storage should have an item")
	} else {
		if written.Value != "test item!?-.,()" {
			t.Errorf("Item should have been updated with clean value, value: %v", written.Value)
		}
	}

	item, err3 := storage.Update(0, "Should have been updated.")
	if err3 != nil {
		t.Errorf("Existing item should have been updated.")
	} else {
		if item.Value != "Should have been updated." {
			t.Errorf("Existing item should have been updated, value: %v", item.Value)
		}
	}

	item3, _ := storage.Read(0)
	if item3 == nil {
		t.Errorf("Item should not be nil.")
	}

	_, err4 := storage.Read(1)
	if err4 == nil {
		t.Errorf("should not be able to read non existing item")
	}

	storage.Remove(0)
	all3 := *storage.ReadAll()
	if len(all3) != 0 {
		t.Errorf("Storage should have zero items")
	}
}
