"use client"; // Error components must be Client Components
import { Button } from "@/components/ui/button";
import { useEffect } from "react";

export default function ErrorHandler({
	error,
	reset,
}: {
	error: Error & { digest?: string };
	reset: () => void;
}) {
	useEffect(() => {
		console.error(error);
	}, [error]);

	return (
		<div className="flex flex-col items-center justify-center min-h-screen">
			<h1 className="text-4xl font-bold text-customcolor-500">
				Something went wrong!
			</h1>
			<p className="text-gray-700 mb-4">
				An unexpected error has occurred. Please try again.
			</p>
			<div className="w-2/3 text-left mb-3">
				<h3 className="text-xl font-bold text-customcolor-700">
					Error message:
				</h3>
			</div>
			<pre className="p-4 bg-gray-800 text-white rounded mb-4 w-2/3">
				{error.message}
			</pre>
			<Button onClick={() => reset()}>Reload page</Button>
		</div>
	);
}
