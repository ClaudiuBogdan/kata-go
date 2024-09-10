package lrucache

import (
    "testing"
)

func TestLRUCache(t *testing.T) {
    t.Run("Basic Operations", func(t *testing.T) {
        lru := NewLRUCache(3)

        lru.Put(1, 1)
        lru.Put(2, 2)
        lru.Put(3, 3)

        if lru.Size() != 3 {
            t.Errorf("Expected size 3, got %d", lru.Size())
        }

        val, err := lru.Get(2)
        if err != nil || val != 2 {
            t.Errorf("Expected value 2 for key 2, got %d", val)
        }

        if lru.String() != "(2:2) (3:3) (1:1) " {
            t.Errorf("Unexpected cache content: %s", lru.String())
        }
    })

    t.Run("Eviction", func(t *testing.T) {
        lru := NewLRUCache(3)

        lru.Put(1, 1)
        lru.Put(2, 2)
        lru.Put(3, 3)
        lru.Put(4, 4)

        if lru.Size() != 3 {
            t.Errorf("Expected size 3 after eviction, got %d", lru.Size())
        }

        _, err := lru.Get(1)
        if err == nil {
            t.Errorf("Expected key 1 to be evicted")
        }

        if lru.String() != "(4:4) (3:3) (2:2) " {
            t.Errorf("Unexpected cache content after eviction: %s", lru.String())
        }
    })

    t.Run("Update Existing", func(t *testing.T) {
        lru := NewLRUCache(3)

        lru.Put(1, 1)
        lru.Put(2, 2)
        lru.Put(3, 3)
        lru.Put(2, 4)

        if lru.Size() != 3 {
            t.Errorf("Expected size 3 after update, got %d", lru.Size())
        }

        val, err := lru.Get(2)
        if err != nil || val != 4 {
            t.Errorf("Expected updated value 4 for key 2, got %d", val)
        }

        if lru.String() != "(2:4) (3:3) (1:1) " {
            t.Errorf("Unexpected cache content after update: %s", lru.String())
        }
    })

    t.Run("Clear", func(t *testing.T) {
        lru := NewLRUCache(3)

        lru.Put(1, 1)
        lru.Put(2, 2)
        lru.Put(3, 3)

        lru.Clear()

        if lru.Size() != 0 {
            t.Errorf("Expected size 0 after clear, got %d", lru.Size())
        }

        if lru.String() != "" {
            t.Errorf("Expected empty cache after clear, got: %s", lru.String())
        }
    })
}
