package wallet

import (
	"be11/apiclean/features/user"
	"time"
)

type Core struct {
	ID        uint
	Jenis     string
	Nomor     string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.Core
}
