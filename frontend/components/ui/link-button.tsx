import { ReactNode } from "react";
import { Tooltip, TooltipContent, TooltipTrigger } from "./tooltip";
import Link from "next/link";
import { Button } from "./button";

export default function LinkButton({
	link,
	title,
	icon,
	variant = "primary",
}: {
	link: string;
	title: string;
	icon: ReactNode;
	variant?:
		| "link"
		| "default"
		| "destructive"
		| "outline"
		| "primary"
		| "secondary"
		| "ghost"
		| null
		| undefined;
}) {
	return (
		<Tooltip>
			<TooltipTrigger asChild>
				<Link href={link}>
					<Button
						variant={variant}
						className="h-8 w-8 items-center flex justify-center align-center"
					>
						{icon}
					</Button>
				</Link>
			</TooltipTrigger>
			<TooltipContent>{title}</TooltipContent>
		</Tooltip>
	);
}
