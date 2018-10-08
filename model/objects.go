package model

import "sort"

// CategoryMapKeys gets a full slice of keys from the CategoryMap below
func CategoryMapKeys() []string {
	keys := make([]string, len(CategoryMap))

	i := 0
	for k := range CategoryMap {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}

// CategoryMap is a map of human-readable categories to their Craigslist shorthand. This map is not currently a comprehensive mapping of the full list of categories
// This map is not currently immutable but should not be mutated anywhere in the codebase
var CategoryMap = map[string]string{
	"apts broker fee":                       "fee",
	"apts broker no fee":                    "nfb",
	"apts/housing for rent":                 "apa",
	"office & commercial":                   "off",
	"real estate - by broker":               "reb",
	"vacation rentals":                      "vac",
	"accounting/finance":                    "acc",
	"admin/office":                          "ofc",
	"architect/engineer/cad":                "egr",
	"art/media/design":                      "med",
	"business/mgmt":                         "bus",
	"customer service":                      "csr",
	"education/teaching":                    "edu",
	"et cetera":                             "etc",
	"food/beverage/hospitality":             "fbh",
	"general labor":                         "lab",
	"government":                            "gov",
	"healthcare":                            "hea",
	"human resource":                        "hum",
	"legal/paralegal":                       "lgl",
	"manufacturing":                         "mnu",
	"marketing/advertising/pr":              "mar",
	"nonprofit":                             "npo",
	"real estate":                           "rej",
	"retail/wholesale":                      "ret",
	"sales":                                 "sls",
	"salon/spa/fitness":                     "spa",
	"science/biotech":                       "sci",
	"security":                              "sec",
	"skilled trades/artisan":                "trd",
	"software/qa/dba/etc":                   "sof",
	"systems/networking":                    "sad",
	"technical support":                     "tch",
	"transportation":                        "trp",
	"tv/film/video/radio":                   "tfr",
	"web/html/info design":                  "web",
	"writing/editing":                       "wri",
	"appliances - by dealer":                "ppd",
	"atvs, utvs, snowmobiles - by dealer":   "snd",
	"auto parts - by dealer":                "ptd",
	"auto wheels & tires - by dealer":       "wtd",
	"boats - by dealer":                     "bod",
	"business/commercial - by dealer":       "bfd",
	"cars & trucks - by owner/dealer":       "cta",
	"cars & trucks - by dealer":             "ctd",
	"cars & trucks - by owner":              "cto",
	"cell phones - by dealer":               "mod",
	"computers - by dealer":                 "syd",
	"electronics - by dealer":               "eld",
	"farm & garden - by dealer":             "grq",
	"furniture - by dealer":                 "fud",
	"general for sale - by dealer":          "fod",
	"heavy equipment - by dealer":           "hvd",
	"household items - by dealer":           "hsd",
	"materials - by dealer":                 "mad",
	"musical instruments - by owner/dealer": "msa",
	"motorcycles/scooters - by dealer":      "mcd",
	"rvs - by dealer":                       "rvd",
	"tickets - by dealer":                   "tid",
	"trailers - by dealer":                  "trd",
}

// Regions is a slice of all Craiglist regions for the application to loop through on its quest to get results from every region
// It is not currently a comprehensive list of all regions
var Regions = []string{
	"albany",
	"allentown",
	"altoona",
	"annapolis",
	"baltimore",
	"binghamton",
	"catskills",
	"cnj",
	"charlottesville",
	"chambersburg",
	"delaware",
	"newlondon",
	"easternshore",
	"martinsburg",
	"elmira",
	"fingerlakes",
	"frederick",
	"fredericksburg",
	"harrisburg",
	"harrisonburg",
	"hartford",
	"hudsonvalley",
	"ithaca",
	"jerseyshore",
	"lancaster",
	"longisland",
	"newhaven",
	"newyork",
	"norfolk",
	"newjersey",
	"nwct",
	"oneonta",
	"philadelphia",
	"poconos",
	"reading",
	"richmond",
	"scranton",
	"southjersey",
	"smd",
	"pennstate",
	"syracuse",
	"utica",
	"washingtondc",
	"westernmass",
	"westmd",
	"williamsport",
	"winchester",
	"york",
}
