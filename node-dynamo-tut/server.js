const {DynamoDBClient, CreateTableCommand} = require("@aws-sdk/client-dynamodb")
const {DynamoDBDocumentClient, PutCommand, GetCommand, UpdateCommand} = require("@aws-sdk/lib-dynamodb")

const client = new DynamoDBClient({
    region: "us-east-1"
})

async function createTable() {
    const command = new CreateTableCommand({
        TableName: "23bcd59_student_table",
        AttributeDefinitions: [
            {
                AttributeName: "studentId", AttributeType: "S"
            }
        ],
        KeySchema: [
            {
                AttributeName: "studentId", KeyType: "HASH"
            }
        ],
        BillingMode: "PAY_PER_REQUEST"

    })
    await client.send(command)
    console.log("Table successfully created!")
}

const data = [
    {
        "studentId": "23bcd59",
    },
    {
        "studentId": "23bcs221",
    }
]

docClient = DynamoDBDocumentClient.from(client)

async function addStudents() {
    const inserts = data.map((student) => docClient.send(new PutCommand({TableName: "23bcd59_student_table", Item: student})))

    await Promise.all(inserts)
    console.log(`Inserted ${data.length} students`)

}

// createTable()
addStudents()


new UpdateCommand({
    TableName: 'sdas',
    Key: {},
    UpdateExpression: "SET age = :a, #role = :r",
    ExpressionAttributeNames: {
        "#role": "role"
    },
    ExpressionAttributeValues: {
        ":a" : 22
    }
})