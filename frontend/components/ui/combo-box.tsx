import * as React from "react";
import { CaretSortIcon, CheckIcon } from "@radix-ui/react-icons";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
	Command,
	CommandEmpty,
	CommandGroup,
	CommandInput,
	CommandItem,
	CommandSeparator,
} from "@/components/ui/command";
import {
	Popover,
	PopoverContent,
	PopoverTrigger,
} from "@/components/ui/popover";

type Option = {
	value: string;
	label: string;
};

export function Combobox({
	options,
	onOptionChange,
	selectedOption,
	itemName = "options",
	nullable = true,
	className,
}: {
	options: Option[];
	selectedOption: string | undefined;
	onOptionChange: (x: Option | undefined) => void;
	itemName?: string;
	className?: string;
	nullable?: boolean;
}) {
	const [open, setOpen] = React.useState(false);

	return (
		<Popover open={open} onOpenChange={setOpen} modal={true}>
			<PopoverTrigger asChild>
				<Button
					variant="outline"
					role="combobox"
					aria-expanded={open}
					className={cn(
						"w-[200px] justify-between",
						!selectedOption && "text-muted-foreground",
						className,
					)}
				>
					{selectedOption ? selectedOption : `Select ${itemName}...`}
					<CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
				</Button>
			</PopoverTrigger>
			<PopoverContent className="w-[200px] p-0 max-h-[200px] overflow-y-scroll">
				<Command>
					<CommandInput placeholder={`Search...`} className="h-9" />
					<CommandEmpty>No {itemName}s found</CommandEmpty>
					<CommandGroup>
						{nullable && (
							<CommandItem
								key="none"
								value="None"
								onSelect={() => {
									onOptionChange(undefined);
									setOpen(false);
								}}
							>
								No {itemName}
								<CheckIcon
									className={cn(
										"ml-auto h-4 w-4",
										selectedOption === null ? "opacity-100" : "opacity-0",
									)}
								/>
							</CommandItem>
						)}
						{options.map((option) => (
							<CommandItem
								key={option.value}
								value={option.value}
								onSelect={(currentValue) => {
									const option = options.find(
										(option) =>
											option.value.toLowerCase() === currentValue.toLowerCase(),
									);
									if (!option) return;
									onOptionChange(option);
									setOpen(false);
								}}
							>
								{option.label}
								<CheckIcon
									className={cn(
										"ml-auto h-4 w-4",
										selectedOption === option.value
											? "opacity-100"
											: "opacity-0",
									)}
								/>
							</CommandItem>
						))}
						<CommandSeparator />
					</CommandGroup>
				</Command>
			</PopoverContent>
		</Popover>
	);
}
