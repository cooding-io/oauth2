{
    allUsers(condition: { email: "%s" }) {
        totalCount
        edges {
        node {
            id
            name
            lastName
            email
            password
            picture
            moodleUdp
        }
        }
    }
}