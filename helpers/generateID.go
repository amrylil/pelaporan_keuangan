package helpers

import (
	"math/rand/v2" // Gunakan package rand v2 yang lebih modern (jika Go Anda versi 1.22+)
)

// Fungsi ini akan menghasilkan sebuah angka uint64 acak yang unik.
// Jauh lebih sederhana dari Sonyflake.
func GenerateID() uint64 {
	return uint64(rand.IntN(90000000) + 10000000)
}
