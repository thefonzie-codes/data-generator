package main

// Thinking about adding season and year as season const enum
// aka spring, summer, fall, winter
// then year as a separate value to make it easier to validate when entering the info
// eg. "spring" + "26" --> Season spring26
// a lot of the data is to be seeded for now - but I'm also debating on if I should make it a "living" mock company

type Apparel struct {
	UUID        string
	Name        string
	Description string
	Season      string
	Colour      string
}

// Size is the general sizing type for now, all the other names I could think of felt too clunky lol

type Size string

const (
	SizeXXS Size = "XXS"
	SizeXS  Size = "XS"
	SizeS   Size = "S"
	SizeM   Size = "M"
	SizeL   Size = "L"
	SizeXL  Size = "XL"
	SizeXXL Size = "XXL"
)

type PantSizeMens string

const (
	PantM26 PantSizeMens = "26"
	PantM28 PantSizeMens = "28"
	PantM30 PantSizeMens = "30"
	PantM32 PantSizeMens = "32"
	PantM34 PantSizeMens = "34"
	PantM36 PantSizeMens = "36"
	PantM38 PantSizeMens = "38"
	PantM40 PantSizeMens = "40"
)

func GenerateMenswear() {
	menswearList := []string{
		"Supima Cotton T-shirt",
		"Boat Neck T-shirt",
		"Raglan Sleeve Tee",
		"Pique Polo",
		"Merino Wool Sweater",
		"Non-Iron Dress Shirt",
		"Oxford Button-down Shirt",
		"Short Sleeve Button Down",
		"Distressed Workwear Jacket",
		"Suit Jacket",
		"Suit Pants",
		"Wool Dress Pant",
		"Chino Pant",
		"Slim-fit Denim",
		"Straight Leg Demim",
		"Silk Tie",
		"Knit Tie",
	}

	womenswear := []string{
		"Blouse",
		"Cotton Tee",
		"Skinny Fit Jeans",
		"Summer Dress",
		"Pleated Skirt",
		"Cozy Wool Sweater",
		"Wool Blazer",
		"Denim Jacket",
		"Dress Pants",
		"Moto Leather Jacket",
		"Distressed Denim Jacket",
		"Merino Wool Cardigan",
		"Cashmere Overcoat",
	}
}
