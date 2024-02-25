"use client";
import { DataTable } from "@/components/table/dataTable";
import { FacetedFilter } from "@/components/table/filter";
import LinkButton from "@/components/ui/link-button";
import { AppleConnection } from "@/gql/queries/appleConnection";
import { AppleVarieties } from "@/lib/types";
import { enumToTitleCase } from "@/lib/utils";
import { PlusIcon } from "@radix-ui/react-icons";
import { Columns } from "./columns";
export default function Comps() {
	return (
		<DataTable columns={Columns(false)} queryFn={AppleConnection}>
			<DataTable.Filters>
				<FacetedFilter
					title="Variety"
					columnName="variety"
					options={AppleVarieties.map((variety) => ({
						label: enumToTitleCase(variety),
						value: variety,
					}))}
				/>
			</DataTable.Filters>
			<DataTable.Actions>
				<LinkButton
					title="Create Apples"
					link="/apples/create"
					icon={<PlusIcon className="h-4 w-4 shrink-0" />}
				/>
			</DataTable.Actions>
		</DataTable>
	);
}
