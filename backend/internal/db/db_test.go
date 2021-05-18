package db

import (
	"strconv"
	"strings"
	"testing"
)

func strArrToStr(arr []string) string {
	var str string = strings.Join(arr, "")
	return str
}

func boolToStr(val bool) string {
	str := strconv.FormatBool(val)
	return str
}

var imgArrayReal []string = []string{"https://www.restu.cz/ir/restaurant/62a/62a74be7fd17903cecc67daaa16e5dd3--cxc400.jpg", "https://www.restu.cz/ir/restaurant/930/930d5a3dd4ee7257056da1f9c2fa0f47--cxc400.jpg", "https://www.restu.cz/ir/restaurant/06c/06c00276d16341d261870d905ec84875--cxc400.jpg", "https://www.restu.cz/ir/restaurant/851/8514b4149f4e3b0aa47f5df296bc34ba--cxc400.jpg", "https://www.restu.cz/ir/restaurant/51a/51a11aabf1adcb805b37d55d3c986da5--cxc400.jpg", "https://www.restu.cz/ir/restaurant/b49/b4942472bd90b49296bfd291ed615bdb--cxc400.jpg", "https://www.restu.cz/ir/restaurant/bd7/bd749067bebcc31eb18adf407da9dbde--cxc400.jpg", "https://www.restu.cz/ir/restaurant/e9c/e9cd098c929b96216135b08bc1df1c5c--cxc400.jpg", "https://www.restu.cz/ir/restaurant/810/8107afb1753aea23471e7d510e7e879e--cxc400.jpg", "https://www.restu.cz/ir/restaurant/695/6950329648ad8c7dad1f424b2a4d2d6d--cxc400.png", "https://www.restu.cz/ir/restaurant/8d9/8d9b7c20ad2dc83921d6cce2da494310--cxc400.jpg", "https://www.restu.cz/ir/restaurant/4fb/4fb3a98011df503cb00f7dd2416c5189--cxc400.jpg", "https://www.restu.cz/ir/restaurant/583/583bc353270048def58bc87b30de7569--cxc400.jpg", "https://www.restu.cz/ir/restaurant/026/026b4da6b58699eddf4dba04e3641209--cxc400.jpg", "https://www.restu.cz/ir/restaurant/7ed/7ed3a893f7e1a6f4840311c4a1bf5c73--cxc400.jpg", "https://www.restu.cz/ir/restaurant/1f7/1f7d7e84cf0c59cb7aa339bf171900dc--cxc400.jpg", "https://www.restu.cz/ir/restaurant/4f2/4f2be5aeef5a184fe149a5b63fd94a3d--cxc400.jpg", "https://www.restu.cz/ir/restaurant/bad/bad5f2e5f0de36d3a90dc12a372f7004--cxc400.jpg", "https://www.restu.cz/ir/restaurant/daa/daa150fa4f4f16f681fc5bdcd9eeae11--cxc400.jpg", "https://www.restu.cz/ir/restaurant/da5/da537d1939448bc30fe1e4cb7d379de1--cxc400.jpg", "https://www.restu.cz/ir/restaurant/02f/02f6a6248e88243cd7a79d15d59fabde--cxc400.jpg", "https://www.restu.cz/ir/restaurant/540/540a5fe6576705ac8c1a2a68346dd1ba--cxc400.jpg", "https://www.restu.cz/ir/restaurant/7fe/7feb8e8e26ea37a4e3912c1b3a2f6744--cxc400.jpg", "https://www.restu.cz/ir/restaurant/c0c/c0caa926306aad9c98c3a37973269451--cxc400.jpg", "https://www.restu.cz/ir/restaurant/695/695e20c211c423a49999a1dae563a7d0--cxc400.jpg", "https://www.restu.cz/ir/restaurant/e04/e042489cfbde3f6c6ed91fa804af3060--cxc400.jpg", "https://www.restu.cz/ir/restaurant/afd/afd028cc32f8a2a11d9acda0d965277e--cxc400.jpg"}
var cuisineArrayReal []string = []string{"Czech", "International"}
var openHoursArrayReal []string = []string{"Friday", ":", "Zavřeno", ",", "Monday", ":", "Zavřeno", ",", "Saturday", ":", "Zavřeno", ",", "Sunday", ":", "Zavřeno", ",", "Thursday", ":", "Zavřeno", ",", "Tuesday", ":", "Zavřeno", ",", "Wednesday", ":", "Zavřeno"}
var openHoursStrReal string = strArrToStr(openHoursArrayReal)

var imgArrayGarbage []string = []string{"bushDidNineEleven.exe"}
var cuisineArrayGarbage []string = []string{"from", "out", "of", "this", "world"}
var openHoursStrGarbage string = "What are opening hours?"

// actual restaurant pulled from the DB
var sampleRestaurant = RestaurantDB{
	1,
	"Zlatá Praha",
	"Pařížská 30",
	"Praha 1",
	imgArrayReal,
	cuisineArrayReal,
	"300 - 600 Kč",
	"5.0",
	"http://www.zlatapraharestaurant.cz",
	"+420 296 630 914",
	50.09185325,
	14.418952034421178,
	false,
	false,
	false,
	"",
	openHoursStrReal,
	false,
	nil,
}

// restaurant struct filled with garbage data
var sampleRestaurantGarbage = RestaurantDB{
	12312413,
	"Not A Real Restaurant",
	"Middle Of Nowhere",
	"In The Boonies",
	imgArrayGarbage,
	cuisineArrayGarbage,
	"1000 - 9999 Kč",
	"10.2",
	"http://penisenlargement.xxx",
	"+7 912 123 45 67 data data data",
	10000000.12345,
	-918629533.98765,
	true,
	true,
	false,
	"don't have one, boss",
	openHoursStrGarbage,
	false,
	nil,
}

func TestIsInRadiusFalse(t *testing.T) {
	got := sampleRestaurant.IsInRadius(50.078702, 14.439827, "100") //random coords next to praguecollege
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInRadiusTrue(t *testing.T) {
	got := sampleRestaurant.IsInRadius(50.078702, 14.439827, "5000") //random coords next to praguecollege
	want := true
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInRadiusGarbage(t *testing.T) {
	got := sampleRestaurantGarbage.IsInRadius(50.078702, 14.439827, "5000") //random coords next to praguecollege
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInPriceRangeFalse(t *testing.T) {
	got := sampleRestaurant.IsInPriceRange("600-")
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInPriceRangeTrue(t *testing.T) {
	got := sampleRestaurant.IsInPriceRange("300-600")
	want := true
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInPiceRangeGarbage(t *testing.T) {
	got := sampleRestaurantGarbage.IsInPriceRange("300-600")
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInDistrictFalse(t *testing.T) {
	got := sampleRestaurant.IsInDistrict("Praha 2")
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInDistrictTrue(t *testing.T) {
	got := sampleRestaurant.IsInDistrict("Praha 1")
	want := true
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestIsInDistrictGarbage(t *testing.T) {
	got := sampleRestaurantGarbage.IsInDistrict("Praha 1")
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestHasCuisineFalse(t *testing.T) {
	got := sampleRestaurant.HasCuisines("Czech, International, American")
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestHasCuisineTrue(t *testing.T) {
	got := sampleRestaurant.HasCuisines("International, Czech")
	want := true
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}

func TestHasCuisineGarbage(t *testing.T) {
	got := sampleRestaurantGarbage.HasCuisines("International")
	want := false
	if got != want {
		t.Errorf("got %q, wanted %q", boolToStr(got), boolToStr(want))
	}
}
