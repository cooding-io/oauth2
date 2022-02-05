{
		allApps {
		  edges {
			node {
			  id
			  name
			  description
			  usersAppsByAppId(condition: { userId: {{.ID}} ,enable:true }) {
				enable: totalCount
			   }
			  userByOwnerId {
				name
				lastName
				email
			  }
			}
		  }
		}
		user: userById(id: {{.ID}}) {
		  id
		  name
		  lastName
		  email
		  picture
		  moodleUdp
		  public
		  cover
		  aboutMe
		  UDP: authsByUserId(condition: { system: 1 }) {
			edges {
			  node {
				custom
			  }
			}
		  }
	  
		  sectionsByOwnerId {
			edges {
			  node {
				id
				semester
				enable
				year
				section
				custom
				filesSectionsBySectionId {
				  edges {
					node {
					  fileByFileId {
						id
						name
						url
						userByOwnerId {
						  id
						  name
						  lastName
						}
					  }
					}
				  }
				}
				tasksBySectionId {
				  edges {
					node {
					  id
					  name
					  description
					  start
					  finish
					  integrates
					  role
					  taskFilesByTaskId {
						edges {
						  node {
							id
							fileByFileId {
							  id
							}
						  }
						}
					  }
					}
				  }
				}
				usersSectionsBySectionId {
				  edges {
					node {
					  id
					  role
					  userByUserId {
						id
						name
						lastName
						email
					  }
					}
				  }
				}
				messagesBySectionId(
				  condition: { forum: true }
				  first: 4
				  orderBy: ID_DESC
				) {
				  edges {
					node {
					  id
					  title
					  message
					  messageByMessageId {
						id
						title
						message
						userByOwnerId {
						  id
						  name
						  lastName
						}
					  }
					  userByOwnerId {
						id
						name
						lastName
					  }
					}
				  }
				}
				courseByCourseId {
				  id
				  name
				  code
				  institutionByOwnerInstitutionId {
					id
					name
				  }
				}
			  }
			}
		  }
	  
		  usersSectionsByUserId {
			edges {
			  node {
				role
				sectionBySectionId {
				  id
				  semester
				  year
				  enable
				  section
				  custom
				  tasksBySectionId(condition: { enable: true }) {
					edges {
					  node {
						id
						name
						description
						start
						finish
						integrates
						role
						taskFilesByTaskId {
						  edges {
							node {
							  id
							  fileByFileId {
								id
							  }
							}
						  }
						}
					  }
					}
				  }
				  oldTasks: tasksBySectionId(condition: { enable: false }) {
					edges {
					  node {
						id
						name
						description
						taskAnswersByTaskId(condition: { ownerId: {{.ID}} }) {
						  edges {
							node {
							  id
							  review
							  calification
							}
						  }
						}
	  
						taskFilesByTaskId {
						  edges {
							node {
							  id
							  fileByFileId {
								id
							  }
							}
						  }
						}
					  }
					}
				  }
				  userByOwnerId {
					id
					name
					lastName
					email
					aboutMe
				  }
				  usersSectionsBySectionId {
					edges {
					  node {
						id
						role
						userByUserId {
						  id
						  name
						  lastName
						  email
						}
					  }
					}
				  }
				  messagesBySectionId(
					condition: { forum: true }
					first: 4
					orderBy: ID_DESC
				  ) {
					edges {
					  node {
						id
						title
						message
						messageByMessageId {
						  id
						  title
						  message
						  userByOwnerId {
							id
							name
							lastName
						  }
						}
						userByOwnerId {
						  id
						  name
						  lastName
						}
					  }
					}
				  }
	  
				  courseByCourseId {
					id
					name
					code
					description
					video
					institutionByOwnerInstitutionId {
					  id
					  name
					}
				  }
	  
				  filesSectionsBySectionId {
					edges {
					  node {
						fileByFileId {
						  id
						  name
						  url
						}
					  }
					}
				  }
				}
			  }
			}
		  }
		}
	  }