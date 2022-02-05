mutation MyMutation {
		__typename
		createUser(input: {user: {lastName: "%s", name: "%s", picture: "", email: "%s", moodleUdp: false, ready: true, role: "basic"}}) {
		  clientMutationId
		}
	  }