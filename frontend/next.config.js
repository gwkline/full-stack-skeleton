// Define base Next.js configuration
const nextConfig = {
	images: {
		remotePatterns: [
			{
				protocol: "https",
				hostname: "imgur.com",
			},
		],
	},
	reactStrictMode: true,
};

// Compose the configurations
const config = nextConfig;
if (process.env.NODE_ENV === "production") {
	module.exports = config;
}

// const withBundleAnalyzer = require("@next/bundle-analyzer")({
// 	enabled: process.env.ANALYZE === "true",
// });
// module.exports = withBundleAnalyzer(nextConfig);
