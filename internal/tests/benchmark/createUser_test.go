package benchmark

import (
	"fmt"
	"github.com/hibbannn/hexagonal-boilerplate/internal/adapters/cache"
	"github.com/hibbannn/hexagonal-boilerplate/internal/adapters/repository"
	"github.com/hibbannn/hexagonal-boilerplate/internal/core/domain"
	"testing"

	"github.com/jinzhu/gorm"
)

func BenchmarkCreateUser(b *testing.B) {
	db, err := gorm.Open("postgres", "postgres://hibban:postgres@localhost:5432/wdp_rnd?sslmode=disable")
	if err != nil {
		panic(err)
	}

	redisCache, err := cache.NewRedisCache("localhost:6379", "")
	if err != nil {
		panic(err)
	}

	store := repository.NewDB(db, redisCache)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		email := fmt.Sprintf("test_user_%d@example.com", i)
		password := "password"
		// Delete user if it exists
		var user domain.User
		if err := db.Where("email = ?", email).First(&user).Error; err == nil {
			if err := db.Delete(&user).Error; err != nil {
				b.Fatalf("failed to delete user: %v", err)
			}
		}
		_, err := store.CreateUser(email, password)
		if err != nil {
			b.Fatalf("failed to create test user: %v", err)
		}
	}
}
