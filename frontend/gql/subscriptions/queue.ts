import { graphql } from "@/gql/graphql";

export const queue = graphql(
	/* GraphQL */ `
subscription queue {
    queue {
    name
	  memoryUsageBytes
	  size
	  groups
	  latencyMsec
	  displayLatency
	  active
	  pending
	  aggregating
	  scheduled
	  retry
	  archived
	  completed
	  processed
	  succeeded
	  failed
	  paused
	  timestamp
    }
}
`,
);
