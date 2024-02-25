import { ActionsCell } from "@/components/table/cells/actions";
import { BadgeCell } from "@/components/table/cells/badge";
import { IdCell } from "@/components/table/cells/id";
import { SelectCell } from "@/components/table/cells/select";
import { badgeVariants } from "@/components/ui/badge";
import { appleEdge } from "@/gql/queries/appleConnection";
import { enumToTitleCase } from "@/lib/utils";
import { ColumnDef } from "@tanstack/react-table";
import { ResultOf } from "gql.tada";

export function Columns(
	isLoading: boolean,
): ColumnDef<ResultOf<typeof appleEdge>>[] {
	return [
		SelectCell(),
		IdCell(isLoading),
		BadgeCell({
			id: "variety",
			title: "Variety",
			isLoading,
			accessorKey: "node.variety",
			accessorFn: (edge) => enumToTitleCase(edge.node.variety),
			variantAccessorFn: (edge) => "outline" as keyof typeof badgeVariants,
		}),
		ActionsCell({ isLoading }),
	];
}
