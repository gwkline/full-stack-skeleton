import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
	schema: '../backend/graph/schema.graphqls',
	// documents: './src/**/*.gql', //uncomment when you have actual code
	generates: {
		'./src/lib/graphql/generated.ts': {
			plugins: [
				'typescript',
				'typescript-operations',
				'typed-document-node',
				'@kitql/graphql-codegen'
			]
		}
	}
};
export default config;
