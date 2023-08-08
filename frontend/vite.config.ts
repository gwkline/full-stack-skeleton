import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import { sentrySvelteKit } from "@sentry/sveltekit";
import dotenv from 'dotenv';
dotenv.config();

export default defineConfig({
	plugins: [sentrySvelteKit({
		sourceMapsUploadOptions: {
		  org: process.env.SENTRY_FE_ORG_NAME,
		  project: process.env.SENTRY_FE_PROJ_NAME,
		  authToken: process.env.SENTRY_FE_AUTH_TOKEN,
		  cleanArtifacts: true,
		},
	  }), sveltekit()],
	server: {
		port: 3000
	}
});
