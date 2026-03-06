package core

import (
	"github.com/teamgram/proto/mtproto"
)

func (c *ConfigurationCore) HelpGetCountriesList(in *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error) {
	countries := []*mtproto.Help_Country{
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "CN", DefaultName: "China",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "86", Prefixes: []string{}, Patterns: []string{"XXX XXXX XXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "US", DefaultName: "United States",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "1", Prefixes: []string{}, Patterns: []string{"XXX XXX XXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "GB", DefaultName: "United Kingdom",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "44", Prefixes: []string{}, Patterns: []string{"XXXX XXXXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "RU", DefaultName: "Russia",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "7", Prefixes: []string{}, Patterns: []string{"XXX XXX XXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "DE", DefaultName: "Germany",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "49", Prefixes: []string{}, Patterns: []string{"XXXX XXXXXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "JP", DefaultName: "Japan",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "81", Prefixes: []string{}, Patterns: []string{"XX XXXX XXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "KR", DefaultName: "South Korea",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "82", Prefixes: []string{}, Patterns: []string{"XX XXXX XXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "IN", DefaultName: "India",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "91", Prefixes: []string{}, Patterns: []string{"XXXXX XXXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "SG", DefaultName: "Singapore",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "65", Prefixes: []string{}, Patterns: []string{"XXXX XXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "HK", DefaultName: "Hong Kong",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "852", Prefixes: []string{}, Patterns: []string{"XXXX XXXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
		mtproto.MakeTLHelpCountry(&mtproto.Help_Country{
			Hidden: false, Iso2: "TW", DefaultName: "Taiwan",
			CountryCodes: []*mtproto.Help_CountryCode{
				mtproto.MakeTLHelpCountryCode(&mtproto.Help_CountryCode{CountryCode: "886", Prefixes: []string{}, Patterns: []string{"XXX XXX XXX"}}).To_Help_CountryCode(),
			},
		}).To_Help_Country(),
	}
	return mtproto.MakeTLHelpCountriesList(&mtproto.Help_CountriesList{Countries: countries, Hash: 0}).To_Help_CountriesList(), nil
}
