package db

func Connect() (*PrismaClient, error) {
	client := NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	return client, nil
}
