"use client";
import { Button } from "@/components/ui/button";
import {
	Dialog,
	DialogClose,
	DialogContent,
	DialogHeader,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { LoginMutation } from "@/gql/mutations/login";
import { SignupMutation } from "@/gql/mutations/signup";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useCookies } from "next-client-cookies";
import { usePathname, useRouter } from "next/navigation";
import { useState } from "react";

export default function Login() {
	const router = useRouter();
	const cookies = useCookies();
	const pathname = usePathname();
	const queryClient = useQueryClient();

	const login = useMutation(LoginMutation(queryClient, router, cookies));
	const signup = useMutation(SignupMutation(queryClient, router, cookies));

	const [email, setEmail] = useState<string>("");
	const [password, setPassword] = useState<string>("");

	return (
		<Dialog
			open={pathname.includes("/login")}
			onOpenChange={(open: boolean) => {
				if (!open) {
					setEmail("");
					setPassword("");
					router.push("/");
				}
			}}
		>
			<DialogContent>
				<Tabs>
					<DialogHeader>
						<TabsList className="grid w-full grid-cols-2 mt-4">
							<TabsTrigger value="login">Login</TabsTrigger>
							<TabsTrigger value="register">Register</TabsTrigger>
						</TabsList>
					</DialogHeader>
					<TabsContent value="login">
						<div className="grid grid-cols-4 space-y-4 items-center">
							<Label>Email</Label>
							<Input
								type="email"
								className="col-span-3"
								value={email}
								onChange={(e) => setEmail(e.target.value)}
							/>
							<Label>Password</Label>
							<Input
								type="password"
								className="col-span-3"
								value={password}
								onChange={(e) => setPassword(e.target.value)}
							/>
						</div>
						<DialogClose asChild>
							<Button
								className="mt-4	"
								onClick={() => login.mutate({ email, password })}
							>
								Login
							</Button>
						</DialogClose>
					</TabsContent>
					<TabsContent value="register">
						<div className="grid grid-cols-4 space-y-4 items-center">
							<Label>Email</Label>
							<Input
								type="email"
								className="col-span-3"
								value={email}
								onChange={(e) => setEmail(e.target.value)}
							/>
							<Label>Password</Label>
							<Input
								type="password"
								className="col-span-3"
								value={password}
								onChange={(e) => setPassword(e.target.value)}
							/>
						</div>
						<DialogClose asChild>
							<Button
								className="mt-4	"
								onClick={() => signup.mutate({ email, password })}
							>
								Sign Up
							</Button>
						</DialogClose>
					</TabsContent>
				</Tabs>
			</DialogContent>
		</Dialog>
	);
}
