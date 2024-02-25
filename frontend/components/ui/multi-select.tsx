import * as React from "react";
import { CheckIcon, PlusCircledIcon } from "@radix-ui/react-icons";

import { cn } from "@/lib/utils";
import {
	Command,
	CommandEmpty,
	CommandGroup,
	CommandInput,
	CommandItem,
	CommandList,
	CommandSeparator,
} from "@/components/ui/command";
import { Popover, PopoverContent, PopoverTrigger } from "./popover";
import { Button, ButtonProps } from "./button";
import { Separator } from "./separator";
import { Badge } from "./badge";

interface MultiSelectProps<TData, TValue> extends ButtonProps {
	title?: string;
	selectedOptions: Set<string>;
	setSelectedOptions: (x: Set<string>) => void;
	options: {
		label: string;
		value: string;
		icon?: React.ComponentType<{ className?: string }>;
	}[];
}

export function MultiSelect<TData, TValue>({
	title,
	className,
	selectedOptions,
	setSelectedOptions,
	options,
	...buttonProps
}: MultiSelectProps<TData, TValue>) {
	return (
		<Popover>
			<PopoverTrigger asChild>
				<Button
					variant="outline"
					size="sm"
					className={cn("h-8 border-dashed", className)}
					{...buttonProps}
				>
					<PlusCircledIcon className="mr-2 h-4 w-4" />
					{title}
					{selectedOptions.size > 0 && (
						<>
							<Separator orientation="vertical" className="mx-2 h-4" />
							<Badge
								variant="secondary"
								className="rounded-sm px-1 font-normal lg:hidden"
							>
								{selectedOptions.size}
							</Badge>
							<div className="hidden space-x-1 lg:flex">
								{selectedOptions.size > 2 ? (
									<Badge
										variant="secondary"
										className="rounded-sm px-1 font-normal"
									>
										{selectedOptions.size} selected
									</Badge>
								) : (
									options
										.filter((option) => selectedOptions.has(option.value))
										.map((option) => (
											<Badge
												variant="secondary"
												key={option.value}
												className="rounded-sm px-1 font-normal"
											>
												{option.label}
											</Badge>
										))
								)}
							</div>
						</>
					)}
				</Button>
			</PopoverTrigger>
			<PopoverContent className="w-[200px] p-0" align="start">
				<Command>
					<CommandInput placeholder={title} />
					<CommandList className="overflow-y-scroll">
						<CommandEmpty>No results found.</CommandEmpty>
						<CommandGroup>
							{options.map((option) => {
								const isSelected = selectedOptions.has(option.value);
								return (
									<CommandItem
										key={option.value}
										onSelect={() => {
											if (isSelected) {
												selectedOptions.delete(option.value);
											} else {
												selectedOptions.add(option.value);
											}
											setSelectedOptions(new Set(selectedOptions));
										}}
									>
										<div
											className={cn(
												"mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary",
												isSelected
													? "bg-primary text-primary-foreground"
													: "opacity-50 [&_svg]:invisible",
											)}
										>
											<CheckIcon className={cn("h-4 w-4")} />
										</div>
										{option.icon && (
											<option.icon className="mr-2 h-4 w-4 text-muted-foreground" />
										)}
										<span>{option.label}</span>
									</CommandItem>
								);
							})}
						</CommandGroup>
						<CommandSeparator />
					</CommandList>
					{selectedOptions.size > 0 && (
						<CommandGroup>
							<CommandItem
								onSelect={() => setSelectedOptions(new Set())}
								className="justify-center text-center"
							>
								Clear filters
							</CommandItem>
						</CommandGroup>
					)}
				</Command>
			</PopoverContent>
		</Popover>
	);
}
