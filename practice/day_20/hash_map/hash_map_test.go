package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
    t.Run("Basic Operations", func(t *testing.T) {
        hm := NewHashMap()

        hm.Put("one", 1)
        hm.Put("two", 2)
        hm.Put("three", 3)

        if hm.Size() != 3 {
            t.Errorf("Expected size 3, got %d", hm.Size())
        }

        val, err := hm.Get("two")
        if err != nil || val != 2 {
            t.Errorf("Expected value 2 for key 'two', got %d", val)
        }

        if hm.IsEmpty() {
            t.Errorf("Expected hash map to be non-empty")
        }
    })

    t.Run("Update and Remove", func(t *testing.T) {
        hm := NewHashMap()

        hm.Put("one", 1)
        hm.Put("two", 2)
        hm.Put("three", 3)

        hm.Put("two", 22)
        val, err := hm.Get("two")
        if err != nil || val != 22 {
            t.Errorf("Expected updated value 22 for key 'two', got %d", val)
        }

        removed := hm.Remove("three")
        if !removed {
            t.Errorf("Expected 'three' to be removed")
        }

        _, err = hm.Get("three")
        if err == nil {
            t.Errorf("Expected 'three' to not be found after removal")
        }

        if hm.Size() != 2 {
            t.Errorf("Expected size 2 after removal, got %d", hm.Size())
        }
    })

    t.Run("Collision Handling", func(t *testing.T) {
        hm := NewHashMap()

        // These keys will likely cause collisions
        hm.Put("Aa", 1)
        hm.Put("BB", 2)

        val1, err1 := hm.Get("Aa")
        val2, err2 := hm.Get("BB")

        if err1 != nil || val1 != 1 {
            t.Errorf("Expected value 1 for key 'Aa', got %d", val1)
        }

        if err2 != nil || val2 != 2 {
            t.Errorf("Expected value 2 for key 'BB', got %d", val2)
        }
    })

    t.Run("Clear", func(t *testing.T) {
        hm := NewHashMap()

        hm.Put("one", 1)
        hm.Put("two", 2)
        hm.Put("three", 3)

        hm.Clear()

        if !hm.IsEmpty() {
            t.Errorf("Expected empty hash map after clear")
        }

        if hm.Size() != 0 {
            t.Errorf("Expected size 0 after clear, got %d", hm.Size())
        }
    })

    t.Run("Resizing", func(t *testing.T) {
        hm := NewHashMap()

        for i := 0; i < 100; i++ {
            hm.Put(fmt.Sprintf("key%d", i), i)
        }

        if hm.Size() != 100 {
            t.Errorf("Expected size 100, got %d", hm.Size())
        }

        for i := 0; i < 100; i++ {
            val, err := hm.Get(fmt.Sprintf("key%d", i))
            if err != nil || val != i {
                t.Errorf("Expected value %d for key 'key%d', got %d", i, i, val)
            }
        }
    })
}
