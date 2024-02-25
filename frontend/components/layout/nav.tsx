"use client";

import { buttonVariants } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import {
	Tooltip,
	TooltipContent,
	TooltipTrigger,
} from "@/components/ui/tooltip";
import { cn } from "@/lib/utils";
import { IconProps } from "@radix-ui/react-icons/dist/types";
import { LucideIcon } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useContext } from "react";
import { AuthContext } from "../providers/authClient";

type Role = "admin" | "user";
interface NavProps {
	isCollapsed: boolean;
	links: {
		title: string;
		link: string;
		role: Role;
		icon:
			| React.ForwardRefExoticComponent<
					IconProps & React.RefAttributes<SVGSVGElement>
			  >
			| LucideIcon;
	}[];
}

export function Nav({ links, isCollapsed }: NavProps) {
	const { user } = useContext(AuthContext);
	const pathname = usePathname();
	if (!user) {
		return (
			<div
				data-collapsed={isCollapsed}
				className="group flex flex-col gap-4 py-2 data-[collapsed=true]:py-2"
			>
				<nav className="grid gap-1 px-2 group-[[data-collapsed=true]]:justify-center group-[[data-collapsed=true]]:px-2">
					{isCollapsed ? (
						<>
							<Skeleton className="h-9 w-9 rounded-md px-3 text-xs flex items-center" />
							<Skeleton className="h-9 w-9 rounded-md px-3 text-xs flex items-center" />
							<Skeleton className="h-9 w-9 rounded-md px-3 text-xs flex items-center" />
							<Skeleton className="h-9 w-9 rounded-md px-3 text-xs flex items-center" />
							<Skeleton className="h-9 w-9 rounded-md px-3 text-xs flex items-center" />
						</>
					) : (
						<>
							<Skeleton className="h-8 rounded-md px-3 text-xs w-full" />
							<Skeleton className="h-8 rounded-md px-3 text-xs w-full" />
							<Skeleton className="h-8 rounded-md px-3 text-xs w-full" />
							<Skeleton className="h-8 rounded-md px-3 text-xs w-full" />
							<Skeleton className="h-8 rounded-md px-3 text-xs w-full" />
						</>
					)}
				</nav>
			</div>
		);
	}
	return (
		<div
			data-collapsed={isCollapsed}
			className="group flex flex-col gap-4 py-2 data-[collapsed=true]:py-2"
		>
			<nav className="grid gap-1 px-2 group-[[data-collapsed=true]]:justify-center group-[[data-collapsed=true]]:px-2">
				{links
					.filter(
						(link) =>
							!(
								link.role === "admin" &&
								(!user || user.role.toLowerCase() !== "admin")
							),
					)
					.map((link, index) =>
						isCollapsed ? (
							<Tooltip key={`${index * 2}b`} delayDuration={0}>
								<TooltipTrigger asChild>
									<Link
										href={link.link}
										className={cn(
											buttonVariants({
												variant: pathname.includes(link.link)
													? "primary"
													: "ghost",
												size: "icon",
											}),
											"h-9 w-9",
										)}
									>
										<link.icon className="h-4 w-4" />
										<span className="sr-only">{link.title}</span>
									</Link>
								</TooltipTrigger>
								<TooltipContent
									side="right"
									className="flex items-center gap-4"
								>
									{link.title}
								</TooltipContent>
							</Tooltip>
						) : (
							<Link
								key={`${index * 2}a`}
								href={link.link}
								className={cn(
									buttonVariants({
										variant: pathname.includes(link.link) ? "primary" : "ghost",
										size: "sm",
									}),
									"justify-start",
								)}
							>
								<link.icon className="mr-2 h-4 w-4" />
								{link.title}
							</Link>
						),
					)}
			</nav>
		</div>
	);
}
