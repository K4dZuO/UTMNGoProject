import requests
import random
import time

URL = "http://localhost:8083/get_top_category?categoryName="

CATEGORIES = [
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
]

def get_category_top():
    category = random.choice(CATEGORIES)
    url = URL + category

    try:
        r = requests.get(url)
        print(f"[{category}] Status:", r.status_code)
        print("Response:", r.text[:500], "\n")  # ограничим вывод
    except Exception as e:
        print("Error:", e)

if __name__ == "__main__":
    print("Requesting top-10 category every second...")
    while True:
        get_category_top()
        time.sleep(1)
