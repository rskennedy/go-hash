package main

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	// "testing"
)


/*** GLOBAL VARIABLES ***/

type Hash_t struct {
	refs	  uint32
	cap       uint32
	hash_dev hash.Hash32
	table     []interface{}
}


/*** API ***/

func CreateHash(num_elem uint32) (*Hash_t, error) {
	var h *Hash_t

	h = new(Hash_t)
	if h == nil {
		return nil, errors.New("Hash table creation failed.")
	}

	h.cap = num_elem
	h.hash_dev = fnv.New32a()
	h.table = make([]interface{}, num_elem)

	if h.table == nil || h.hash_dev == nil {
		return nil, errors.New("Failed to initialize hash table.")
	}
	return h, nil
}

func SetEntry(h *Hash_t, key string, value interface{}) (bool, error) {
	if value == nil || h == nil || key == "" {
		return false, errors.New("Failed to set hash key.")
	}
	i, err := hash_compute(h, key)
	if err != nil {
		return false, err
	}
	if h.table[i] != nil {
		return false, nil
	}

	h.table[i] = value
	return true, nil
}

/* returns a value based on the key. It can also optionally provide
 * an index in addition to the value
 * */
func GetEntry(h *Hash_t, key string, index *uint32) (interface{}, error) {
	if (h == nil) {
		return nil, errors.New("Hash Table does not exist!")
	}
	i, err := hash_compute(h, key)
	if err != nil {
		return nil, err
	}
	if index != nil {
		*index = i
	}
	return h.table[i], nil
}

func DeleteEntry(h* Hash_t, key string) (interface{}, error) {
	var index uint32

	tmp, err := GetEntry(h, key, &index)
	if (err != nil) {
		return nil, err
	}
	h.table[index] = nil
	return tmp, nil
}

func GetLoad(h* Hash_t) (float32, error) {
	count := 0

	if (h == nil) {
		return -1, errors.New("Hash Table does not exist!")
	}

	for _, x := range h.table {
		if x == nil {
			continue
		}
		count += 1
	}
	return float32(count)/float32(h.cap), nil
}


/*** INTERNAL FUNCTIONS ***/

/* Hash compute uses the hashing device found in  Hash_t.
 * That hashing device uses a non-cryptographic hashing function
 * written by Glenn Fowler, Landon Curt Noll, and Phong Vo.
 * */
func hash_compute(h *Hash_t, key string) (uint32, error) {
	var err error

	if (h == nil) {
		return 0, errors.New("Hash Table does not exist!")
	}

	h.hash_dev.Reset()			//clear previous input
	_, err = h.hash_dev.Write([]byte(key))
	if err != nil {
		return 0, errors.New("Failed to compute hash.")
	}
	return h.hash_dev.Sum32() % h.cap, nil
}


/*** TESTING ***/

func main() {
	var ret bool

	/* Testing for correctness */
	h, _ := CreateHash(5)
	ret,_ = SetEntry(h, "hello", "World!")
	ret,_ = SetEntry(h, "hello", "World!")
	s, _ := GetEntry(h, "hello", nil)
	ret,_ = SetEntry(h, "yo", "What up!")
	ret,_ = SetEntry(h, "ayy", "lmao")
	val,_ := DeleteEntry(h, "ayy")
	SetEntry(h, "Hi", "fam!")
	SetEntry(h, "aaldfjalkfj", "Fam!")
	SetEntry(h, "hi hi hi", "Bruv!")
	fmt.Println(*h)
	fmt.Println(s)
	fmt.Println(ret)
	fmt.Println(val)
}
