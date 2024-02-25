import { Row } from "@tanstack/react-table";

export type JWT = {
	AccessToken: string;
	RefreshToken: string;
};

export interface DataTableRowActionsProps<TData> {
	row: Row<TData>;
}

export const AppleVarieties = [
	"FUJI",
	"GALA",
	"HONEYCRISP",
	"GOLDEN_DELICIOUS",
	"RED_DELICIOUS",
	"GRANNY_SMITH",
	"BRAEBURN",
	"JONAGOLD",
	"CRIPPS_PINK",
	"MCINTOSH",
	"EMPIRE",
	"JONATHAN",
	"CORTLAND",
	"WINESAP",
	"AMBROSIA",
	"COSMIC_CRISP",
	"ENVY",
	"JAZZ",
];

export type AppleVariety =
	| "FUJI"
	| "GALA"
	| "HONEYCRISP"
	| "GOLDEN_DELICIOUS"
	| "RED_DELICIOUS"
	| "GRANNY_SMITH"
	| "BRAEBURN"
	| "JONAGOLD"
	| "CRIPPS_PINK"
	| "MCINTOSH"
	| "EMPIRE"
	| "JONATHAN"
	| "CORTLAND"
	| "WINESAP"
	| "AMBROSIA"
	| "COSMIC_CRISP"
	| "ENVY"
	| "JAZZ";
