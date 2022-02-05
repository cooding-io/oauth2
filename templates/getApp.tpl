{
		allApps(condition:{clientId:"%s"}) {
		  totalCount
		  edges {
			node {
			  id
			  clientId
			  secret
			  url
			}
		  }
		}
	  }