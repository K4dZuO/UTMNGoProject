package seeder

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)
var categories = []string{
    "Smartphones", "Laptops", "Dishes", "Shoes", "Shirts", "Headphones",
    "Monitors", "Keyboards", "Mice", "Cameras",
    "Furniture", "Books", "Toys", "Food", "Drinks",
    "Sports", "Tools", "Garden", "Auto", "Beauty",
    "Jewelry", "Watches", "Bags", "Wallets", "Sunglasses",
    "Tablets", "Printers", "Speakers", "Gaming", "Consoles",
    "TVs", "Home Appliances", "Kitchenware", "Bedding", "Bath",
    "Pet Supplies", "Baby Products", "Stationery", "Office Supplies",
    "Musical Instruments", "Art Supplies", "Fitness", "Outdoor Gear",
    "Travel", "Party Supplies", "Cleaning", "Lighting", "Decor",
    "Smart Home", "Wearables", "Drones", "Software", "Movies",
    "Music", "Video Games", "Collectibles", "Antiques", "Crafts",
    "Industrial", "Medical", "Safety", "Educational",
}

func SeedCategories(ctx context.Context, pool *pgxpool.Pool) error {
	for _, name := range categories {
		_, err := pool.Exec(ctx,
			`INSERT INTO categories(name)
			 VALUES ($1)
			 ON CONFLICT (name) DO NOTHING`,
			name,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
