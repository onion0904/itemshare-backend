# Query

user(id: String!): User! -> OK
group(id: String!): Group! -> OK
event(id: String!): Event! -> OK
eventsByMonth(input: MonthlyEventInput!): [String!]! -> OK



# Mutation

createUser(input: CreateUserInput!): User! -> OK

updateUser( -> OK
id: String!
input: UpdateUserInput!
): User!

deleteUser(id: String!): Boolean! -> OK

createGroup(input: CreateGroupInput!): Group! -> OK

updateGroup( -> OK
id: String!
input: UpdateGroupInput!
): Group!

deleteGroup(id: String!): Boolean! -> OK

addUserToGroup( -> OK
groupID: String!
userID: String!
): Group!

removeUserFromGroup( -> OK
groupID: String!
userID: String!
): Group!

addEventToGroup( -> OK
groupID: String!
eventID: String!
): Group!

createEvent(input: CreateEventInput!): Event! -> OK

deleteEvent(id: String!): Boolean! -> OK

sendVerificationCode(email: String!): Boolean! -> OK

signup( ->OK
input: CreateUserInput!
vcode: String!
): AuthUserResponse!

signin( -> OK
email: String!
password: String!
): AuthUserResponse!