import LoginButton from "@/components/login/server";
import { Separator } from "@/components/ui/separator";
import { ThemeToggle } from "./theme";

export default function Navbar() {
	return (
		<div className="sticky top-0 z-50 bg-background">
			<div className="grid grid-cols-3 items-center py-2 px-8">
				<div className="col-span-2 justify-start flex gap-4" />
				<div className="col-span-1 justify-self-end flex gap-4">
					<ThemeToggle />
					<LoginButton />
				</div>
			</div>
			<Separator />
		</div>
	);
}
