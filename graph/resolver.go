package graph

import (
	"gitlab.com/morawskioz/camp_me_go/db"
)

type Resolver struct {
	Prisma *db.PrismaClient
}
