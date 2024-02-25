"use client";
import { AuthContext } from "@/components/providers/authClient";
import { Button } from "@/components/ui/button";
import {
	Dialog,
	DialogContent,
	DialogFooter,
	DialogHeader,
	DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "@/components/ui/select";
import { toast } from "@/components/ui/use-toast";
import { CreateApple } from "@/gql/mutations/createApple";
import { AppleVarieties, AppleVariety } from "@/lib/types";
import { enumToTitleCase } from "@/lib/utils";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useCookies } from "next-client-cookies";
import { usePathname } from "next/navigation";
import { useRouter } from "next/navigation";
import { useContext, useState } from "react";

export default function Create() {
	const router = useRouter();
	const cookies = useCookies();
	const pathname = usePathname();
	const queryClient = useQueryClient();

	const { mutate } = useMutation(CreateApple(queryClient, cookies, router));
	const { user } = useContext(AuthContext);

	const [variety, setVariety] = useState<AppleVariety | undefined>(undefined);
	const [quantity, setQuantity] = useState<number>(0);

	return (
		<Dialog
			open={pathname.includes("/create")}
			onOpenChange={(open: boolean) => {
				if (!open) {
					setVariety(undefined);
					setQuantity(0);
					router.push("/apples");
				}
			}}
		>
			<DialogContent>
				<DialogHeader>
					<DialogTitle>Create Apple</DialogTitle>
				</DialogHeader>
				<div className="grid gap-4 grid-cols-6">
					<div className="col-span-3 items-center gap-4">
						<Label htmlFor="variety" className="whitespace-nowrap">
							Variety
						</Label>
						<Select
							name="variety"
							value={variety}
							onValueChange={(e) => setVariety(e as AppleVariety)}
						>
							<SelectTrigger className="col-span-2 h-8">
								<SelectValue placeholder="Pick a variety..." />
							</SelectTrigger>
							<SelectContent>
								{AppleVarieties.map((variety) => (
									<SelectItem value={variety}>
										{enumToTitleCase(variety)}
									</SelectItem>
								))}
							</SelectContent>
						</Select>
					</div>

					<div className="col-span-3 items-center gap-4">
						<Label htmlFor="quantity" className="whitespace-nowrap">
							Quantity
						</Label>
						<Input
							id="quantity"
							name="quantity"
							value={quantity}
							onChange={(e) => setQuantity(parseFloat(e.target.value))}
							className="col-span-2 h-8"
							type="number"
						/>
					</div>
				</div>

				<DialogFooter>
					<Button
						className="bg-customcolor-500"
						onClick={() => {
							if (quantity === 0 || variety === undefined || !user) {
								toast({
									title: "Error",
									description: "Please enter a variety / quantity",
									variant: "destructive",
								});
								return;
							}

							mutate({ variety, quantity, userId: user?.id });
						}}
					>
						Submit
					</Button>
				</DialogFooter>
			</DialogContent>
		</Dialog>
	);
}
