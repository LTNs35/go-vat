package vat

import (
	"fmt"
	"testing"

	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/mocks"
)

func TestCheckVAT(t *testing.T) {

	type args struct {
		values []string
		want   bool
	}

	tests := []struct {
		name    string
		args    []args
		country countries.Calculer
	}{
		{
			name: "Albania",
			args: []args{
				{
					values: mocks.AlbaniaValidTests,
					want:   true,
				},
				{
					values: mocks.AlbaniaInvalidTests,
					want:   false,
				},
			},
			country: countries.Albania,
		},
		{
			name: "Andorra",
			args: []args{
				{
					values: mocks.AndorraValidTests,
					want:   true,
				},
				{
					values: mocks.AndorraInvalidTests,
					want:   false,
				},
			},
			country: countries.Andorra,
		},
		{
			name: "Argentina",
			args: []args{
				{
					values: mocks.ArgentinaValidTests,
					want:   true,
				},
				{
					values: mocks.ArgentinaInvalidTests,
					want:   false,
				},
			},
			country: countries.Argentina,
		},
		{
			name: "Austria",
			args: []args{
				{
					values: mocks.AustriaValidTests,
					want:   true,
				},
				{
					values: mocks.AustriaInvalidTests,
					want:   false,
				},
			},
			country: countries.Austria,
		},
		{
			name: "Belarus",
			args: []args{
				{
					values: mocks.BelarusValidTests,
					want:   true,
				},
				{
					values: mocks.BelarusInvalidTests,
					want:   false,
				},
			},
			country: countries.Belarus,
		},
		{
			name: "Belgium",
			args: []args{
				{
					values: mocks.BelgiumValidTests,
					want:   true,
				},
				{
					values: mocks.BelgiumInvalidTests,
					want:   false,
				},
			},
			country: countries.Belgium,
		},
		{
			name: "Bolivia",
			args: []args{
				{
					values: mocks.BoliviaValidTests,
					want:   true,
				},
				{
					values: mocks.BoliviaInvalidTests,
					want:   false,
				},
			},
			country: countries.Bolivia,
		},
		{
			name: "Brazil",
			args: []args{
				{
					values: mocks.BrazilValidTests,
					want:   true,
				},
				{
					values: mocks.BrazilInvalidTests,
					want:   false,
				},
			},
			country: countries.Brazil,
		},
		{
			name: "Bulgaria",
			args: []args{
				{
					values: mocks.BulgariaValidTests,
					want:   true,
				},
				{
					values: mocks.BulgariaInvalidTests,
					want:   false,
				},
			},
			country: countries.Bulgaria,
		},
		{
			name: "Canada",
			args: []args{
				{
					values: mocks.CanadaValidTests,
					want:   true,
				},
				{
					values: mocks.CanadaInvalidTests,
					want:   false,
				},
			},
			country: countries.Canada,
		},
		{
			name: "Croatia",
			args: []args{
				{
					values: mocks.CroatiaValidTests,
					want:   true,
				},
				{
					values: mocks.CroatiaInvalidTests,
					want:   false,
				},
			},
			country: countries.Croatia,
		},
		{
			name: "Cyprus",
			args: []args{
				{
					values: mocks.CyprusValidTests,
					want:   true,
				},
				{
					values: mocks.CyprusInvalidTests,
					want:   false,
				},
			},
			country: countries.Cyprus,
		},
		{
			name: "Czech Republic",
			args: []args{
				{
					values: mocks.CzechRepublicValidTests,
					want:   true,
				},
				{
					values: mocks.CzechRepublicInvalidTests,
					want:   false,
				},
			},
			country: countries.CzechRepublic,
		},
		{
			name: "Denmark",
			args: []args{
				{
					values: mocks.DenmarkValidTests,
					want:   true,
				},
				{
					values: mocks.DenmarkInvalidTests,
					want:   false,
				},
			},
			country: countries.Denmark,
		},
		{
			name: "Estonia",
			args: []args{
				{
					values: mocks.EstoniaValidTests,
					want:   true,
				},
				{
					values: mocks.EstoniaInvalidTests,
					want:   false,
				},
			},
			country: countries.Estonia,
		},
		{
			name: "Finland",
			args: []args{
				{
					values: mocks.FinlandValidTests,
					want:   true,
				},
				{
					values: mocks.FinlandInvalidTests,
					want:   false,
				},
			},
			country: countries.Finland,
		},
		{
			name: "France",
			args: []args{
				{
					values: mocks.FranceValidTests,
					want:   true,
				},
				{
					values: mocks.FranceInvalidTests,
					want:   false,
				},
			},
			country: countries.France,
		},
		{
			name: "Germany",
			args: []args{
				{
					values: mocks.GermanyValidTests,
					want:   true,
				},
				{
					values: mocks.GermanyInvalidTests,
					want:   false,
				},
			},
			country: countries.Germany,
		},
		{
			name: "Greece",
			args: []args{
				{
					values: mocks.GreeceValidTests,
					want:   true,
				},
				{
					values: mocks.GreeceInvalidTests,
					want:   false,
				},
			},
			country: countries.Greece,
		},
		{
			name: "Hong Kong",
			args: []args{
				{
					values: mocks.HongKongValidTests,
					want:   true,
				},
				{
					values: mocks.HongKongInvalidTests,
					want:   false,
				},
			},
			country: countries.HongKong,
		},
		{
			name: "Hungary",
			args: []args{
				{
					values: mocks.HungaryValidTests,
					want:   true,
				},
				{
					values: mocks.HungaryInvalidTests,
					want:   false,
				},
			},
			country: countries.Hungary,
		},
		{
			name: "Italy",
			args: []args{
				{
					values: mocks.ItalyValidTests,
					want:   true,
				},
				{
					values: mocks.ItalyInvalidTests,
					want:   false,
				},
			},
			country: countries.Italy,
		},
		{
			name: "Ireland",
			args: []args{
				{
					values: mocks.IrelandValidTests,
					want:   true,
				},
				{
					values: mocks.IrelandInvalidTests,
					want:   false,
				},
			},
			country: countries.Ireland,
		},
		{
			name: "Kazakhstan",
			args: []args{
				{
					values: mocks.KazakhstanValidTests,
					want:   true,
				},
				{
					values: mocks.KazakhstanInvalidTests,
					want:   false,
				},
			},
			country: countries.Kazakhstan,
		},
		{
			name: "Latvia",
			args: []args{
				{
					values: mocks.LatviaValidTests,
					want:   true,
				},
				{
					values: mocks.LatviaInvalidTests,
					want:   false,
				},
			},
			country: countries.Latvia,
		},
		{
			name: "Liechtenstein",
			args: []args{
				{
					values: mocks.LiechtensteinValidTests,
					want:   true,
				},
				{
					values: mocks.LiechtensteinInvalidTests,
					want:   false,
				},
			},
			country: countries.Liechtenstein,
		},
		{
			name: "Lithuania",
			args: []args{
				{
					values: mocks.LithuaniaValidTests,
					want:   true,
				},
				{
					values: mocks.LithuaniaInvalidTests,
					want:   false,
				},
			},
			country: countries.Lithuania,
		},
		{
			name: "Luxembourg",
			args: []args{
				{
					values: mocks.LuxembourgValidTests,
					want:   true,
				},
				{
					values: mocks.LuxembourgInvalidTests,
					want:   false,
				},
			},
			country: countries.Luxembourg,
		},
		{
			name: "Malta",
			args: []args{
				{
					values: mocks.MaltaValidTests,
					want:   true,
				},
				{
					values: mocks.MaltaInvalidTests,
					want:   false,
				},
			},
			country: countries.Malta,
		},
		{
			name: "Netherlands",
			args: []args{
				{
					values: mocks.NetherlandsValidTests,
					want:   true,
				},
				{
					values: mocks.NetherlandsInvalidTests,
					want:   false,
				},
			},
			country: countries.Netherlands,
		},
		{
			name: "North Macedonia",
			args: []args{
				{
					values: mocks.NorthMacedoniaValidTests,
					want:   true,
				},
				{
					values: mocks.NorthMacedoniaInvalidTests,
					want:   false,
				},
			},
			country: countries.NorthMacedonia,
		},
		{
			name: "Norway",
			args: []args{
				{
					values: mocks.NorwayValidTests,
					want:   true,
				},
				{
					values: mocks.NorwayInvalidTests,
					want:   false,
				},
			},
			country: countries.Norway,
		},
		{
			name: "Peru",
			args: []args{
				{
					values: mocks.PeruValidTests,
					want:   true,
				},
				{
					values: mocks.PeruInvalidTests,
					want:   false,
				},
			},
			country: countries.Peru,
		},
		{
			name: "Poland",
			args: []args{
				{
					values: mocks.PolandValidTests,
					want:   true,
				},
				{
					values: mocks.PolandInvalidTests,
					want:   false,
				},
			},
			country: countries.Poland,
		},
		{
			name: "Portugal",
			args: []args{
				{
					values: mocks.PortugalValidTests,
					want:   true,
				},
				{
					values: mocks.PortugalInvalidTests,
					want:   false,
				},
			},
			country: countries.Portugal,
		},
		{
			name: "Romania",
			args: []args{
				{
					values: mocks.RomaniaValidTests,
					want:   true,
				},
				{
					values: mocks.RomaniaInvalidTests,
					want:   false,
				},
			},
			country: countries.Romania,
		},
		{
			name: "Russia",
			args: []args{
				{
					values: mocks.RussiaValidTests,
					want:   true,
				},
				{
					values: mocks.RussiaInvalidTests,
					want:   false,
				},
			},
			country: countries.Russia,
		},
		{
			name: "San Marino",
			args: []args{
				{
					values: mocks.SanMarinoValidTests,
					want:   true,
				},
				{
					values: mocks.SanMarinoInvalidTests,
					want:   false,
				},
			},
			country: countries.SanMarino,
		},
		{
			name: "Serbia",
			args: []args{
				{
					values: mocks.SerbiaValidTests,
					want:   true,
				},
				{
					values: mocks.SerbiaInvalidTests,
					want:   false,
				},
			},
			country: countries.Serbia,
		},
		{
			name: "Slovakia",
			args: []args{
				{
					values: mocks.SlovakiaValidTests,
					want:   true,
				},
				{
					values: mocks.SlovakiaInvalidTests,
					want:   false,
				},
			},
			country: countries.Slovakia,
		},
		{
			name: "Slovenia",
			args: []args{
				{
					values: mocks.SloveniaValidTests,
					want:   true,
				},
				{
					values: mocks.SloveniaInvalidTests,
					want:   false,
				},
			},
			country: countries.Slovenia,
		},
		{
			name: "Spain",
			args: []args{
				{
					values: mocks.SpainValidTests,
					want:   true,
				},
				{
					values: mocks.SpainInvalidTests,
					want:   false,
				},
			},
			country: countries.Spain,
		},
		{
			name: "Sweden",
			args: []args{
				{
					values: mocks.SwedenValidTests,
					want:   true,
				},
				{
					values: mocks.SwedenInvalidTests,
					want:   false,
				},
			},
			country: countries.Sweden,
		},
		{
			name: "Switzerland",
			args: []args{
				{
					values: mocks.SwitzerlandValidTests,
					want:   true,
				},
				{
					values: mocks.SwitzerlandInvalidTests,
					want:   false,
				},
			},
			country: countries.Switzerland,
		},
		{
			name: "Turkey",
			args: []args{
				{
					values: mocks.TurkeyValidTests,
					want:   true,
				},
				{
					values: mocks.TurkeyInvalidTests,
					want:   false,
				},
			},
			country: countries.Turkey,
		},
		{
			name: "Ukraine",
			args: []args{
				{
					values: mocks.UkraineValidTests,
					want:   true,
				},
				{
					values: mocks.UkraineInvalidTests,
					want:   false,
				},
			},
			country: countries.Ukraine,
		},
		{
			name: "United Kingdom",
			args: []args{
				{
					values: mocks.UnitedKingdomValidTests,
					want:   true,
				},
				{
					values: mocks.UnitedKingdomInvalidTests,
					want:   false,
				},
			},
			country: countries.UnitedKingdom,
		},
		{
			name: "Multiple countries",
			args: []args{
				{
					values: mocks.SpainValidTests,
					want:   true,
				},
				{
					values: mocks.FranceValidTests,
					want:   true,
				},
				{
					values: mocks.GermanyValidTests,
					want:   true,
				},
				{
					values: mocks.ItalyValidTests,
					want:   true,
				},
			},
		},
	}
	for _, tt := range tests {
		for _, arg := range tt.args {

			count := 0

			for _, value := range arg.values {

				count++

				valid := "VALID"

				if !arg.want {
					valid = "INVALID"
				}

				name := fmt.Sprintf("%s %s nº %d: %s", tt.name, valid, count, value)

				t.Run(
					name, func(t *testing.T) {

						got := new(CheckResult)

						if tt.country == nil {
							got, _ = Validate(value)
						} else {
							got, _ = Validate(value, tt.country)
						}

						if got.IsValid != arg.want {
							t.Errorf("Validate() got = %v, want %v", got, arg.want)
						}
					},
				)

			}

		}

	}
}
