import { sveltekit } from '@sveltejs/kit/vite';
import houdini from 'houdini/vite';
import { defineConfig } from 'vite';
import { sentrySvelteKit } from "@sentry/sveltekit";

export default defineConfig({
	plugins: [
		sentrySvelteKit({
			sourceMapsUploadOptions: {
				org: process.env.SENTRY_FE_ORG_NAME,
				project: process.env.SENTRY_FE_PROJ_NAME,
				authToken: process.env.SENTRY_FE_AUTH_TOKEN,
				cleanArtifacts: true
			}
		}),
		houdini(),
		sveltekit()
	],
	server: {
		port: 3000
	}
});
