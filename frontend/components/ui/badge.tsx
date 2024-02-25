import { type VariantProps, cva } from "class-variance-authority";
import * as React from "react";

import { cn } from "@/lib/utils";

const badgeVariants = cva(
	"inline-flex items-center rounded-md border border-stone-200 px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-stone-950 focus:ring-offset-2 dark:border-stone-800 dark:focus:ring-stone-300",
	{
		variants: {
			variant: {
				default:
					"border-transparent bg-stone-900 text-stone-50 shadow hover:bg-stone-900/80 dark:bg-stone-50 dark:text-stone-900 dark:hover:bg-stone-50/80",
				secondary:
					"border-transparent bg-stone-100 text-stone-900 hover:bg-stone-100/80 dark:bg-stone-800 dark:text-stone-200 dark:hover:bg-stone-800/80",
				destructive:
					"border-transparent bg-red-500 text-stone-50 shadow hover:bg-red-500/80 dark:bg-red-900 dark:text-stone-200 dark:hover:bg-red-900/80",
				outline: "text-stone-950 dark:text-stone-200",
				pending:
					"border-transparent bg-yellow-500 text-stone-50 shadow hover:bg-yellow-500/80 dark:bg-yellow-700 dark:text-stone-200 dark:hover:bg-yellow-700/80",
				queued:
					"border-transparent bg-orange-500 text-stone-50 shadow hover:bg-orange-500/80 dark:bg-orange-700 dark:text-stone-200 dark:hover:bg-orange-700/80",
				working:
					"border-transparent bg-blue-500 text-stone-50 shadow hover:bg-blue-500/80 dark:bg-blue-900 dark:text-stone-200 dark:hover:bg-blue-900/80",
				completed:
					"border-transparent bg-purple-500 text-stone-50 shadow hover:bg-purple-500/80 dark:bg-purple-900 dark:text-stone-200 dark:hover:bg-purple-900/80",
				failed:
					"border-transparent bg-red-500 text-stone-50 shadow hover:bg-red-500/80 dark:bg-red-900 dark:text-stone-200 dark:hover:bg-red-900/80",
				stopped:
					"border-transparent bg-gray-500 text-stone-50 shadow hover:bg-gray-500/80 dark:bg-gray-700 dark:text-stone-200 dark:hover:bg-gray-700/80",
				loading:
					"border-transparent bg-stone-900/10 text-stone-900/10 shadow dark:bg-stone-50/10 dark:text-stone-200/10 animate-pulse",
			},
		},
		defaultVariants: {
			variant: "default",
		},
	},
);

export interface BadgeProps
	extends React.HTMLAttributes<HTMLDivElement>,
		VariantProps<typeof badgeVariants> {}

function Badge({ className, variant, ...props }: BadgeProps) {
	return (
		<div className={cn(badgeVariants({ variant }), className)} {...props} />
	);
}

export { Badge, badgeVariants };
