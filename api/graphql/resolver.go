package graphql

import (
	"gitlab.com/morawskioz/camp_me_go/internal/prisma/db"
)

type Resolver struct {
	Prisma *db.PrismaClient
}
